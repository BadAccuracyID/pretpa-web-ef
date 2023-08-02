package service

import (
	"context"
	"github.com/badaccuracyid/tpa-web-ef/internal/graph/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"sync"
)

type ChatService interface {
	GetMessage(id string) (*model.Message, error)
	GetConversation(id string) (*model.Conversation, error)

	CreateConversation(input model.CreateConversationInput) (*model.Conversation, error)

	SendMessage(input model.SendMessageInput) (*model.Message, error)
	NewMessageSubscription(conversationID string) (<-chan *model.Message, chan<- struct{}, error)
}

type chatService struct {
	ctx context.Context
	db  *gorm.DB
}

func NewChatService(ctx context.Context, db *gorm.DB) ChatService {
	return &chatService{ctx: ctx, db: db}
}

func (s *chatService) GetMessage(id string) (*model.Message, error) {
	var message *model.Message
	if err := s.db.Where("id = ?", id).First(&message).Error; err != nil {
		return nil, err
	}

	return message, nil
}

func (s *chatService) GetConversation(id string) (*model.Conversation, error) {
	var conversation *model.Conversation
	if err := s.db.Preload("Messages").Preload("Users").Where("id = ?", id).First(&conversation).Error; err != nil {
		return nil, err
	}

	return conversation, nil
}

func (s *chatService) CreateConversation(input model.CreateConversationInput) (*model.Conversation, error) {
	conversation := &model.Conversation{
		ID: uuid.New().String(),
	}

	if err := s.db.Create(conversation).Error; err != nil {
		return nil, err
	}

	return conversation, nil
}

// Subscription system
type Subscription struct {
	messageChan chan *model.Message
	doneChan    chan struct{}
}

var (
	subscriptions      = make(map[string][]*Subscription)
	subscriptionsMutex sync.Mutex
)

func (s *chatService) SendMessage(input model.SendMessageInput) (*model.Message, error) {
	message := &model.Message{
		UserID:         input.UserID,
		ConversationID: input.ConversationID,
		Content:        input.Content,
	}

	if err := s.db.Create(message).Error; err != nil {
		return nil, err
	}

	triggerSubscription(input.ConversationID, message)
	return message, nil
}

func (s *chatService) NewMessageSubscription(conversationId string) (<-chan *model.Message, chan<- struct{}, error) {
	subscription := &Subscription{
		messageChan: make(chan *model.Message),
		doneChan:    make(chan struct{}),
	}

	onSubscribe(conversationId, subscription)
	return subscription.messageChan, subscription.doneChan, nil
}

func onSubscribe(conversationId string, subscription *Subscription) {
	subscriptionsMutex.Lock()
	defer subscriptionsMutex.Unlock()
	subscriptions[conversationId] = append(subscriptions[conversationId], subscription)
}

func triggerSubscription(conversationId string, message *model.Message) {
	subscriptionsMutex.Lock()
	defer subscriptionsMutex.Unlock()

	subscribers, found := subscriptions[conversationId]
	if found {
		for _, subscriber := range subscribers {
			select {
			case <-subscriber.doneChan:
				// Subscription was cancelled, remove it.
				// Here I'm setting the subscriber to nil, but in a production setting,
				// you would probably want to remove it from the slice completely.
				subscriber = nil
			case subscriber.messageChan <- message:
				// Our message went through, do nothing
			}
		}
	}
}
