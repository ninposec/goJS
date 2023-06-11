# goJS
 
goJS is a simple command line tool written in Go, designed to download files from a list of URLs concurrently. For each URL, the tool creates a directory named after the domain part of the URL and downloads the corresponding file into this directory.


### Usage

Example:

```bash
cat urls.txt | gojs -output jsfiles
````


````bash
Usage: gojs
Download files from a list of URLs concurrently.

Options:
  -urls string
        List of URLs separated by commas.
  -output string
        Output directory (default ".")
  -concurrency int
        Number of concurrent downloads (default 10)
  -help
        Display help

````



### Installation

```bash
go install -v https://github.com/ninposec/gojs@latest
````

or

```bash
git clone https://github.com/ninposec/gojs.git
cd gojs
go build .
```
