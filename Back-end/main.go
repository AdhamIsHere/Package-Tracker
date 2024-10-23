package main

import (
	"Package-Tracker/database"
	"Package-Tracker/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// CORS Middleware
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allowing specific origins (in this case localhost:4200, where Angular is running)
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// If it's a preflight OPTIONS request, respond with OK and return
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Proceed to the next handler
		next.ServeHTTP(w, r)
	})
}

func main() {
	router := mux.NewRouter()

	// user routes
	router.HandleFunc("/register", handlers.RegisterUser).Methods("POST")
	router.HandleFunc("/login", handlers.Login).Methods("POST")
	router.HandleFunc("/logout", handlers.Logout).Methods("POST")
	router.HandleFunc("/order/create", handlers.CreateOrder).Methods("POST")
	router.HandleFunc("/order/myorders", handlers.GetUserOrders).Methods("GET")
	router.HandleFunc("/order/view", handlers.ViewUserOrderDetails).Methods("GET")

	//admin routes
	router.HandleFunc("/order/viewall", handlers.ViewAllOrders).Methods("GET")
	router.HandleFunc("/order/update", handlers.UpdateOrderStatus).Methods("PUT")
	router.HandleFunc("/order/delete", handlers.DeleteOrder).Methods("DELETE")

	// Apply the CORS middleware to the router
	corsRouter := enableCORS(router)

	// Connect to the database
	database.ConnectDB()

	// Start the server on port 8080
	if err := http.ListenAndServe(":8080", corsRouter); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
