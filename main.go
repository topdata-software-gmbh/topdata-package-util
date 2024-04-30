package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Config struct {
	Repositories []string
}

func LoadConfig(filename string) (Config, error) {
	var config Config
	file, err := os.Open(filename)
	if err != nil {
		return config, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return config, err
	}
	return config, nil
}

var config Config

func main() {
	var err error
	configFile := "config.json"
	fmt.Printf("Reading config file: %s\n", configFile)
	config, err = LoadConfig(configFile)
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
