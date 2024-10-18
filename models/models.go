package models

import (
	"time"
)

type Order struct {
	ID              uint      `json:"id" gorm:"primary_key"`
	UserID          uint      `json:"user_id" gorm:"not null"`
	PickupLocation  string    `json:"pickup_location" gorm:"not null"`
	DropoffLocation string    `json:"dropoff_location" gorm:"not null"`
	PackageDetails  string    `json:"package_details" gorm:"not null"`
	DeliveryTime    time.Time `json:"delivery_time"`
	Status          string    `json:"status" gorm:"default:'pending'"`
	CreatedAt       time.Time `json:"created_at"`
}

type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Name     string `json:"name" gorm:"not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	Phone    string `json:"phone" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
}
