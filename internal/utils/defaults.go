package utils

import (
	"time"
)

type DefaultModel struct {
	ID string `json:"id" gorm:"primaryKey"`
	DefaultCreateUpdate
}

type DefaultCreateUpdate struct {
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime:milli"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime:milli"`
}
