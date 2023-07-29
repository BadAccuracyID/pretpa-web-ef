package utils

import (
	"time"
)

type DefaultIncrementModel struct {
	ID uint `json:"id" gorm:"primaryKey;autoIncrement"`
	DefaultCreateUpdate
}

type DefaultCreateUpdate struct {
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime:milli"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime:milli"`
}
