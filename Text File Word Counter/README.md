Text File Word Counter
A Go program that reads a .txt file and counts word frequencies, displaying the top 5 most frequent words. It uses the os and bufio packages to read the file and a map to store word counts.
Features

Reads a text file specified via command-line argument
Counts word frequencies, ignoring case and punctuation
Displays the top 5 most frequent words, sorted by frequency (and alphabetically for ties)
Handles errors (e.g., file not found, empty files)

Prerequisites

Go 1.16 or higher

Installation

Clone or download the project.
Navigate to the project directory.
Run the program with a text file:go run wordcounter.go sample.txt



Usage
Run the program with a .txt file as an argument:
go run wordcounter.go <filename.txt>

The program will print the top 5 most frequent words in the file.
Example with sample.txt:
At the time, no single team member knew Go, but within a month, everyone was writing in Go and we were building out the endpoints. It was the flexibility, how easy it was to use, and the really cool concept behind Go (how Go handles native concurrency, garbage collection, and of course safety+speed.) that helped engage us during the build. Also, who can beat that cute mascot!

Output:
Top 5 most frequent words in sample.txt:
1. the: 5
2. go: 4
3. and: 3
4. was: 3
5. how: 2

File Structure

wordcounter.go: Main source code
README.md: This file
(Optional) sample.txt: Sample input file for testing

Notes

Words are normalized to lowercase and stripped of punctuation.
The program handles empty files and missing files with appropriate error messages.
The code follows Go best practices, including proper error handling and formatting with gofmt.
To extend the program, you could add support for multiple files or output to a file.

