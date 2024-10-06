package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/NishimuraTakuya-nt/go-rest-clean-plane/internal/infrastructure/auth"
	"github.com/NishimuraTakuya-nt/go-rest-clean-plane/internal/infrastructure/logger"
)

type AuthHandler struct {
	authService auth.TokenService
}

func NewAuthHandler(authService auth.TokenService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodPost && r.URL.Path == "/auth/login":
		h.Login(w, r)
	// 将来的に他の認証関連エンドポイントを追加する場合、ここに追加します
	// 例: case r.Method == http.MethodPost && r.URL.Path == "/auth/logout":
	//        h.Logout(w, r)
	default:
		http.NotFound(w, r)
	}
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

// fixme usecase に処理を移す
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	log := logger.GetLogger()

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Error("Failed to decode login request", "error", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// TODO: ユーザー認証のロジックを実装する
	// この例では、単純化のためにユーザー名とパスワードのチェックを省略しています
	userID := req.Username
	roles := []string{"role:teamA:editor", "role:teamB:viewer"} // 実際のアプリケーションでは、データベースからユーザーのロールを取得する必要があります

	token, err := h.authService.GenerateToken(userID, roles)
	if err != nil {
		log.Error("Failed to generate token", "error", err)
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	response := LoginResponse{Token: token}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error("Failed to encode login response", "error", err)
		return
	}

	log.Info("Login successful", "user_id", userID)
}
