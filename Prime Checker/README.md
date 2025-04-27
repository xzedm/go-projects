Goroutine-based Prime Checker
A Go program that checks if a number is prime using goroutines to parallelize divisor checks. It splits the divisor range into parts, assigns each to a goroutine, and uses a channel to collect results.
Features

Checks if a user-provided number is prime.
Uses multiple goroutines to check divisor ranges concurrently.
Sends results via a channel to determine primality.
Handles invalid inputs (e.g., negative numbers, non-integers).

Prerequisites

Go 1.16 or higher

Installation

Clone or download the project.
Navigate to the project directory.
Run the program:go run primechecker.go



Usage
Run the program and enter a positive integer to check if it’s prime:
Prime Number Checker (using goroutines)
Enter a positive integer to check if it's prime, or 'exit' to quit.
> 17
17 is a prime number.
> 24
24 is not a prime number.
> invalid
Please enter a valid positive integer.
> exit
Goodbye!

File Structure

primechecker.go: Main source code
README.md: This file

Notes

The program uses 4 goroutines by default (configurable via numGoroutines constant).
Divisor checks are limited to the square root of the input number for efficiency.
The range of divisors (3 to sqrt(n)) is split evenly among goroutines.
The code follows Go best practices, including proper error handling and formatting with gofmt.

