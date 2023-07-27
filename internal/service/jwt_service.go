package service

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JWTService interface {
	GenerateToken(userID string) (string, error)
	ValidateToken(tokenString string) (string, error)
}

type jwtService struct {
	secretKey string
}

type CustomClaim struct {
	jwt.RegisteredClaims
	UserID string `json:"user_id"`
}

func NewJWTService(secretKey string) JWTService {
	return &jwtService{
		secretKey: secretKey,
	}
}

func (s *jwtService) GenerateToken(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "EF23-1",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		UserID: userID,
	})

	return token.SignedString([]byte(s.secretKey))
}

func (s *jwtService) ValidateToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaim{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(s.secretKey), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*CustomClaim)
	if !ok || !token.Valid {
		return "", err
	}

	userId := claims.UserID
	if userId == "" {
		return "", errors.New("user id is empty")
	}

	return userId, nil
}
