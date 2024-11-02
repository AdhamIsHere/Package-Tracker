package main

import (
	"Package-Tracker/database"
	"Package-Tracker/handlers"
	"Package-Tracker/models"
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

	// courier routes
	router.HandleFunc("/order/assigned", handlers.ViewAssignedOrders).Methods("GET")

	//admin routes
	router.HandleFunc("/order/viewall", handlers.ViewAllOrders).Methods("GET")
	router.HandleFunc("/order/assign", handlers.AssignOrder).Methods("PUT")
	router.HandleFunc("/order/delete", handlers.DeleteOrder).Methods("DELETE")
	router.HandleFunc("/order/updatestatus", handlers.UpdateOrderStatus).Methods("PUT")

	// Apply the CORS middleware to the router
	corsRouter := enableCORS(router)

	// Connect to the database
	database.ConnectDB()
	items := []models.Item{
		{ID: "1", Name: "Laptop", Quantity: 10},
		{ID: "2", Name: "Phone", Quantity: 20},
		{ID: "3", Name: "Headphones", Quantity: 15},
		{ID: "4", Name: "Keyboard", Quantity: 25},
		{ID: "5", Name: "Mouse", Quantity: 30},
	}

	// Check if the item already exists in the database
	var existingItem models.Item
	result := database.DB.Where("id = ?", items[0].ID).First(&existingItem)

	if result.RowsAffected == 0 {
		// Item does not exist, insert it
		if err := database.DB.Create(&items).Error; err != nil {
			log.Printf("Failed to insert items : %v", err)
		} else {
			log.Println("Inserted items")
		}
	} else {
		log.Println("Items already loaded in DB")
	}

	// Start the server on port 8080
	if err := http.ListenAndServe(":8080", corsRouter); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
