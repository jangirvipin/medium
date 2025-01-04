package main

import (
	"fmt"
	"log"
	"net/http"

	"medium/server/config"
	"medium/server/handlers"
	"medium/server/routes"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables directly")
	}

	fmt.Printf("hello world\n")
	db := config.ConnectDatabase()
	handlers.SetDB(db) // Pass the DB instance to handlers

	// Register routes
	routes.RegisterRoutes()

	// Start the HTTP server
	port:=":3000"
	fmt.Printf("Server running on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
