package service

import (
	"context"
	"errors"
	"github.com/badaccuracyid/tpa-web-ef/internal/graph/model"
	"github.com/badaccuracyid/tpa-web-ef/internal/utils"
	"gorm.io/gorm"
)

type UserService interface {
	UpdateUser(id string, user *model.UserInput) (*model.User, error)
	UpdateCurrentUser(user *model.UserInput) (*model.User, error)
	DeleteUser(id string) (*model.User, error)
	DeleteCurrentUser() (*model.User, error)
	GetCurrentUser() (*model.User, error)
	GetUser(id string) (*model.User, error)
	GetAllUsers() ([]*model.User, error)
}

type userService struct {
	ctx context.Context
	db  *gorm.DB
}

func NewUserService(ctx context.Context, db *gorm.DB) UserService {
	return &userService{ctx: ctx, db: db}
}

// Common update logic for both UpdateUser and UpdateCurrentUser
func (s *userService) updateUserByID(id string, input *model.UserInput) (*model.User, error) {
	user, err := s.GetUser(id)
	if err != nil {
		return nil, err
	}

	user.Name = input.Name
	user.Email = input.Email
	user.Username = input.Username

	if err := s.db.Model(user).Where("id = ?", id).Updates(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) UpdateUser(id string, input *model.UserInput) (*model.User, error) {
	return s.updateUserByID(id, input)
}

func (s *userService) UpdateCurrentUser(input *model.UserInput) (*model.User, error) {
	userId := utils.GetCurrentUserID(s.ctx)
	if userId == "" {
		return nil, errors.New("user not found")
	}

	return s.updateUserByID(userId, input)
}

func (s *userService) deleteUserByID(id string) (*model.User, error) {
	user, err := s.GetUser(id)
	if err != nil {
		return nil, err
	}

	if err := s.db.Where("id = ?", id).Delete(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) DeleteUser(id string) (*model.User, error) {
	return s.deleteUserByID(id)
}

func (s *userService) DeleteCurrentUser() (*model.User, error) {
	userId := utils.GetCurrentUserID(s.ctx)
	if userId == "" {
		return nil, errors.New("user not found")
	}

	return s.DeleteUser(userId)
}

func (s *userService) GetCurrentUser() (*model.User, error) {
	userId := utils.GetCurrentUserID(s.ctx)
	if userId == "" {
		return nil, errors.New("user not found")
	}

	return s.GetUser(userId)
}

func (s *userService) GetUser(id string) (*model.User, error) {
	var user *model.User
	if err := s.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) GetAllUsers() ([]*model.User, error) {
	var users []*model.User
	if err := s.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
