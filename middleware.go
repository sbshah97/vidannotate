package main

import (
	"encoding/json"
	"net/http"
)

// Middleware to validate API key in header
func validateAPIKeyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("API-Key")
		if apiKey != "valid_api_key" {
			json.NewEncoder(w).Encode(Error{Message: "Invalid API key"})
			return
		}
		next.ServeHTTP(w, r)
	})
}

// Middleware to validate JWT token
func validateJWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jwt := r.Header.Get("Authorization")
		if jwt != "valid_jwt" {
			json.NewEncoder(w).Encode(Error{Message: "Invalid JWT token"})
			return
		}
		next.ServeHTTP(w, r)
	})
}
