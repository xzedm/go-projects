Simple Web Scraper
A Go program that scrapes webpage titles from a list of URLs and displays them along with the URLs. It uses the net/http package for HTTP requests and golang.org/x/net/html for HTML parsing.
Features

Fetches page titles from a list of URLs
Handles HTTP errors (e.g., 404, 500) and invalid URLs
Stores and displays URL-title pairs
Adds "https://" to URLs without a scheme for convenience

Prerequisites

Go 1.16 or higher
Install the golang.org/x/net/html package:go get golang.org/x/net/html



Installation

Clone or download the project.
Navigate to the project directory.
Run the program:go run scraper.go



Usage
The program scrapes titles from a hardcoded list of URLs defined in scraper.go. To modify the URLs, edit the urls slice in the main function.
Example output:
Web Scraper: Fetching page titles...
Error for invalid-url: failed to fetch https://invalid-url: Get "https://invalid-url": dial tcp: lookup invalid-url: no such host

Results:
URL: example.com
Title: Example Domain

URL: golang.org
Title: The Go Programming Language

URL: github.com
Title: GitHub: Let’s build from here · GitHub

URL: invalid-url
Title: Error: failed to fetch https://invalid-url: Get "https://invalid-url": dial tcp: lookup invalid-url: no such host

File Structure

scraper.go: Main source code
README.md: This file

Notes

The program automatically adds "https://" to URLs lacking a scheme.
Errors (e.g., invalid URLs, HTTP errors, missing titles) are displayed in the results.
The code follows Go best practices, including proper error handling and formatting with gofmt.
To extend the program, you can modify it to accept URLs via command-line arguments or a file.

