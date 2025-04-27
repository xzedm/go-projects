package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// checkRange checks if any number in the range [start, end] divides n evenly
func checkRange(n, start, end int, result chan<- bool) {
	for i := start; i <= end; i++ {
		if i != 0 && n%i == 0 {
			result <- true // Found a divisor
			return
		}
	}
	result <- false // No divisors found in this range
}

// isPrime checks if n is prime using goroutines
func isPrime(n int, numGoroutines int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	// Calculate the square root of n to limit divisor checks
	sqrtN := int(math.Sqrt(float64(n)))
	// Adjust numGoroutines if too large for the range
	if numGoroutines > sqrtN/2 {
		numGoroutines = sqrtN / 2
	}
	if numGoroutines < 1 {
		numGoroutines = 1
	}

	// Divide the range [3, sqrtN] into numGoroutines parts
	rangeSize := (sqrtN-2)/numGoroutines + 1
	result := make(chan bool, numGoroutines)

	// Launch goroutines
	for i := 0; i < numGoroutines; i++ {
		start := 3 + i*rangeSize
		end := start + rangeSize - 1
		if end > sqrtN {
			end = sqrtN
		}
		go checkRange(n, start, end, result)
	}

	// Collect results
	for i := 0; i < numGoroutines; i++ {
		if <-result {
			return false // A divisor was found
		}
	}
	return true // No divisors found
}

func main() {
	const numGoroutines = 4 // Number of goroutines to use
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Prime Number Checker (using goroutines)")
	fmt.Println("Enter a positive integer to check if it's prime, or 'exit' to quit.")

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		input = strings.TrimSpace(input)
		if input == "exit" {
			fmt.Println("Goodbye!")
			return
		}

		n, err := strconv.Atoi(input)
		if err != nil || n < 0 {
			fmt.Println("Please enter a valid positive integer.")
			continue
		}

		if isPrime(n, numGoroutines) {
			fmt.Printf("%d is a prime number.\n", n)
		} else {
			fmt.Printf("%d is not a prime number.\n", n)
		}
	}
}
