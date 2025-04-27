CLI Calculator
A Go command-line tool that performs basic mathematical operations, including addition, subtraction, multiplication, division, square roots, and exponentiation. It accepts user input in a simple format and handles errors like division by zero.
Features

Supports operations: +, -, *, /, sqrt, pow.
Accepts two numbers and an operation (or one number for sqrt).
Handles invalid inputs, division by zero, and negative square roots.
Provides an interactive CLI with an exit command.

Prerequisites

Go 1.16 or higher

Installation

Clone or download the project.
Navigate to the project directory.
Run the program:go run calculator.go



Usage
Run the program and enter inputs in the format <num1> <num2> <operation> (or <num> sqrt for square root). Type exit to quit.
Example session:
CLI Calculator (supports +, -, *, /, sqrt, pow)
Enter two numbers (or one for sqrt) and an operation, separated by spaces (e.g., '5 3 +' or '16 sqrt'). Type 'exit' to quit.
> 5 3 +
Result: 8.00
> 10 2 *
Result: 20.00
> 15 0 /
Error: Division by zero is not allowed.
> 16 sqrt
Result: 4.00
> 2 3 pow
Result: 8.00
> exit
Goodbye!

File Structure

calculator.go: Main source code
README.md: This file

Notes

Numbers can be integers or decimals (e.g., 5.5).
The program validates inputs and provides clear error messages.
Results are formatted to two decimal places for readability.
The code follows Go best practices, including proper error handling and formatting with gofmt.
On Windows CMD, ensure inputs are space-separated without quotes (e.g., 5 3 +).

