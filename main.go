package main

import (
	"Package-Tracker/database"
	"Package-Tracker/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	// Public routes
	router.HandleFunc("/register", handlers.RegisterUser).Methods("POST")
	router.HandleFunc("/login", handlers.Login).Methods("POST")

	// Protected routes (in the future)
	// router.Handle("/orders", middleware.Auth(http.HandlerFunc(handlers.GetOrders))).Methods("GET")

	// Connect to the database
	database.ConnectDB()

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
