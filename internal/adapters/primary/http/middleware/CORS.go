package middleware

import (
	"net/http"
	"strings"
)

// CORSConfig defines the configuration options for CORS middleware
type CORSConfig struct {
	AllowOrigins     []string
	AllowMethods     []string
	AllowHeaders     []string
	AllowCredentials bool
	ExposeHeaders    []string
	MaxAge           int
}

// DefaultCORSConfig returns a default CORS configuration
func DefaultCORSConfig() CORSConfig {
	return CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: false,
		ExposeHeaders:    []string{},
		MaxAge:           300, // 5 minutes
	}
}

func CORS(config CORSConfig) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// オリジンの検証
			origin := r.Header.Get("Origin")
			allowOrigin := ""
			for _, o := range config.AllowOrigins {
				if o == "*" || o == origin {
					allowOrigin = o
					break
				}
			}

			// CORSヘッダーの設定
			if allowOrigin != "" {
				w.Header().Set("Access-Control-Allow-Origin", allowOrigin)
				w.Header().Set("Access-Control-Allow-Methods", strings.Join(config.AllowMethods, ", "))
				w.Header().Set("Access-Control-Allow-Headers", strings.Join(config.AllowHeaders, ", "))

				if config.AllowCredentials {
					w.Header().Set("Access-Control-Allow-Credentials", "true")
				}

				if len(config.ExposeHeaders) > 0 {
					w.Header().Set("Access-Control-Expose-Headers", strings.Join(config.ExposeHeaders, ", "))
				}

				if config.MaxAge > 0 {
					w.Header().Set("Access-Control-Max-Age", string(rune(config.MaxAge)))
				}
			}

			// プリフライトリクエスト（OPTIONS）の処理
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			// 次のハンドラを呼び出す
			next.ServeHTTP(w, r)
		})
	}
}
