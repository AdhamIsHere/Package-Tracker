package handlers

import (
	"Package-Tracker/database"
	"Package-Tracker/models"
	"encoding/json"
	"net/http"
	"strconv"
)

func IsCourier(req *http.Request) bool {
	id, err := GetIDFromToken(req)
	if err != nil {
		panic(err)
	}
	role, err := GetRoleFromID(id)
	if err != nil {
		panic(err)
	}
	if role == "courier" {
		return true
	}
	return false
}

func ViewAssignedOrders(writer http.ResponseWriter, req *http.Request) {
	if !IsCourier(req) {
		http.Error(writer, "Unauthorized", http.StatusUnauthorized)
		return
	}
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

func UpdateOrderStatus(writer http.ResponseWriter, req *http.Request) {

	if !IsCourier(req) {
		http.Error(writer, "Unauthorized", http.StatusUnauthorized)
		return
	}

	status := req.URL.Query().Get("status")

	orderID := req.URL.Query().Get("id")
	if orderID == "" {
		http.Error(writer, "Order ID is required", http.StatusBadRequest)
		return
	}
	if status == "" {
		http.Error(writer, "Status is required", http.StatusBadRequest)
		return
	}

	oid, err := strconv.Atoi(orderID)
	if err != nil {
		http.Error(writer, "Invalid order ID", http.StatusBadRequest)
		return
	}

	var order models.Order
	if err := database.DB.First(&order, oid).Error; err != nil {
		http.Error(writer, "Order not found", http.StatusNotFound)
		return
	}

	order.Status = status
	if err := database.DB.Save(&order).Error; err != nil {
		http.Error(writer, "Could not update order status", http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	msg := "Order no." + string(order.ID) + " status updated to " + status
	json.NewEncoder(writer).Encode(map[string]string{"message": msg})
}
