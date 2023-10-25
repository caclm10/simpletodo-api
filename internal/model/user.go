package model

import "time"

type User struct {
	ID         uint `gorm:"primaryKey"`
	Name       string
	Email      string
	Password   string
	PictureURL string
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
}
