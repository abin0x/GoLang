package main

import (
	"encoding/json"
	"net/http"
)

// senData sends a JSON response with the given data and status code
func senData(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}
