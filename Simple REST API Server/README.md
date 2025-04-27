Simple REST API Server (Bookstore)
A Go RESTful API for a mock Bookstore system, built using the net/http package. It supports listing, adding, and deleting books via HTTP endpoints.
Features

GET /books: List all books
POST /books: Add a new book
DELETE /books/{id}: Delete a book by its ID
Stores books in memory (slice)
Returns JSON responses
Thread-safe with mutex for concurrent access

Prerequisites

Go 1.16 or higher

Installation

Clone or download the project.
Navigate to the project directory.
Run the server:go run bookstore.go



Usage
The server runs on http://localhost:8080. Use a tool like curl or Postman to interact with the API.
Endpoints

GET /books

Lists all books.
Example:curl http://localhost:8080/books

Response:[
  {"id": 1, "title": "The Go Programming Language", "author": "Donovan"},
]




POST /books

Adds a new book. Send JSON with title and author.
Example:curl -v -X POST -H "Content-Type: application/json" -d "{\"title\":\"The Go Programming Language\",\"author\":\"Alan Donovan\"}" http://localhost:8080/books


Response:{"id": 1}




DELETE /books/{id}

Deletes a book by ID.
Example:curl -X DELETE http://localhost:8080/books/1

Response: (204 No Content on success, 404 if not found)



File Structure

bookstore.go: Main source code
README.md: This file

Notes

Books are stored in memory and reset when the server restarts. For persistence, you could add a database like BoltDB.
The API is thread-safe using a mutex to protect the book slice.
Errors (e.g., invalid JSON, missing fields, non-existent IDs) return appropriate HTTP status codes and messages.
The code follows Go best practices, including proper error handling and formatting with gofmt.

