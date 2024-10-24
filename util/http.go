package util

import (
	"encoding/json"
	"errors"
	"net/http"

	validator "github.com/go-playground/validator/v10"
)

var validate = validator.New()

func WriteJson(w http.ResponseWriter, status int, v any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, message string) {
	WriteJson(w, status, map[string]string{"error": message})
}

func ReadJson[T any](r *http.Request) (T, error) {
	var payload T
	if r.Body == nil {
		return payload, errors.New("missing request body")
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		return payload, err
	}

	if err := validate.Struct(payload); err != nil {
		return payload, err
	}

	return payload, nil
}
