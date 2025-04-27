package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"sync"
)

// URLStore manages the mapping of short keys to full URLs
type URLStore struct {
	mu   sync.Mutex
	urls map[string]string
}

// NewURLStore initializes a new URLStore
func NewURLStore() *URLStore {
	return &URLStore{
		urls: make(map[string]string),
	}
}

// AddURL adds a URL with a generated short key and returns the key
func (s *URLStore) AddURL(fullURL string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Validate URL
	if !isValidURL(fullURL) {
		return "", fmt.Errorf("invalid URL: %s", fullURL)
	}

	// Generate unique short key
	shortKey := generateShortKey()
	for _, exists := s.urls[shortKey]; exists; {
		shortKey = generateShortKey()
	}

	s.urls[shortKey] = fullURL
	return shortKey, nil
}

// GetURL retrieves the full URL for a short key
func (s *URLStore) GetURL(shortKey string) (string, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	url, exists := s.urls[shortKey]
	return url, exists
}

// isValidURL checks if the URL starts with http:// or https://
func isValidURL(url string) bool {
	return strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")
}

// generateShortKey creates a random 6-character key
func generateShortKey() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const keyLength = 6
	b := make([]byte, keyLength)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func main() {
	store := NewURLStore()
	mux := http.NewServeMux()

	// POST /shorten
	mux.HandleFunc("POST /shorten", func(w http.ResponseWriter, r *http.Request) {
		var payload struct {
			URL string `json:"url"`
		}
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
			return
		}
		if payload.URL == "" {
			http.Error(w, "URL is required", http.StatusBadRequest)
			return
		}

		shortKey, err := store.AddURL(payload.URL)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response := map[string]string{"short_url": fmt.Sprintf("http://localhost:8080/%s", shortKey)}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	})

	// GET /{short_key}
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		shortKey := strings.TrimPrefix(r.URL.Path, "/")
		if shortKey == "" {
			http.Error(w, "Short key is required", http.StatusBadRequest)
			return
		}

		fullURL, exists := store.GetURL(shortKey)
		if !exists {
			http.Error(w, "Short URL not found", http.StatusNotFound)
			return
		}

		http.Redirect(w, r, fullURL, http.StatusMovedPermanently)
	})

	fmt.Println("URL Shortener starting on :8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Fprintf(os.Stderr, "Server failed: %v\n", err)
		os.Exit(1)
	}
}
