package models

import "time"

type Order struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	BookRefer int  `json:"book_id"`
	Book      Book `gorm:"foreignKey:BookRefer"`
	UserRefer int  `json:"user_id"`
	User      User `gorm:"foreignKey:UserRefer"`
}
type OrderInput struct {
	BookID     int    `json:"book_id"`
	NationalID string `json:"national_id"`
}
