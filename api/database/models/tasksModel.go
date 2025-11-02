package models

import "github.com/google/uuid"

type Task struct {
	ID 				uuid.UUID `json:"id" gorm:"primaryKey"`
	UserID 			uuid.UUID `json:"user_id"`
	Title 			string `json:"title"`
	Description 	string `json:"description"`
	Completed 		bool `json:"completed"`
}