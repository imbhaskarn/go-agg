package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func handlerError(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 400, "Internal Server Error")
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	if code > 499 {
		log.Println("Server error:", message)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, code, errorResponse{Error: message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload any) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON : %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}