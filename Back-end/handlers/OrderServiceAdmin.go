package handlers

import (
	"Package-Tracker/database"
	"Package-Tracker/models"
	"encoding/json"
	"net/http"
	"strconv"
)

func IsAdmin(req *http.Request) bool {
	id, err := GetIDFromToken(req)
	if err != nil {
		panic(err)
	}
	role, err := GetRoleFromID(id)
	if err != nil {
		panic(err)
	}
	if role == "admin" {
		return true
	}
	return false
}

func ViewAllOrders(writer http.ResponseWriter, req *http.Request) {
	if !IsAdmin(req) {
		http.Error(writer, "Unauthorized", http.StatusUnauthorized)
		return
	}
	// Get all orders
	orders := []models.Order{}
	database.DB.Preload("Items").Find(&orders)
	json.NewEncoder(writer).Encode(orders)
}

func DeleteOrder(writer http.ResponseWriter, req *http.Request) {
	if !IsAdmin(req) {
		http.Error(writer, "Unauthorized", http.StatusUnauthorized)
		return
	}

	orderID := req.URL.Query().Get("id")
	if orderID == "" {
		http.Error(writer, "Order ID is required", http.StatusBadRequest)
		return
	}

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

	// Find the order
	var order models.Order
	if err := tx.Preload("Items").First(&order, id).Error; err != nil {
		tx.Rollback()
		http.Error(writer, "Order not found", http.StatusNotFound)
		return
	}

	// Clear the many-to-many relationship from the join table
	if err := tx.Model(&order).Association("Items").Clear(); err != nil {
		tx.Rollback()
		http.Error(writer, "Could not delete order items from join table", http.StatusInternalServerError)
		return
	}

	// Delete the order by ID
	if err := tx.Delete(&order).Error; err != nil {
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

func AssignOrder(writer http.ResponseWriter, req *http.Request) {
	if !IsAdmin(req) {
		http.Error(writer, "Unauthorized", http.StatusUnauthorized)
		return
	}

	orderID := req.URL.Query().Get("oid")
	courierID := req.URL.Query().Get("cid")

	if orderID == "" {
		http.Error(writer, "Order ID is required", http.StatusBadRequest)
		return
	}
	if courierID == "" {
		http.Error(writer, "Courier ID is required", http.StatusBadRequest)
		return
	}

	oid, err := strconv.Atoi(orderID)
	if err != nil {
		http.Error(writer, "Invalid order ID", http.StatusBadRequest)
		return
	}
	cid, err := strconv.Atoi(courierID)
	if err != nil {
		http.Error(writer, "Invalid courier ID", http.StatusBadRequest)
		return
	}

	// Get the courier
	var courier models.User
	if err := database.DB.First(&courier, cid).Error; err != nil {
		http.Error(writer, "Courier not found", http.StatusNotFound)
		return
	}

	// Check if the user is a courier
	if courier.Role != "courier" {
		http.Error(writer, "User is not a courier", http.StatusBadRequest)
		return
	}

	// Get the order
	var order models.Order
	if err := database.DB.First(&order, oid).Error; err != nil {
		http.Error(writer, "Order not found", http.StatusNotFound)
		return
	}

	// Assign the order to the courier
	order.CourierID = &cid
	if err := database.DB.Save(&order).Error; err != nil {
		http.Error(writer, "Could not assign order", http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	msg := "Order no." + strconv.Itoa(order.ID) + " assigned to " + courier.Name
	json.NewEncoder(writer).Encode(map[string]string{"message": msg})
}

func UpdateOrderDetails(writer http.ResponseWriter, req *http.Request) {
	if !IsAdmin(req) {
		http.Error(writer, "Unauthorized", http.StatusUnauthorized)
		return
	}

	orderID := req.URL.Query().Get("id")
	if orderID == "" {
		http.Error(writer, "Order ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(orderID)
	if err != nil {
		http.Error(writer, "Invalid order ID", http.StatusBadRequest)
		return
	}

	// Get the order
	var order models.Order
	if err := database.DB.First(&order, id).Error; err != nil {
		http.Error(writer, "Order not found", http.StatusNotFound)
		return
	}

	// Decode the request body
	var updatedOrder models.Order
	if err := json.NewDecoder(req.Body).Decode(&updatedOrder); err != nil {
		http.Error(writer, "Could not decode request body", http.StatusBadRequest)
		return
	}

	// Update the order
	order.PickupLocation = updatedOrder.PickupLocation
	order.DropOffLocation = updatedOrder.DropOffLocation
	order.DeliveryTime = updatedOrder.DeliveryTime
	order.Status = updatedOrder.Status
	order.Items = updatedOrder.Items

	if err := database.DB.Save(&order).Error; err != nil {
		http.Error(writer, "Could not update order", http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(map[string]string{"message": "Order updated successfully"})
}
