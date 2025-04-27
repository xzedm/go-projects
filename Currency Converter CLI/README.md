Currency Converter CLI
A Go command-line tool that converts an amount from one currency to another using the ExchangeRate-API. It accepts an amount and currency codes as input and displays the converted amount.
Features

Converts currencies using real-time exchange rates
Accepts input in the format <amount> <from_currency> <to_currency>
Handles errors for invalid currency codes, negative amounts, and network issues
Supports uppercase and lowercase currency codes (normalized to uppercase)

Prerequisites

Go 1.16 or higher
An API key from ExchangeRate-API
Sign up for a free account to obtain a key (free tier available).



Installation

Clone or download the project.
Navigate to the project directory.
Replace YOUR_API_KEY in currency_converter.go with your ExchangeRate-API key.
Run the program:go run currency_converter.go



Usage
Run the program and enter an amount, source currency, and target currency:
Currency Converter CLI
Enter amount, from currency, and to currency (e.g., '100 USD EUR'). Type 'exit' to quit.
> 100 USD EUR
100.00 USD = 94.50 EUR
> 50 EUR GBP
50.00 EUR = 42.30 GBP
> invalid
Usage: <amount> <from_currency> <to_currency> (e.g., '100 USD EUR')
> 100 XYZ EUR
Error: API error: invalid-base-currency
> exit
Goodbye!

File Structure

currency_converter.go: Main source code
README.md: This file

Notes

Obtain an API key from ExchangeRate-API and update the apiKey constant in currency_converter.go.
The program normalizes currency codes to uppercase for consistency.
Errors (e.g., invalid currency codes, network issues) are displayed with clear messages.
The code follows Go best practices, including proper error handling and formatting with gofmt.
To extend the program, you could add support for a list of supported currencies or caching of exchange rates.

