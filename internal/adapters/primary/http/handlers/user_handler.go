package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/NishimuraTakuya-nt/go-rest-clean-plane/internal/adapters/primary/http/middleware"
	"github.com/NishimuraTakuya-nt/go-rest-clean-plane/internal/apperrors"
	"github.com/NishimuraTakuya-nt/go-rest-clean-plane/internal/core/usecases"
	"github.com/NishimuraTakuya-nt/go-rest-clean-plane/internal/infrastructure/logger"
)

type UserHandler struct {
	userUseCase usecases.UserUseCase
}

func NewUserHandler(userUseCase usecases.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet && strings.HasPrefix(r.URL.Path, "/user/"):
		h.Get(w, r)

	case r.Method == http.MethodGet && r.URL.Path == "/users":
		h.List(w, r)
	case r.Method == http.MethodPost && r.URL.Path == "/user":
		h.Create(w, r)
	case r.Method == http.MethodPut && strings.HasPrefix(r.URL.Path, "/user"):
		h.Update(w, r)
	case r.Method == http.MethodDelete && strings.HasPrefix(r.URL.Path, "/user"):
		h.Delete(w, r)
	default:
		http.NotFound(w, r)
	}
}

func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.GetLogger()
	requestID, _ := ctx.Value(middleware.RequestIDKey).(string)

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 {
		log.Warn("Invalid user ID in request", "path", r.URL.Path, "request_id", requestID)
		writeError(w, apperrors.NewBadRequestError("Invalid user ID", nil))
		return
	}
	userID := parts[2]

	user, err := h.userUseCase.GetUser(ctx, userID)
	// ///////////////////////////////////////////////////////////////////////
	num, _ := strconv.Atoi(userID)
	switch num {
	case 0:
		e := fmt.Errorf("this is a test validation error")
		err = apperrors.NewValidationError(e.Error(), e)

	case 1:
		e := fmt.Errorf("this is a test Resource not found error")
		err = apperrors.NewNotFoundError(e.Error(), e)

	case 2:
		e := fmt.Errorf("this is a test Unauthorized error")
		err = apperrors.NewUnauthorizedError(e.Error(), e)

	case 3:
		panic("This is a test panic in healthcheck")

	case 10:
		e := apperrors.NewValidationErrors()
		e.AddError("test_field1", "This is a test validation error1")
		e.AddError("test_field2", "This is a test validation error2")
		err = e

	default:
		err = nil
	}

	// ///////////////////////////////////////////////////////////////////////
	if err != nil {
		log.Error("Failed to get user", "error", err, "user_id", userID, "request_id", requestID)
		writeError(w, err)
		return
	}

	writeJSONResponse(w, user, requestID)
}

func (h *UserHandler) List(w http.ResponseWriter, r *http.Request) {
	// ユーザー一覧の取得処理
	json.NewEncoder(w).Encode(map[string]string{"message": "List users"})
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	// ユーザー作成処理
	json.NewEncoder(w).Encode(map[string]string{"message": "Create user"})
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	// ユーザー更新処理
	json.NewEncoder(w).Encode(map[string]string{"message": "Update user"})
}

func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// ユーザー削除処理
	json.NewEncoder(w).Encode(map[string]string{"message": "Delete user"})
}
