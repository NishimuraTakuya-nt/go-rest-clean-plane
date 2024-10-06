package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
)

type OrderHandler struct{}

func NewOrderHandler() *OrderHandler {
	return &OrderHandler{}
}

func (h *OrderHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet && r.URL.Path == "/orders":
		h.List(w, r)
	case r.Method == http.MethodPost && r.URL.Path == "/order":
		h.Create(w, r)
	case r.Method == http.MethodGet && strings.HasPrefix(r.URL.Path, "/order/"):
		h.Get(w, r)
	case r.Method == http.MethodPut && strings.HasPrefix(r.URL.Path, "/order/"):
		h.Update(w, r)
	case r.Method == http.MethodDelete && strings.HasPrefix(r.URL.Path, "/order/"):
		h.Delete(w, r)
	default:
		http.NotFound(w, r)
	}
}

func (h *OrderHandler) List(w http.ResponseWriter, _ *http.Request) {
	// 注文一覧の取得処理
	// nolint:errcheck
	json.NewEncoder(w).Encode(map[string]string{"message": "List orders"})
}

func (h *OrderHandler) Create(w http.ResponseWriter, _ *http.Request) {
	// 注文作成処理
	// nolint:errcheck
	json.NewEncoder(w).Encode(map[string]string{"message": "Create order"})
}

func (h *OrderHandler) Get(w http.ResponseWriter, _ *http.Request) {
	// 特定の注文取得処理
	// nolint:errcheck
	json.NewEncoder(w).Encode(map[string]string{"message": "Get order"})
}

func (h *OrderHandler) Update(w http.ResponseWriter, _ *http.Request) {
	// 注文更新処理
	// nolint:errcheck
	json.NewEncoder(w).Encode(map[string]string{"message": "Update order"})
}

func (h *OrderHandler) Delete(w http.ResponseWriter, _ *http.Request) {
	// 注文削除処理
	// nolint:errcheck
	json.NewEncoder(w).Encode(map[string]string{"message": "Delete order"})
}
