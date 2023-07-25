package model

type User struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email" gorm:"unique"`
	Username       string `json:"username" gorm:"unique"`
	HashedPassword string `json:"hashedPassword"`
}
