package handlers

import (
	"Package-Tracker/database"
	"Package-Tracker/models"
	"encoding/json"
	"net/http"
	"strconv"
)

func ViewAllOrders(writer http.ResponseWriter, req *http.Request) {
	// Get all orders
	orders := []models.Order{}
	database.DB.Preload("Items").Find(&orders)
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
	json.NewEncoder(writer).Encode(map[string]string{"message": "Order updated successfully"})
}

func DeleteOrder(writer http.ResponseWriter, req *http.Request) {
	// Get the order ID from the request URL
	orderID := req.URL.Query().Get("id")
	if orderID == "" {
		http.Error(writer, "Order ID is required", http.StatusBadRequest)
		return
	}

	// Convert orderID to integer
	id, err := strconv.Atoi(orderID)
	if err != nil {
		http.Error(writer, "Invalid order ID", http.StatusBadRequest)
		return
	}

	// Start a transaction
	tx := database.DB.Begin()
	if tx.Error != nil {
		http.Error(writer, "Could not start transaction", http.StatusInternalServerError)
		return
	}

	// Delete associated items
	if err := tx.Where("order_id = ?", id).Delete(&models.Item{}).Error; err != nil {
		tx.Rollback()
		http.Error(writer, "Could not delete order items", http.StatusInternalServerError)
		return
	}

	// Delete the order by ID
	if err := tx.Delete(&models.Order{}, id).Error; err != nil {
		tx.Rollback()
		http.Error(writer, "Could not delete order", http.StatusInternalServerError)
		return
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		http.Error(writer, "Could not commit transaction", http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(map[string]string{"message": "Order deleted successfully"})
}
