package models

import "time"

// User represents the users table
type User struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	Name      string  `json:"name" gorm:"size:255;not null"`
	Email     string  `json:"email" gorm:"size:191;unique;not null"`
	Phone     *string `json:"phone" gorm:"size:20;unique"`
	Password  string  `gorm:"not null"` // Hashed password
	CreatedAt time.Time
	UpdatedAt time.Time
}

var Users []User
