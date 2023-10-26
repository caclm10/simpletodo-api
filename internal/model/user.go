package model

import (
	"time"

	"github.com/caclm10/simpletodo-api/internal/model/response"
)

type User struct {
	ID         uint `gorm:"primaryKey"`
	Name       string
	Email      string
	Password   string
	PictureURL string
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
}

func (u User) ToResponse() response.UserResponse {
	return response.UserResponse{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
}
