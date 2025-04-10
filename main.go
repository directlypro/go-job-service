package main

import (
	//"encoding/json"
	"fmt"
	"go-job-service/router"
	"log"
	"net/http"
	"os"
	"time"

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

	appRouter := router.NewServiceRouter()

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      appRouter,
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
