package utils

import (
	"context"
)

const (
	UserIDKey = "UserID"
)

func GetCurrentUserID(ctx context.Context) string {
	userId := ctx.Value(UserIDKey).(string)
	return userId
}
