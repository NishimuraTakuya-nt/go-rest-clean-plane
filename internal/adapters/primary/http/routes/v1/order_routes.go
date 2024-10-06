package v1

import (
	"net/http"

	"github.com/NishimuraTakuya-nt/go-rest-clean-plane/internal/adapters/primary/http/handlers"
)

func SetupOrderRoutes(mux *http.ServeMux) {
	orderHandler := handlers.NewOrderHandler()
	mux.Handle("/orders", orderHandler)
	mux.Handle("/order/", orderHandler)
}
