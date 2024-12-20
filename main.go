package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Response struct for JSON response
type Response struct {
	Result float64 `json:"result"`
}

// Handler for addition
func adding(w http.ResponseWriter, r *http.Request) {
	num1, num2 := parseNumbers(r)
	result := num1 + num2
	sendResponse(w, result)
}

// Handler for subtraction
func subtracting(w http.ResponseWriter, r *http.Request) {
	num1, num2 := parseNumbers(r)
	result := num1 - num2
	sendResponse(w, result)
}

// Handler for multiplication
func multiplying(w http.ResponseWriter, r *http.Request) {
	num1, num2 := parseNumbers(r)
	result := num1 * num2
	sendResponse(w, result)
}

// Handler for division
func dividing(w http.ResponseWriter, r *http.Request) {
	num1, num2 := parseNumbers(r)
	if num2 == 0 {
		http.Error(w, "Cannot divide by zero", http.StatusBadRequest)
		return
	}
	result := num1 / num2
	sendResponse(w, result)
}

// Parse numbers from query parameters
func parseNumbers(r *http.Request) (float64, float64) {
	query := r.URL.Query()
	num1, _ := strconv.ParseFloat(query.Get("num1"), 64)
	num2, _ := strconv.ParseFloat(query.Get("num2"), 64)
	return num1, num2
}

// Send JSON response
func sendResponse(w http.ResponseWriter, result float64) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{Result: result})
}

// Handle routes
func handleRequests() {
	http.HandleFunc("/add", adding)
	http.HandleFunc("/subtract", subtracting)
	http.HandleFunc("/multiply", multiplying)
	http.HandleFunc("/divide", dividing)
}

// Main function
func main() {
	fmt.Println("Starting API on port 8080...")
	handleRequests()
	http.ListenAndServe(":8080", nil)
}
