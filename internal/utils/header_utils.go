package utils

import (
	"context"
)

const (
	UserIDKey = "UserID"
)

func GetCurrentUserID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}

	if ctx.Value(UserIDKey) == nil {
		return ""
	}

	userId := ctx.Value(UserIDKey).(string)
	return userId
}
