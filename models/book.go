package models

import "time"

type Book struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	BookID    int  `json:"book_id"`
	CreatedAt time.Time
	Name      string `json:"name"`
	State     string `json:"state"`
}
