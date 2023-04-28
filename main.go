package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JwtSecretKey should be replaced with a strong secret key in a production environment.
// This secret key should be stored in a secrets vault (e.g. Conjur).
const (
	JwtSecretKey = "your-secret-key-here" // Replace this with a strong secret key
)

// In-memory data structure to store items
type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var items = []Item{} // The in-memory items list is shared between all requests
var idCounter int
var itemsMutex = &sync.Mutex{} // Mutex to protect the items list

func main() {
	// Define the endpoints and their handlers
	http.HandleFunc("/items", jwtAuthMiddleware(itemsHandler))
	http.HandleFunc("/login", loginHandler)

	// Start the server
	log.Println("Starting server on :9292")
	log.Fatal(http.ListenAndServe(":9292", nil))
}

func itemsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getItems(w, r)
	case http.MethodPost:
		addItem(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getItems(w http.ResponseWriter, r *http.Request) {
	itemsMutex.Lock()
	defer itemsMutex.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func addItem(w http.ResponseWriter, r *http.Request) {
	var item Item

	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	itemsMutex.Lock()
	defer itemsMutex.Unlock()

	idCounter++
	item.ID = idCounter
	items = append(items, item)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

// Middleware to verify the JWT token in the Authorization header of incoming requests
func jwtAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")
		if authorizationHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		bearerToken := strings.Split(authorizationHeader, " ")
		if len(bearerToken) != 2 {
			http.Error(w, "Invalid token format", http.StatusBadRequest)
			return
		}

		token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(JwtSecretKey), nil
		})

		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}

// Create and return a JWT token when valid credentials are provided
func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	// Replace this with your actual user authentication logic
	if username != "testuser" || password != "testpassword" {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(JwtSecretKey))
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}
