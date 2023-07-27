package service

import (
	"context"
	"errors"
	"github.com/badaccuracyid/tpa-web-ef/internal/graph/model"
	"github.com/badaccuracyid/tpa-web-ef/internal/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"os"
)

type AuthService interface {
	Login(input *model.LoginInput) (*model.User, error)
	Register(user *model.RegisterInput) (*model.User, error)
	ChangePassword(oldPassword string, newPassword string) (*model.User, error)
}

type authService struct {
	ctx context.Context
	db  *gorm.DB
}

func NewAuthService(ctx context.Context, db *gorm.DB) AuthService {
	return &authService{ctx: ctx, db: db}
}

func (a *authService) Login(input *model.LoginInput) (*model.User, error) {
	var user *model.User
	if err := a.db.Where("username = ?", input.Username).First(&user).Error; err != nil {
		return nil, err
	}

	if err := utils.CheckHash(input.Password, user.HashedPassword); err != nil {
		return nil, err
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	jwtToken, err := NewJWTService(jwtSecret).GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	user.JWTToken = jwtToken
	return user, nil
}

func (a *authService) Register(input *model.RegisterInput) (*model.User, error) {
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		ID:             uuid.New().String(),
		Name:           input.Username,
		Email:          input.Email,
		Username:       input.Username,
		HashedPassword: hashedPassword,
	}

	//if err := a.db.Create(user).Error; err != nil {
	//	return nil, err
	//}

	jwtSecret := os.Getenv("JWT_SECRET")
	jwtToken, err := NewJWTService(jwtSecret).GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	user.JWTToken = jwtToken
	return user, nil
}

func (a *authService) ChangePassword(oldPassword string, newPassword string) (*model.User, error) {
	userId := utils.GetCurrentUserID(a.ctx)
	if userId == "" {
		return nil, errors.New("userid key not found")
	}

	userService := NewUserService(a.ctx, a.db)
	user, err := userService.GetUser(userId)
	if err != nil {
		return nil, err
	}

	if err := utils.CheckHash(oldPassword, user.HashedPassword); err != nil {
		return nil, err
	}

	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return nil, err
	}

	user.HashedPassword = hashedPassword
	if err := a.db.Model(user).Where("id = ?", userId).Updates(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
