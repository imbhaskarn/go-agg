package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	port := os.Getenv("PORT")
	router := chi.NewRouter()
	if port == "" {
		port = "8080" // Default port
	}

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	fmt.Printf("Server is running on port %s\nhttp://localhost:%s\n", port, port)
	err := srv.ListenAndServe()
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		return
	}

}
