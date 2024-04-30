package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/topdata-software-gmbh/topdata-package-service/pkg"
	"log"
	"net/http"
)

var config pkg.Config

var port string

func init() {
	flag.StringVar(&port, "port", "8080", "port to run the server on")
}

func main() {
	flag.Parse()

	var err error
	configFile := "config.json5"
	fmt.Printf("Reading config file: %s\n", configFile)
	config, err = pkg.LoadConfig(configFile)
	if err != nil {
		log.Fatalf("Failed to load config: %s", err)
	}

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the TopData Package Service!")
	})

	router.HandleFunc("/repositories", getRepositories)

	fmt.Printf("Server started at http://localhost:%s\n", port)
	fmt.Println("API Endpoints:")
	fmt.Printf("http://localhost:%s/\n", port)
	fmt.Printf("http://localhost:%s/repositories\n", port)
	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}

func getRepositories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(config.Repositories)
}
