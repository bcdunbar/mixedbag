package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

// Data represents a JSON response with a message
type Data struct {
	Message string `json:"message"`
}

// MultiplyRequest represents a JSON request with a number to multiply
type MultiplyRequest struct {
	Number int `json:"number"`
}

// MultiplyResponse represents a JSON response with the result of multiplying a number
type MultiplyResponse struct {
	Result int `json:"result"`
}

// QueryRequest represents a JSON request with two numbers to multiply
type QueryRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}

// QueryResponse represents a JSON response with the result of multiplying two numbers
type QueryResponse struct {
	Result int `json:"result"`
}

func main() {
	// Create a new router
	router := http.NewServeMux()

	// Register the GET and POST endpoints
	router.HandleFunc("/api/get", GetHandler)
	router.HandleFunc("/api/post", PostHandler)

	// Register the QUERY endpoint
	router.HandleFunc("/api/query", QueryHandler)

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", router))
}

// getHandler handles GET requests at /api/get
func GetHandler(w http.ResponseWriter, r *http.Request) {
	// Create a new Data object with a message
	data := Data{Message: "Hello, World!"}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Encode the Data object as JSON and write it to the response
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Log the response message
	log.Printf("GET Response: %s", data.Message)
}

// postHandler handles POST requests at /api/post
func PostHandler(w http.ResponseWriter, r *http.Request) {
	// Decode the request body into a MultiplyRequest object
	var req MultiplyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Multiply the input number by 2 and create a new MultiplyResponse object
	res := MultiplyResponse{Result: req.Number * 2}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Encode the MultiplyResponse object as JSON and write it to the response
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Log the response message
	log.Printf("POST Response: %d", res.Result)
}

// queryHandler handles GET requests at /api/query
func QueryHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the query parameters a and b as integers
	aStr := r.URL.Query().Get("a")
	bStr := r.URL.Query().Get("b")

	a, err := strconv.Atoi(aStr)
	if err != nil {
		http.Error(w, `{"error": "Invalid value for parameter 'a'"}`, http.StatusBadRequest)
		return
	}

	b, err := strconv.Atoi(bStr)
	if err != nil {
		http.Error(w, `{"error": "Invalid value for parameter 'b'"}`, http.StatusBadRequest)
		return
	}

	// Multiply the input numbers and create a new QueryResponse object
	res := QueryResponse{Result: a * b}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Encode the QueryResponse object as JSON and write it to the response
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, `{"error": "Internal server error"}`, http.StatusInternalServerError)
		return
	}

	// Log the response message
	log.Printf("QUERY Response: %d", res.Result)
}
