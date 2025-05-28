package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	
)




func main() {
	godotenv.Load(".env")
	port := os.Getenv("PORT")

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		ExposedHeaders:   []string{"Content-Length", "Authorization"},
	}))

	if port == "" {
		port = "8080" // Default port
	}

	v1Router := chi.NewRouter()
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to the API!")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"message":"Welcome to the API!"}`)
	})
	v1Router.Get("/healthz", readinessHandler)
	v1Router.Get("/err", handlerError)
	router.Mount("/v1", v1Router)

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
