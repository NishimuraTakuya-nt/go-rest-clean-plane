package v1

import (
	"net/http"

	"github.com/NishimuraTakuya-nt/go-rest-clean-plane/internal/adapters/primary/http/handlers"
)

func SetupHealthcheckRoutes(mux *http.ServeMux) {
	healthcheckHandler := handlers.NewHealthcheckHandler()
	mux.Handle("/healthcheck", healthcheckHandler)
}
