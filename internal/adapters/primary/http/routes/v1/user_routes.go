package v1

import (
	"net/http"

	"github.com/NishimuraTakuya-nt/go-rest-clean-plane/internal/adapters/primary/http/handlers"
	"github.com/NishimuraTakuya-nt/go-rest-clean-plane/internal/core/usecases"
)

func SetupUserRoutes(mux *http.ServeMux, userUseCase usecases.UserUseCase) {
	userHandler := handlers.NewUserHandler(userUseCase)
	mux.Handle("/users", userHandler)
	mux.Handle("/user/", userHandler)
}
