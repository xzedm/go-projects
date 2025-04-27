## URL Shortener (Local)
A Go-based URL shortener service that maps short keys to full URLs. It provides endpoints to create and resolve short URLs, storing mappings in memory.

---

## Features

- POST /shorten: Creates a short URL for a given full URL.
- GET /{short_key}: Redirects to the full URL associated with the short key.
- Generates random 6-character short keys.
- Validates URLs (must start with http:// or https://).
- Thread-safe with mutex for concurrent access.

---

## Prerequisites

- Go 1.16 or higher

---

## Installation

- Clone or download the project.
- Navigate to the project directory.
- Run the server:go run urlshortener.go

---

## Usage
- The server runs on http://localhost:8080. Use a tool like curl or Postman to interact with the API.
Endpoints

- POST /shorten

- Creates a short URL.
- Payload: JSON with url field.
- Example:curl -X POST -H "Content-Type: application/json" -d '{"url":"https://www.example.com"}' http://localhost:8080/shorten

Windows CMD:
- curl -X POST -H "Content-Type: application/json" -d "{\"url\":\"https://www.example.com\"}" http://localhost:8080/shorten

- Response:{"short_url":"http://localhost:8080/abc123"}




GET /{short_key}

Redirects to the full URL.
Example:curl -L http://localhost:8080/abc123

Response: Redirects to https://www.example.com.



File Structure

urlshortener.go: Main source code
README.md: This file

Notes

URLs are stored in memory and reset when the server restarts. For persistence, you could integrate BoltDB.
Short keys are 6-character alphanumeric strings, randomly generated.
The API validates URLs and returns appropriate error messages for invalid inputs.
The code follows Go best practices, including proper error handling and formatting with gofmt.
On Windows CMD, use escaped quotes for JSON payloads:curl -X POST -H "Content-Type: application/json" -d "{\"url\":\"https://www.example.com\"}" http://localhost:8080/shorten



