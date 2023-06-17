package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var (
	urls        = flag.String("urls", "", "File List of URLs separated by new lines")
	output      = flag.String("output", ".", "Output directory")
	concurrency = flag.Int("concurrency", 10, "Number of concurrent downloads")
)

func init() {
	flag.Usage = func() {
		fmt.Println(`
		
██████╗  ██████╗      ██╗███████╗
██╔════╝ ██╔═══██╗     ██║██╔════╝
██║  ███╗██║   ██║     ██║███████╗
██║   ██║██║   ██║██   ██║╚════██║
╚██████╔╝╚██████╔╝╚█████╔╝███████║
 ╚═════╝  ╚═════╝  ╚════╝ ╚══════╝
																																																					 
				
			`)

		fmt.Println("goJS v.0.1")
		fmt.Println("Author: ninposec")
		fmt.Println("")
		fmt.Println("Usage: cat urls.txt | gojs -output jsfiles")
		fmt.Println("Download JS files from a list of JS URLs concurrently. Save files to folder per target")
		fmt.Println()
		fmt.Println("Options:")
		flag.PrintDefaults()
	}
}

func downloadFile(filepath, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to start download: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	out, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

func worker(id int, jobs <-chan string, results chan<- error, wg *sync.WaitGroup, outputDir string) {
	for j := range jobs {
		u, err := url.Parse(j)
		if err != nil {
			results <- fmt.Errorf("Worker%d: failed to parse URL %s: %v", id, j, err)
			continue
		}

		domainDir := filepath.Join(outputDir, u.Host)
		err = os.MkdirAll(domainDir, os.ModePerm)
		if err != nil {
			results <- fmt.Errorf("Worker%d: failed to create directory for URL %s: %v", id, j, err)
			continue
		}

		targetPath := filepath.Join(domainDir, filepath.Base(u.Path))
		targetPath = filepath.Clean(targetPath) // Sanitize the target path
		err = downloadFile(targetPath, j)
		if err != nil {
			results <- fmt.Errorf("Worker%d: failed to download %s: %v", id, j, err)
		} else {
			results <- nil
		}
	}
	wg.Done()
}

func main() {
	flag.Parse()

	if flag.NFlag() == 0 {
		flag.Usage()
		os.Exit(0)
	}

	var urlsList []string
	if *urls != "" {
		urlsList = strings.Split(*urls, "\n")
	} else {
		for {
			var u string
			if _, err := fmt.Scan(&u); err != nil {
				break
			}
			urlsList = append(urlsList, u)
		}
	}

	if len(urlsList) == 0 {
		fmt.Println("No URLs provided. Exiting.")
		os.Exit(1)
	}

	if *output == "" {
		fmt.Println("No output directory provided. Using current directory.")
		*output = "."
	}

	errCh := make(chan error, *concurrency) // Error channel to collect worker errors
	var wg sync.WaitGroup
	jobs := make(chan string, *concurrency) // Use a buffered channel to avoid blocking the sender goroutine

	for i := 1; i <= *concurrency; i++ {
		go worker(i, jobs, errCh, &wg, *output)
	}

	for _, url := range urlsList {
		wg.Add(1)
		jobs <- url
	}
	close(jobs)

	wg.Wait()
	close(errCh) // Close the error channel to signal the completion of error collection

	for err := range errCh {
		if err != nil {
			fmt.Println(err) // Handle or log the errors as desired
		}
	}
}
