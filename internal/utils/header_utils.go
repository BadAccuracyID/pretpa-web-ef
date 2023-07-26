package utils

import (
	"context"
	"github.com/badaccuracyid/tpa-web-ef/internal/service"
	"os"
)

func getCurrentUserID(jwtToken string) (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")

	userId, err := service.NewJWTService(jwtSecret).ValidateToken(jwtToken)
	if err != nil {
		return "", err
	}
	return userId, nil
}

func GetCurrentUserID(ctx context.Context) (string, error) {
	jwtToken := ctx.Value("token").(string)
	return getCurrentUserID(jwtToken)
}
