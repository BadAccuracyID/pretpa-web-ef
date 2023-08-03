package model

import "github.com/badaccuracyid/pretpa-web-ef/internal/utils"

type User struct {
	ID string `json:"user_id" gorm:"primaryKey"`
	utils.DefaultCreateUpdate
	Name           string `json:"name"`
	Email          string `json:"email" gorm:"uniqueIndex;not null"`
	Username       string `json:"username" gorm:"uniqueIndex;not null"`
	HashedPassword string `json:"hashedPassword"`

	JWTToken string `json:"jwtToken" gorm:"-"`
}
