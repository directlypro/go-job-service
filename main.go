package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Starting Go Job Service")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed, could not load .env file")
	}

	// Port configuration
	port := os.Getenv("PORT")
	if port == "" {
		port = "8083"
		log.Printf("Defaulting to port %s", port)
	}

	r := mux.NewRouter()

	r.HandleFunc("/healthz", healthCheckHandler).Methods(http.MethodGet)

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	fmt.Printf("Server listening on port: %s\n", port)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string]string{"status": "ok"}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Printf("Failed writing health check response: %v", err)
	}
}

//func hello() {
//	fmt.Println("Hello wor")
//}
