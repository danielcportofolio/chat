package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/danielcportofolio/chat/handlers"
	"github.com/danielcportofolio/chat/repositories"
	"github.com/gorilla/mux"
)

func main() {
	// Database connection string
	dbConfig := "postgresql://username:password@host:port/database_name?sslmode=disable"

	// Open a connection to the database
	db, err := sql.Open("pgx", dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Test the database connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}

	// Create the user repository
	userRepo := repositories.NewUserRepository(db)

	// Create the user handler
	userHandler := handlers.NewUserHandler(userRepo)

	// Create a new router
	router := mux.NewRouter()

	// Define the user endpoint
	router.HandleFunc("/users", userHandler.CreateUser).Methods(http.MethodPost)

	// Start the HTTP server
	port := 8080
	log.Printf("Server started on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
