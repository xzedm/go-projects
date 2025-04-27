package main

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// Result stores a URL and its page title
type Result struct {
	URL   string
	Title string
}

// fetchTitle retrieves the title of a webpage from a given URL
func fetchTitle(url string) (Result, error) {
	// Ensure URL has a scheme
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "https://" + url
	}

	// Send HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		return Result{URL: url, Title: ""}, fmt.Errorf("failed to fetch %s: %w", url, err)
	}
	defer resp.Body.Close()

	// Check for non-200 status codes
	if resp.StatusCode != http.StatusOK {
		return Result{URL: url, Title: ""}, fmt.Errorf("HTTP error %d: %s", resp.StatusCode, resp.Status)
	}

	// Parse HTML
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return Result{URL: url, Title: ""}, fmt.Errorf("failed to parse HTML for %s: %w", url, err)
	}

	// Extract title
	title := extractTitle(doc)
	if title == "" {
		return Result{URL: url, Title: ""}, fmt.Errorf("no title found for %s", url)
	}

	return Result{URL: url, Title: title}, nil
}

// extractTitle traverses the HTML node tree to find the <title> tag content
func extractTitle(n *html.Node) string {
	if n.Type == html.ElementNode && n.Data == "title" {
		if n.FirstChild != nil {
			return strings.TrimSpace(n.FirstChild.Data)
		}
		return ""
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if title := extractTitle(c); title != "" {
			return title
		}
	}
	return ""
}

func main() {
	// List of URLs to scrape
	urls := []string{
		"example.com",
		"golang.org",
		"github.com",
		"invalid-url", // For error handling demonstration
	}

	fmt.Println("Web Scraper: Fetching page titles...")
	results := make([]Result, 0, len(urls))

	// Fetch titles for each URL
	for _, url := range urls {
		result, err := fetchTitle(url)
		if err != nil {
			fmt.Printf("Error for %s: %v\n", url, err)
			results = append(results, Result{URL: url, Title: "Error: " + err.Error()})
			continue
		}
		results = append(results, result)
	}

	// Display results
	fmt.Println("\nResults:")
	for _, result := range results {
		fmt.Printf("URL: %s\nTitle: %s\n\n", result.URL, result.Title)
	}
}
