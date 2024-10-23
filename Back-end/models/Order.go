package models

import (
	"fmt"
	"time"
)

type Order struct {
	ID              int64     `json:"id" gorm:"primary_key"`
	UserID          int       `json:"user_id" gorm:"not null"`
	PickupLocation  string    `json:"pickup_location" gorm:"not null"`
	DropOffLocation string    `json:"dropoff_location" gorm:"not null"`
	DeliveryTime    time.Time `json:"delivery_time"`
	Status          string    `json:"status" gorm:"default:'pending'"`
	CreatedAt       time.Time `json:"created_at"`
	Items           []Item    `json:"items" gorm:"many2many:order_items;"`
}

func (o Order) String() string {
	return fmt.Sprintf(
		"Order: ID: %d, UserID: %d, PickupLocation: %s, DropOffLocation: %s, DeliveryTime: %s, Status: %s, CreatedAt: %s Items: %v",
		o.ID, o.UserID, o.PickupLocation, o.DropOffLocation, o.DeliveryTime, o.Status, o.CreatedAt, o.Items,
	)
}
