package models

type User struct {
	ID       uint    `json:"id" gorm:"primary_key"`
	Name     string  `json:"name" gorm:"not null"`
	Email    string  `json:"email" gorm:"unique;not null"`
	Phone    string  `json:"phone" gorm:"not null"`
	Password string  `json:"password" gorm:"not null"`
	Role     string  `json:"role" gorm:"default:'user';not null"` // user, courier
	Orders   []Order `json:"orders" gorm:"foreignKey:UserID"`
}
