package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("CLI Calculator (supports +, -, *, /, sqrt, pow)")
	fmt.Println("Enter two numbers (or one for sqrt) and an operation, separated by spaces (e.g., '5 3 +' or '16 sqrt'). Type 'exit' to quit.")

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
			break
		}

		parts := strings.Fields(input)
		if len(parts) < 2 || len(parts) > 3 {
			fmt.Println("Invalid input. Use format: 'num1 num2 op' or 'num op' for sqrt.")
			continue
		}

		var num1, num2 float64
		var op string

		if parts[len(parts)-1] == "sqrt" {
			if len(parts) != 2 {
				fmt.Println("Sqrt requires one number: 'num sqrt'")
				continue
			}
			num1, err = strconv.ParseFloat(parts[0], 64)
			if err != nil {
				fmt.Println("Invalid number:", parts[0])
				continue
			}
			op = "sqrt"
		} else {
			if len(parts) != 3 {
				fmt.Println("Expected format: 'num1 num2 op'")
				continue
			}
			num1, err = strconv.ParseFloat(parts[0], 64)
			if err != nil {
				fmt.Println("Invalid first number:", parts[0])
				continue
			}
			num2, err = strconv.ParseFloat(parts[1], 64)
			if err != nil {
				fmt.Println("Invalid second number:", parts[1])
				continue
			}
			op = parts[2]
		}

		var result float64
		switch op {
		case "+":
			result = num1 + num2
		case "-":
			result = num1 - num2
		case "*":
			result = num1 * num2
		case "/":
			if num2 == 0 {
				fmt.Println("Error: Division by zero is not allowed.")
				continue
			}
			result = num1 / num2
		case "sqrt":
			if num1 < 0 {
				fmt.Println("Error: Square root of negative number is not allowed.")
				continue
			}
			result = math.Sqrt(num1)
		case "pow":
			result = math.Pow(num1, num2)
		default:
			fmt.Println("Invalid operation. Supported: +, -, *, /, sqrt, pow")
			continue
		}

		fmt.Printf("Result: %.2f\n", result)
	}
}
