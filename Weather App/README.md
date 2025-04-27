Weather App CLI
A Go command-line tool that fetches and displays the current weather for a city using the OpenWeatherMap API. It accepts a city name as input and shows the temperature and weather description.
Features

Fetches real-time weather data for a specified city.
Displays temperature in Celsius and weather description.
Handles errors for invalid city names, network issues, and API failures.
Supports an interactive CLI with an exit command.

Prerequisites

Go 1.16 or higher
An API key from OpenWeatherMap




Installation

Clone or download the project.
Navigate to the project directory.
Run the program:go run weather.go



Usage
Run the program and enter a city name to get the current weather:
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

File Structure

weather.go: Main source code
README.md: This file

Notes

Obtain an API key from OpenWeatherMap and update the apiKey constant in weather.go.
The program uses metric units (Celsius) for temperature.
Errors (e.g., invalid city names, network issues) are displayed with clear messages.
The code follows Go best practices, including proper error handling and formatting with gofmt.
On Windows CMD, enter city names without spaces or use quotes for multi-word cities (e.g., "New York").

