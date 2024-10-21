package models

import (
	"fmt"
	"time"
)

type Order struct {
	ID              int       `json:"id" gorm:"primary_key"`
	UserID          int       `json:"user_id" gorm:"not null"`
	PickupLocation  string    `json:"pickup_location" gorm:"not null"`
	DropOffLocation string    `json:"dropoff_location" gorm:"not null"`
	PackageDetails  string    `json:"package_details" gorm:"not null"`
	DeliveryTime    time.Time `json:"delivery_time"`
	Status          string    `json:"status" gorm:"default:'pending'"`
	CreatedAt       time.Time `json:"created_at"`
}

func (o Order) String() string {
	return fmt.Sprintf(
		"Order{ID: %d, UserID: %d, PickupLocation: %s, DropOffLocation: %s, PackageDetails: %s, DeliveryTime: %s, Status: %s, CreatedAt: %s}",
		o.ID,
		o.UserID,
		o.PickupLocation,
		o.DropOffLocation,
		o.PackageDetails,
		o.DeliveryTime.Format(time.RFC3339), // Format the time if needed
		o.Status,
		o.CreatedAt.Format(time.RFC3339), // Format the time if needed
	)
}
