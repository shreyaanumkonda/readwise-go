// this is utility file is used to write json to the response writer
package main

import (
	"encoding/json"
	"net/http"
)

// WriteJSON writes JSON response with status code
func WriteJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}
