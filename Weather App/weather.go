package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

// WeatherResponse represents the OpenWeatherMap API response
type WeatherResponse struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
	Name string `json:"name"`
	Cod  int    `json:"cod"`
}

// fetchWeather fetches weather data for a city using OpenWeatherMap API
func fetchWeather(city, apiKey string) (*WeatherResponse, error) {
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch weather data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP error %d: %s", resp.StatusCode, resp.Status)
	}

	var weather WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	if weather.Cod != 200 {
		return nil, fmt.Errorf("API error: city not found or invalid request")
	}

	return &weather, nil
}

func main() {
	const apiKey = "7128b60745a7dc3d585def9822b203c9" // Replace with your OpenWeatherMap API key
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Weather App CLI")
	fmt.Println("Enter a city name to get the current weather, or 'exit' to quit.")

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
		if input == "" {
			fmt.Println("Please enter a city name.")
			continue
		}

		weather, err := fetchWeather(input, apiKey)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		// Display weather information
		fmt.Printf("Weather in %s:\n", weather.Name)
		fmt.Printf("Temperature: %.1f°C\n", weather.Main.Temp)
		if len(weather.Weather) > 0 {
			fmt.Printf("Description: %s\n", strings.Title(weather.Weather[0].Description))
		} else {
			fmt.Println("Description: N/A")
		}
	}
}
