# goJS
 
goJS is a simple command line tool written in Go, designed to download files from a list of URLs concurrently. For each URL, the tool creates a directory structure named after the domain part of the URL and downloads the corresponding file.

The main purpose of this tool is to download Javascript files at scale.

Can be combined with e.g. GetJS to fetch a list of Javascript URLs, then download found JS to a structured folder.


### Usage

Example:

```bash
cat urls.txt | gojs -output jsfiles
cat alive.txt | getJS --complete | gojs -output jsfiles
````


````bash
gojs -h
		
		
██████╗  ██████╗      ██╗███████╗
██╔════╝ ██╔═══██╗     ██║██╔════╝
██║  ███╗██║   ██║     ██║███████╗
██║   ██║██║   ██║██   ██║╚════██║
╚██████╔╝╚██████╔╝╚█████╔╝███████║
 ╚═════╝  ╚═════╝  ╚════╝ ╚══════╝
															
				
			
goJS v.0.1
Author: ninposec

Usage: cat urls.txt | gojs -output jsfiles
Download JS files from a list of JS URLs concurrently. Save files to folder per target

Options:
  -concurrency int
    	Number of concurrent downloads (default 10)
  -output string
    	Output directory (default ".")
  -urls string
    	File List of URLs separated by new lines

````



### Installation

```bash
go install -v https://github.com/ninposec/goJS@latest
````

or

```bash
git clone https://github.com/ninposec/gojs.git
cd gojs
go build .
```

### ToDo

* Add SourceMappingURL functionality (Download corresponding js map files if exist)