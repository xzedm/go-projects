package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

// WordCount stores a word and its frequency
type WordCount struct {
	Word  string
	Count int
}

func main() {
	// Check for command-line argument
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: go run wordcounter.go <filename.txt>")
		os.Exit(1)
	}
	filename := os.Args[1]

	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// Count word frequencies
	wordFreq := make(map[string]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Split line into words
		words := strings.Fields(line)
		for _, word := range words {
			// Normalize: convert to lowercase, remove punctuation
			word = strings.ToLower(strings.TrimFunc(word, func(r rune) bool {
				return !unicode.IsLetter(r) && !unicode.IsNumber(r)
			}))
			if word != "" {
				wordFreq[word]++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Check if map is empty
	if len(wordFreq) == 0 {
		fmt.Println("No words found in the file.")
		return
	}

	// Convert map to slice for sorting
	counts := make([]WordCount, 0, len(wordFreq))
	for word, count := range wordFreq {
		counts = append(counts, WordCount{Word: word, Count: count})
	}

	// Sort by count (descending) and word (ascending for ties)
	sort.Slice(counts, func(i, j int) bool {
		if counts[i].Count == counts[j].Count {
			return counts[i].Word < counts[j].Word
		}
		return counts[i].Count > counts[j].Count
	})

	// Print top 5 words (or fewer if less than 5)
	fmt.Printf("Top %d most frequent words in %s:\n", min(5, len(counts)), filename)
	for i, wc := range counts[:min(5, len(counts))] {
		fmt.Printf("%d. %s: %d\n", i+1, wc.Word, wc.Count)
	}
}

// min returns the smaller of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
