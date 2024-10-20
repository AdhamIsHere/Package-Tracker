package handlers

import (
	"Package-Tracker/database"
	"Package-Tracker/models"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"strconv"
)

func GetIDFromToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil // Replace with your actual secret
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id := claims["id"].(string)
		return id, nil
	}

	return "", fmt.Errorf("invalid token")
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
	// Get the user ID from the token
	cookie, _ := req.Cookie("token")
	num, err := GetIDFromToken(cookie.Value)
	if err != nil {
		http.Error(writer, "Could not get user ID "+err.Error(), http.StatusInternalServerError)
		return
	}
	id, _ := strconv.Atoi(num)
	order.UserID = uint(id)
	// save the order
	if err := database.DB.Create(&order).Error; err != nil {
		http.Error(writer, "Could not create order", http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(order)

}
