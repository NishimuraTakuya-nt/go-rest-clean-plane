package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/NishimuraTakuya-nt/go-rest-clean-plane/internal/apperrors"
	"github.com/NishimuraTakuya-nt/go-rest-clean-plane/internal/core/domain/models"
	"github.com/NishimuraTakuya-nt/go-rest-clean-plane/internal/infrastructure/auth"
	"github.com/NishimuraTakuya-nt/go-rest-clean-plane/internal/infrastructure/logger"
)

var UserKey = struct{}{}

var excludedPaths = []string{
	"/api/v1/auth/login",
	"/api/v1/healthcheck",
	"/swagger/",
	"/docs/swagger/",
}

func Authenticate(tokenService auth.TokenService) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log := logger.GetLogger()

			rw, ok := w.(*ResponseWriter)
			if !ok {
				log.Warn("ResponseWriter is not of type *ResponseWriter")
				return
			}

			// 除外パスのチェック
			for _, path := range excludedPaths {
				if strings.HasPrefix(r.URL.Path, path) {
					next.ServeHTTP(rw, r)
					return
				}
			}

			// 認証方法は適宜変更する
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				log.Warn("Missing authorization header")
				rw.WriteError(apperrors.NewUnauthorizedError("Missing authorization header", nil))
				return
			}

			if len(authHeader) <= 7 || authHeader[:7] != "Bearer " {
				log.Warn("Invalid token format", "header", authHeader)
				rw.WriteError(apperrors.NewUnauthorizedError("Invalid token format", nil))
				return
			}

			tokenString := authHeader[7:]
			claims, err := tokenService.ValidateToken(tokenString)
			if err != nil {
				log.Error("Token validation failed", "error", err)
				rw.WriteError(apperrors.NewUnauthorizedError("Invalid or expired token", nil))
				return
			}

			user := &models.User{
				ID:    claims.UserID,
				Roles: claims.Roles,
			}

			// nolint:staticcheck
			ctx := context.WithValue(r.Context(), UserKey, user)
			log.Info("User authenticated", "user_id", user.ID)
			next.ServeHTTP(rw, r.WithContext(ctx))
		})
	}
}
