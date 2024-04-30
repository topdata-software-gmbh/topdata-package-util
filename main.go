package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	config, err := LoadConfig("config.json")
	if err != nil {
		log.Fatalf("Failed to load config: %s", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the TopData Package Service!")
	})

	http.HandleFunc("/repositories", getRepositories)

	fmt.Println("Server started at http://localhost:8080")
	fmt.Println("API Endpoints:")
	fmt.Println("http://localhost:8080/")
	fmt.Println("http://localhost:8080/repositories")
	http.ListenAndServe(":8080", nil)
}
func getRepositories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(config.Repositories)
}
