package service

import (
	"github.com/badaccuracyid/tpa-web-ef/internal/graph/model"
	"github.com/badaccuracyid/tpa-web-ef/internal/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserService interface {
	CreateUser(user *model.UserInput) (*model.User, error)
	UpdateUser(id string, user *model.UserInput) (*model.User, error)
	DeleteUser(id string) (*model.User, error)
	GetAllUsers() ([]*model.User, error)
	GetUser(id string) (*model.User, error)
}

type userService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) UserService {
	return &userService{db: db}
}

func (s *userService) CreateUser(input *model.UserInput) (*model.User, error) {
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		ID:             uuid.New().String(),
		Name:           input.Name,
		Email:          input.Email,
		Username:       input.Username,
		HashedPassword: hashedPassword,
	}

	if err := s.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) UpdateUser(id string, input *model.UserInput) (*model.User, error) {
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		ID:             id,
		Name:           input.Name,
		Email:          input.Email,
		Username:       input.Username,
		HashedPassword: hashedPassword,
	}

	if err := s.db.Model(user).Where("id = ?", id).Updates(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) DeleteUser(id string) (*model.User, error) {
	user := &model.User{
		ID: id,
	}

	if err := s.db.Where("id = ?", id).Delete(user).Error; err != nil {
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

func (s *userService) GetUser(id string) (*model.User, error) {
	var user *model.User
	if err := s.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
