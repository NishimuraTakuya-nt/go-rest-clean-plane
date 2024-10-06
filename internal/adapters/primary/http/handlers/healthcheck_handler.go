package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
)

type HealthcheckHandler struct{}

func NewHealthcheckHandler() *HealthcheckHandler {
	return &HealthcheckHandler{}
}

func (h *HealthcheckHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet && strings.HasPrefix(r.URL.Path, "/healthcheck"):
		h.Get(w, r)
	default:
		http.NotFound(w, r)
	}
}

func (h *HealthcheckHandler) Get(w http.ResponseWriter, _ *http.Request) {
	// healthcheck
	// nolint:errcheck
	json.NewEncoder(w).Encode(map[string]string{"message": "healthcheck ok"})
}
