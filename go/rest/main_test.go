package main_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	main "github.com/bcdunbar/mixedbag/go/rest/main"
)

func TestGetHandler(t *testing.T) {
	// Create a new request to the GET endpoint
	req, err := http.NewRequest("GET", "/api/get", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new recorder to capture the response
	rr := httptest.NewRecorder()

	// Create a new test server and register the GET endpoint
	handler := http.HandlerFunc(main.GetHandler)
	handler.ServeHTTP(rr, req)

	// Check the status code of the response
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the content type of the response
	if ctype := rr.Header().Get("Content-Type"); ctype != "application/json" {
		t.Errorf("handler returned wrong content type: got %v want %v",
			ctype, "application/json")
	}

	// Decode the response body into a Data object
	var data main.Data
	if err := json.NewDecoder(rr.Body).Decode(&data); err != nil {
		t.Fatal(err)
	}

	// Check the message in the response
	expected := "Hello, World!"
	if data.Message != expected {
		t.Errorf("handler returned unexpected message: got %v want %v",
			data.Message, expected)
	}
}

func TestPostHandler(t *testing.T) {
	// Create a new request to the POST endpoint with a JSON body
	body := strings.NewReader(`{"number": 5}`)
	req, err := http.NewRequest("POST", "/api/post", body)
	if err != nil {
		t.Fatal(err)
	}

	// Set the Content-Type header to application/json
	req.Header.Set("Content-Type", "application/json")

	// Create a new recorder to capture the response
	rr := httptest.NewRecorder()

	// Create a new test server and register the POST endpoint
	handler := http.HandlerFunc(main.PostHandler)
	handler.ServeHTTP(rr, req)

	// Check the status code of the response
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the content type of the response
	if ctype := rr.Header().Get("Content-Type"); ctype != "application/json" {
		t.Errorf("handler returned wrong content type: got %v want %v",
			ctype, "application/json")
	}

	// Decode the response body into a MultiplyResponse object
	var res main.MultiplyResponse
	if err := json.NewDecoder(rr.Body).Decode(&res); err != nil {
		t.Fatal(err)
	}

	// Check the result in the response
	expected := 10
	if res.Result != expected {
		t.Errorf("handler returned unexpected result: got %v want %v",
			res.Result, expected)
	}
}

func TestQueryHandler(t *testing.T) {
	// Create a new request to the QUERY endpoint with query parameters
	req, err := http.NewRequest("GET", "/api/query?a=5&b=10", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new recorder to capture the response
	rr := httptest.NewRecorder()

	// Create a new test server and register the QUERY endpoint
	handler := http.HandlerFunc(main.QueryHandler)
	handler.ServeHTTP(rr, req)

	// Check the status code of the response
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the content type of the response
	if ctype := rr.Header().Get("Content-Type"); ctype != "application/json" {
		t.Errorf("handler returned wrong content type: got %v want %v",
			ctype, "application/json")
	}

	// Decode the response body into a QueryResponse object
	var res main.QueryResponse
	if err := json.NewDecoder(rr.Body).Decode(&res); err != nil {
		t.Fatal(err)
	}

	// Check the result in the response
	expected := 50
	if res.Result != expected {
		t.Errorf("handler returned unexpected result: got %v want %v",
			res.Result, expected)
	}
}
