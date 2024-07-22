package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

// Validate is the validator instance
var Validate = validator.New()

// ParseJSON parses the request body and decodes it into the payload
func ParseJSON(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("request body can not be empty")
	}

	return json.NewDecoder(r.Body).Decode(payload)
}

// WriteJSON writes the response as JSON
func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")

	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

// WriteError writes the error as JSON
func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{
		"error": err.Error(),
	})
}
