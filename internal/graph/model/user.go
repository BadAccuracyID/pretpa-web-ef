package model

type User struct {
	ID             uint   `json:"id" gorm:"primaryKey"`
	Name           string `json:"name"`
	Email          string `json:"email" gorm:"uniqueIndex;not null"`
	Username       string `json:"username" gorm:"uniqueIndex;not null"`
	HashedPassword string `json:"hashedPassword"`

	JWTToken string `json:"jwtToken" gorm:"-"`
}
