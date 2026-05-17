package main

import (
	"fmt"
	"log"
	"net/http"

	"health-crud/app/handlers"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	// Serve static files via a fallback handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.ServeFile(w, r, "./static/index.html")
			return
		}
		http.ServeFile(w, r, "./static"+r.URL.Path)
	})

	// Register routes
	http.HandleFunc("/create", handlers.CreatePatientHandler)
	http.HandleFunc("/search", handlers.SearchPatientHandler)
	http.HandleFunc("/update", handlers.UpdatePatientHandler)
	http.HandleFunc("/delete", handlers.DeletePatientHandler)

	fmt.Println("Server is running on http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
