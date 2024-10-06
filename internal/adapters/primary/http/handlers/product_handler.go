package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
)

type ProductHandler struct{}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}

func (h *ProductHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet && r.URL.Path == "/products":
		h.List(w, r)
	case r.Method == http.MethodPost && r.URL.Path == "/product":
		h.Create(w, r)
	case r.Method == http.MethodGet && strings.HasPrefix(r.URL.Path, "/product/"):
		h.Get(w, r)
	case r.Method == http.MethodPut && strings.HasPrefix(r.URL.Path, "/product/"):
		h.Update(w, r)
	case r.Method == http.MethodDelete && strings.HasPrefix(r.URL.Path, "/product/"):
		h.Delete(w, r)
	default:
		http.NotFound(w, r)
	}
}

func (h *ProductHandler) List(w http.ResponseWriter, _ *http.Request) {
	// 商品一覧の取得処理
	// nolint:errcheck
	json.NewEncoder(w).Encode(map[string]string{"message": "List products"})
}

func (h *ProductHandler) Create(w http.ResponseWriter, _ *http.Request) {
	// 商品作成処理
	// nolint:errcheck
	json.NewEncoder(w).Encode(map[string]string{"message": "Create product"})
}

func (h *ProductHandler) Get(w http.ResponseWriter, _ *http.Request) {
	// 特定の商品取得処理
	// nolint:errcheck
	json.NewEncoder(w).Encode(map[string]string{"message": "Get product"})
}

func (h *ProductHandler) Update(w http.ResponseWriter, _ *http.Request) {
	// 商品更新処理
	// nolint:errcheck
	json.NewEncoder(w).Encode(map[string]string{"message": "Update product"})
}

func (h *ProductHandler) Delete(w http.ResponseWriter, _ *http.Request) {
	// 商品削除処理
	// nolint:errcheck
	json.NewEncoder(w).Encode(map[string]string{"message": "Delete product"})
}
