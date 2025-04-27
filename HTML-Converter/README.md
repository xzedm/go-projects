Markdown to HTML Converter
A Go program that converts .md (Markdown) files to .html files using the github.com/gomarkdown/markdown package. It reads the input file, converts it to HTML, and writes the output to a file with a .html extension.
Features

Converts Markdown to HTML with support for common extensions (e.g., headings, bold, links, code blocks).
Wraps output in a basic HTML structure with styling.
Accepts a .md file via command-line argument.
Validates input file extension and handles errors.

Prerequisites

Go 1.16 or higher
Install the gomarkdown package:go get github.com/gomarkdown/markdown



Installation

Clone or download the project.
Navigate to the project directory.
Install dependencies:go get github.com/gomarkdown/markdown
or run in the project folder:

go mod init project-name
go mod tidy


Run the program:go run md2html.go <input.md>



Usage
Run the program with a .md file as an argument:
go run md2html.go sample.md

This generates sample.html in the same directory.
Example input (sample.md):
# Sample Markdown Document

This is a **sample** Markdown file.

## Features
- Supports headings
- Supports **bold** and *italic* text
- Includes [links](https://example.com)

Output: Creates sample.html with styled HTML content viewable in a browser.
File Structure

md2html.go: Main source code
README.md: This file
(Optional) sample.md: Sample input file for testing

Notes

The program requires the github.com/gomarkdown/markdown package.
Output HTML includes a basic CSS style for readability.
The code follows Go best practices, including proper error handling and formatting with gofmt.
On Windows CMD, ensure file paths are correctly formatted (e.g., sample.md instead of ./sample.md).

