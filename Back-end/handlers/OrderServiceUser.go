package handlers

import (
	"Package-Tracker/database"
	"Package-Tracker/models"
	"encoding/json"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
)

// function to get id from token
func GetIDFromToken(req *http.Request) (int, error) {
	// Get the user ID from the token
	cookie, _ := req.Cookie("token")
	claims, err := ParseToken(cookie)
	if err != nil {
		return 0, err
	}
	id := claims.ID
	return id, nil
}

func ParseToken(cookie *http.Cookie) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(cookie.Value, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("Invalid token")
	}
	return claims, nil
}

func CreateOrder(writer http.ResponseWriter, req *http.Request) {
	var order models.Order
	err := json.NewDecoder(req.Body).Decode(&order)
	if err != nil {
		http.Error(writer, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if order.DropOffLocation == "" {
		http.Error(writer, "Drop off location is required", http.StatusBadRequest)
		return
	}
	if order.PickupLocation == "" {
		http.Error(writer, "Pickup location is required", http.StatusBadRequest)
		return
	}

	if len(order.Items) == 0 {
		http.Error(writer, "at least 1 item is needed", http.StatusBadRequest)
		return
	}
	// ------------------------------------
	for i := 0; i < len(order.Items); i++ {
		// check if item exists in items table
		var item models.Item
		if err := database.DB.Where("id = ?", order.Items[i].ID).First(&item).Error; err != nil {
			http.Error(writer, "Item(s) does not exist", http.StatusBadRequest)
			return
		}
	}

	// ------------------------------------
	if order.DeliveryTime.IsZero() {
		http.Error(writer, "Delivery time is required", http.StatusBadRequest)
		return
	}

	// --------------------------------------------------

	// get id from token
	id, err := GetIDFromToken(req)
	if err != nil {
		http.Error(writer, "Could not get user ID "+err.Error(), http.StatusInternalServerError)
		return
	}
	order.SellerID = id

	// --------------------------------------------------
	// save the order
	if err := database.DB.Create(&order).Error; err != nil {
		http.Error(writer, "Could not create order", http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(order)
}

func GetUserOrders(writer http.ResponseWriter, req *http.Request) {
	id, err := GetIDFromToken(req)
	if err != nil {
		http.Error(writer, "Could not get user ID "+err.Error(), http.StatusInternalServerError)
		return
	}

	var orders []models.Order
	if err := database.DB.Where("user_id = ?", id).Find(&orders).Error; err != nil {
		http.Error(writer, "Could not get orders", http.StatusInternalServerError)
		return
	}

	type OrderSummary struct {
		ID     int64  `json:"order_id"`
		UserID int    `json:"user_id"`
		Status string `json:"status"`
	}

	var orderSummaries []OrderSummary
	for _, order := range orders {
		orderSummaries = append(orderSummaries, OrderSummary{
			ID:     int64(order.ID),
			UserID: order.SellerID,
			Status: order.Status,
		})
	}

	json.NewEncoder(writer).Encode(orderSummaries)
}

func ViewUserOrderDetails(writer http.ResponseWriter, req *http.Request) {
	id, err := GetIDFromToken(req)
	if err != nil {
		http.Error(writer, "Could not get user ID "+err.Error(), http.StatusInternalServerError)
		log.Printf("Error getting user ID from token: %v", err)
		return
	}

	var order models.Order
	orderID := req.URL.Query().Get("id")
	if orderID == "" {
		http.Error(writer, "Order ID is required", http.StatusBadRequest)
		return
	}

	if err := database.DB.Where("id = ? AND user_id = ?", orderID, id).Preload("Items").First(&order).Error; err != nil {
		http.Error(writer, "Could not get order", http.StatusInternalServerError)
		log.Printf("Error fetching order: %v", err)
		return
	}

	json.NewEncoder(writer).Encode(order)
}
