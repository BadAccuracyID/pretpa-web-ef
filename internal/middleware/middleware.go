package middleware

import (
	"context"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		token := request.Header.Get("Authorization")
		if token != "" {
			ctx := context.WithValue(request.Context(), "token", token)
			request = request.WithContext(ctx)
		}

		next.ServeHTTP(writer, request)
	})
}
