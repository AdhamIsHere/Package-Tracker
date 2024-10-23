package handlers

import (
	"Package-Tracker/database"
	"Package-Tracker/models"
	"encoding/json"
	"net/http"
)

func ViewAllOrders(writer http.ResponseWriter, req *http.Request) {
	// Get all orders
	orders := []models.Order{}
	database.DB.Find(&orders)
	json.NewEncoder(writer).Encode(orders)
}

func UpdateOrderStatus(writer http.ResponseWriter, req *http.Request) {
	var order models.Order
	err := json.NewDecoder(req.Body).Decode(&order)
	if err != nil {
		http.Error(writer, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := database.DB.Save(&order).Error; err != nil {
		http.Error(writer, "Could not update order", http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(order)
}

func DeleteOrder(writer http.ResponseWriter, req *http.Request) {
	var order models.Order
	err := json.NewDecoder(req.Body).Decode(&order)
	if err != nil {
		http.Error(writer, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := database.DB.Delete(&order).Error; err != nil {
		http.Error(writer, "Could not delete order", http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(map[string]string{"message": "Order deleted successfully"})
}
