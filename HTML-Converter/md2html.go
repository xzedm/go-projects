package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

// convertMarkdownToHTML converts Markdown content to HTML
func convertMarkdownToHTML(md []byte) string {
	// Create parser with common extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	p := parser.NewWithExtensions(extensions)

	// Parse Markdown
	doc := p.Parse(md)

	// Create HTML renderer
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	// Convert to HTML
	htmlContent := markdown.Render(doc, renderer)

	// Wrap in basic HTML structure
	return fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Markdown to HTML</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 40px; }
        h1, h2, h3 { color: #333; }
        p { line-height: 1.6; }
    </style>
</head>
<body>
%s
</body>
</html>`, htmlContent)
}

func main() {
	// Check for command-line argument
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: go run md2html.go <input.md>")
		os.Exit(1)
	}
	inputFile := os.Args[1]

	// Verify file has .md extension
	if !strings.HasSuffix(strings.ToLower(inputFile), ".md") {
		fmt.Fprintln(os.Stderr, "Input file must have .md extension")
		os.Exit(1)
	}

	// Open and read the Markdown file
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// Read file content
	var mdContent []byte
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		mdContent = append(mdContent, scanner.Bytes()...)
		mdContent = append(mdContent, '\n')
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Convert Markdown to HTML
	htmlContent := convertMarkdownToHTML(mdContent)

	// Generate output filename (replace .md with .html)
	outputFile := strings.TrimSuffix(inputFile, filepath.Ext(inputFile)) + ".html"

	// Write HTML to output file
	err = os.WriteFile(outputFile, []byte(htmlContent), 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing output file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully converted %s to %s\n", inputFile, outputFile)
}
