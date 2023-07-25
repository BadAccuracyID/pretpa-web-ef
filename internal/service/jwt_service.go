package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTService interface {
	GenerateToken(userID string) (string, error)
	ValidateToken(tokenString string) (string, error)
}

type jwtCustomClaims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issure    string
}

func NewJWTService(secretKey string, issuer string) JWTService {
	return &jwtService{
		secretKey: secretKey,
		issure:    issuer,
	}
}

func (s *jwtService) GenerateToken(userID string) (string, error) {
	claims := &jwtCustomClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(), // 2 hours
			Issuer:    s.issure,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(s.secretKey))
}

func (s *jwtService) ValidateToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&jwtCustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return []byte(s.secretKey), nil
		},
	)

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*jwtCustomClaims); ok && token.Valid {
		return claims.UserID, nil
	}

	return "", errors.New("invalid token")
}
