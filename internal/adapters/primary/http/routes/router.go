package routes

import (
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/NishimuraTakuya-nt/go-rest-clean-plane/internal/adapters/primary/http/handlers"
	"github.com/NishimuraTakuya-nt/go-rest-clean-plane/internal/adapters/primary/http/middleware"
	v1 "github.com/NishimuraTakuya-nt/go-rest-clean-plane/internal/adapters/primary/http/routes/v1"
	"github.com/NishimuraTakuya-nt/go-rest-clean-plane/internal/core/usecases"
	"github.com/NishimuraTakuya-nt/go-rest-clean-plane/internal/infrastructure/auth"
	"github.com/NishimuraTakuya-nt/go-rest-clean-plane/internal/infrastructure/config"
)

func SetupRouter(
	cfg *config.Config,
	tokenService auth.TokenService,
	userUseCase usecases.UserUseCase,
) http.Handler {
	mux := http.NewServeMux()

	// ルートハンドラの登録
	mux.HandleFunc("/", handlers.HomeHandler)
	// Swagger 2.0
	mux.Handle("/swagger/2.0/", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))
	// OAS 3.0
	mux.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("/docs/swagger/openapi3.json"),
	))
	mux.Handle("/docs/swagger/", http.StripPrefix("/docs/swagger/", http.FileServer(http.Dir("./docs/swagger"))))

	// API v1 ルート
	apiV1 := http.NewServeMux()
	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", apiV1))

	v1.SetupHealthcheckRoutes(apiV1)
	v1.SetupAuthRoutes(apiV1, tokenService)

	v1.SetupUserRoutes(apiV1, userUseCase)
	v1.SetupProductRoutes(apiV1)
	v1.SetupOrderRoutes(apiV1)

	// CORSの設定
	corsConfig := middleware.DefaultCORSConfig()
	corsConfig.AllowOrigins = cfg.AllowedOrigins

	// ミドルウェアの適用
	handler := middleware.Chain(
		mux,
		middleware.CORS(corsConfig),
		middleware.Logging(),
		middleware.ErrorHandler(),
		//middleware.Authenticate(tokenService), // fix 面倒なので一旦OFF
	)
	return handler
}
