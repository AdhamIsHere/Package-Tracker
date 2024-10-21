package handlers

import (
	"Package-Tracker/database"
	"Package-Tracker/models"
	"encoding/json"
	"errors"
	"github.com/dgrijalva/jwt-go"
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
	if order.PackageDetails == "" {
		http.Error(writer, "Package details is required", http.StatusBadRequest)
		return
	}
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
	order.UserID = id

	// --------------------------------------------------
	// save the order
	if err := database.DB.Create(&order).Error; err != nil {
		http.Error(writer, "Could not create order", http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(order)
}
