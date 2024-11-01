package handlers

import (
	"Package-Tracker/database"
	"Package-Tracker/models"
	"encoding/json"
	"net/http"
)

func ViewAssignedOrders(writer http.ResponseWriter, req *http.Request) {
	// Get all orders
	orders := []models.Order{}
	id, err := GetIDFromToken(req)
	if err != nil {
		http.Error(writer, "Invalid token", http.StatusBadRequest)
		return
	}
	database.DB.Where("courier_id = ?", id).Preload("Items").Find(&orders)
	json.NewEncoder(writer).Encode(orders)
}
