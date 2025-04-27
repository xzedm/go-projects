package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
)

// Book represents a book in the bookstore
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// BookStore manages the collection of books
type BookStore struct {
	books  []Book
	nextID int
	mu     sync.Mutex
}

// NewBookStore initializes a new BookStore
func NewBookStore() *BookStore {
	return &BookStore{
		books:  []Book{},
		nextID: 1,
	}
}

// addBook adds a new book and returns its ID
func (bs *BookStore) addBook(book Book) int {
	bs.mu.Lock()
	defer bs.mu.Unlock()
	book.ID = bs.nextID
	bs.nextID++
	bs.books = append(bs.books, book)
	return book.ID
}

// getBooks returns all books
func (bs *BookStore) getBooks() []Book {
	bs.mu.Lock()
	defer bs.mu.Unlock()
	return bs.books
}

// deleteBook removes a book by ID and returns true if found
func (bs *BookStore) deleteBook(id int) bool {
	bs.mu.Lock()
	defer bs.mu.Unlock()
	for i, book := range bs.books {
		if book.ID == id {
			bs.books = append(bs.books[:i], bs.books[i+1:]...)
			return true
		}
	}
	return false
}

func main() {
	store := NewBookStore()
	mux := http.NewServeMux()

	// GET /books
	mux.HandleFunc("GET /books", func(w http.ResponseWriter, r *http.Request) {
		books := store.getBooks()
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(books); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		}
	})

	// POST /books
	mux.HandleFunc("POST /books", func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
			return
		}
		fmt.Printf("Received payload: %s\n", body)
		r.Body = io.NopCloser(bytes.NewReader(body)) // Restore body
		var book Book
		if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
			http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
			return
		}
		if book.Title == "" || book.Author == "" {
			http.Error(w, "Title and Author are required", http.StatusBadRequest)
			return
		}
		id := store.addBook(book)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]int{"id": id})
	})

	// DELETE /books/{id}
	mux.HandleFunc("DELETE /books/", func(w http.ResponseWriter, r *http.Request) {
		parts := strings.Split(r.URL.Path, "/")
		if len(parts) != 3 {
			http.Error(w, "Invalid URL format", http.StatusBadRequest)
			return
		}
		id, err := strconv.Atoi(parts[2])
		if err != nil {
			http.Error(w, "Invalid book ID", http.StatusBadRequest)
			return
		}
		if !store.deleteBook(id) {
			http.Error(w, "Book not found", http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	})

	fmt.Println("Server starting on :8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Fprintf(os.Stderr, "Server failed: %v\n", err)
		os.Exit(1)
	}
}
