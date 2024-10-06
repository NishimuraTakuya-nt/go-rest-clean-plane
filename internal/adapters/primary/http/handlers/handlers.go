package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/NishimuraTakuya-nt/go-rest-clean-plane/internal/apperrors"
	"github.com/NishimuraTakuya-nt/go-rest-clean-plane/internal/infrastructure/logger"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	err := json.NewEncoder(w).Encode(map[string]string{"message": "Welcome to the API"})
	if err != nil {
		logger.GetLogger().Warn("Failed to encode response", "error", err)
		return
	}
}

func writeError(w http.ResponseWriter, err error) {
	if ew, ok := w.(apperrors.ErrorWriter); ok {
		ew.WriteError(err)
	} else {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func writeJSONResponse(w http.ResponseWriter, data any, requestID string) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log := logger.GetLogger()
		log.Error("Failed to encode response", "error", err, "request_id", requestID)
		writeError(w, apperrors.NewInternalError("Failed to encode response", err))
	}
}
