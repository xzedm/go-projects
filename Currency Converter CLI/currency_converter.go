package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// ExchangeRateResponse represents the API response structure
type ExchangeRateResponse struct {
	Result          string             `json:"result"`
	BaseCode        string             `json:"base_code"`
	ConversionRates map[string]float64 `json:"conversion_rates"`
	ErrorType       string             `json:"error-type,omitempty"`
}

// fetchExchangeRate fetches the exchange rate for a base currency
func fetchExchangeRate(baseCurrency string, apiKey string) (map[string]float64, error) {
	url := fmt.Sprintf("https://v6.exchangerate-api.com/v6/%s/latest/%s", apiKey, baseCurrency)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch exchange rates: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP error %d: %s", resp.StatusCode, resp.Status)
	}

	var result ExchangeRateResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	if result.Result != "success" {
		return nil, fmt.Errorf("API error: %s", result.ErrorType)
	}

	return result.ConversionRates, nil
}

// convertCurrency converts an amount from one currency to another
func convertCurrency(amount float64, from, to string, apiKey string) (float64, error) {
	rates, err := fetchExchangeRate(from, apiKey)
	if err != nil {
		return 0, err
	}

	rate, exists := rates[to]
	if !exists {
		return 0, fmt.Errorf("invalid target currency code: %s", to)
	}

	return amount * rate, nil
}

func main() {
	const apiKey = "4165fbf75802179ecb67c7f3" // Replace with your ExchangeRate-API key
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Currency Converter CLI")
	fmt.Println("Enter amount, from currency, and to currency (e.g., '100 USD EUR'). Type 'exit' to quit.")

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

		parts := strings.Fields(input)
		if len(parts) != 3 {
			fmt.Println("Usage: <amount> <from_currency> <to_currency> (e.g., '100 USD EUR')")
			continue
		}

		amount, err := strconv.ParseFloat(parts[0], 64)
		if err != nil {
			fmt.Println("Invalid amount:", parts[0])
			continue
		}
		if amount < 0 {
			fmt.Println("Amount must be non-negative")
			continue
		}

		fromCurrency := strings.ToUpper(parts[1])
		toCurrency := strings.ToUpper(parts[2])

		convertedAmount, err := convertCurrency(amount, fromCurrency, toCurrency, apiKey)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		fmt.Printf("%.2f %s = %.2f %s\n", amount, fromCurrency, convertedAmount, toCurrency)
	}
}
