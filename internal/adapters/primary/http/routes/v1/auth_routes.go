package v1

import (
	"net/http"

	"github.com/NishimuraTakuya-nt/go-rest-clean-plane/internal/adapters/primary/http/handlers"
	"github.com/NishimuraTakuya-nt/go-rest-clean-plane/internal/infrastructure/auth"
)

func SetupAuthRoutes(mux *http.ServeMux, authService auth.TokenService) {
	authHandler := handlers.NewAuthHandler(authService)
	mux.Handle("/auth/", authHandler)
}
