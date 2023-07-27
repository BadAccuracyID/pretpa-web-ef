package middleware

import (
	"context"
	"github.com/badaccuracyid/tpa-web-ef/internal/service"
	"net/http"
)

const (
	AuthorizationHeader = "Authorization"
	UserIDKey           = "UserID"
)

type AuthMiddleware struct {
	jwtService service.JWTService
}

func NewAuthMiddleware(service service.JWTService) *AuthMiddleware {
	return &AuthMiddleware{
		jwtService: service,
	}
}

func (a *AuthMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		token := request.Header.Get(AuthorizationHeader)

		if token == "" {
			next.ServeHTTP(writer, request)
			return
		}

		userId, err := a.jwtService.ValidateToken(token)
		if err != nil {
			next.ServeHTTP(writer, request)
			return
		}

		ctx := context.WithValue(request.Context(), UserIDKey, userId)
		request = request.WithContext(ctx)

		next.ServeHTTP(writer, request)
	})
}
