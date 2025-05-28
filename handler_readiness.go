package main

import (
	"net/http"
)

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	println("Readiness check received")
	respondWithJSON(w, 200, struct {
	}{})
}
