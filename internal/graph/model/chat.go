package model

import (
	"github.com/badaccuracyid/tpa-web-ef/internal/utils"
)

type Conversation struct {
	utils.DefaultIncrementModel
	Users    []User `gorm:"many2many:user_conversations;"`
	Messages []Message
}

type Message struct {
	utils.DefaultIncrementModel
	UserID         string       `json:"user_id"`
	User           User         `gorm:"not null;foreignKey:UserID"`
	ConversationID uint         `json:"conversation_id"`
	Conversation   Conversation `gorm:"foreignKey:ConversationID"`
	Content        string       `json:"content"`
}

type ConversationSubscription struct {
	utils.DefaultIncrementModel
	UserID         string       `json:"user_id"`
	User           User         `gorm:"foreignKey:UserID"`
	ConversationID uint         `json:"conversation_id"`
	Conversation   Conversation `gorm:"foreignKey:ConversationID"`
}
