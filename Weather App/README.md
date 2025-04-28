# Weather App CLI

A simple Go command-line tool that fetches and displays real-time weather information for any city using the OpenWeatherMap API.

---

## Features

- Fetches live weather data for a specified city.
- Displays temperature in Celsius and a short weather description.
- Handles errors like invalid city names, network issues, and API failures.
- Supports an interactive CLI with an exit command.

---

## Prerequisites

- Go 1.16 or higher
- An API key from OpenWeatherMap

---

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/go-projects.git
   cd weather-app

2. Run the application:
```bash
go run weather.go
```

## Usage
After running the program, enter a city name to retrieve the current weather information.

Example:
Weather App CLI
Enter a city name to get the current weather, or 'exit' to quit.

> London
Weather in London:
Temperature: 15.2°C
Description: Scattered Clouds

> InvalidCity
Error: API error: city not found or invalid request

> exit
Goodbye!
