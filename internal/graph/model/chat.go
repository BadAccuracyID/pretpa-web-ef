package model

import (
	"time"
)

type Conversation struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Users     []User    `gorm:"many2many:user_conversations;"`
	Messages  []Message `gorm:"foreignKey:ConversationID"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime:milli"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime:milli"`
}

type Message struct {
	ID             string       `json:"id" gorm:"primaryKey"`
	UserID         string       `json:"user_id"`
	User           User         `gorm:"not null;foreignKey:UserID"`
	ConversationID string       `json:"conversation_id"`
	Conversation   Conversation `gorm:"foreignKey:ConversationID"`
	Content        string       `json:"content"`
	CreatedAt      time.Time    `json:"createdAt" gorm:"autoCreateTime:milli"`
	UpdatedAt      time.Time    `json:"updatedAt" gorm:"autoUpdateTime:milli"`
}
