package v1

import (
	"net/http"

	"github.com/NishimuraTakuya-nt/go-rest-clean-plane/internal/adapters/primary/http/handlers"
)

func SetupProductRoutes(mux *http.ServeMux) {
	productHandler := handlers.NewProductHandler()
	mux.Handle("/products", productHandler)
	mux.Handle("/product/", productHandler)
}
