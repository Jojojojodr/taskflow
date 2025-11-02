package models

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `json:"id" gorm:"primaryKey"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
}
