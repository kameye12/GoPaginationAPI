package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// User represents a user object
type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
	Address   string `json:"address"`
}

// Dummy data
var users = []User{
	{"John", "Doe", "john@example.com", 30, "123 Elm St"},
	{"Jane", "Doe", "jane@example.com", 25, "456 Oak St"},
	{"Alice", "Johnson", "alice@example.com", 28, "789 Pine St"},
	{"Bob", "Smith", "bob@example.com", 32, "101 Maple St"},
	{"Carol", "White", "carol@example.com", 27, "102 Birch St"},
	{"Dave", "Black", "dave@example.com", 35, "103 Cedar St"},
	{"Eve", "Brown", "eve@example.com", 26, "104 Walnut St"},
	{"Frank", "Green", "frank@example.com", 29, "105 Spruce St"},
	{"Grace", "Blue", "grace@example.com", 24, "106 Aspen St"},
	{"Hank", "Yellow", "hank@example.com", 31, "107 Willow St"},
	// Add more dummy data as needed
}

// getUsers handles the request to get users with pagination
func getUsers(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page < 1 {
		page = 1
	}

	const pageSize = 10
	start := (page - 1) * pageSize
	end := start + pageSize

	if start > len(users) {
		start = len(users)
	}
	if end > len(users) {
		end = len(users)
	}

	paginatedUsers := users[start:end]
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(paginatedUsers)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", getUsers).Methods("GET")

	log.Println("Server started at :8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
