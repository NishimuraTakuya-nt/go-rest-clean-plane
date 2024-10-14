package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/NishimuraTakuya-nt/go-rest-clean-plane/docs/swagger"
	"github.com/NishimuraTakuya-nt/go-rest-clean-plane/internal/adapters/secondary/piyographql"

	"github.com/NishimuraTakuya-nt/go-rest-clean-plane/internal/adapters/primary/http/routes"
	"github.com/NishimuraTakuya-nt/go-rest-clean-plane/internal/core/usecases"
	"github.com/NishimuraTakuya-nt/go-rest-clean-plane/internal/infrastructure/auth"
	"github.com/NishimuraTakuya-nt/go-rest-clean-plane/internal/infrastructure/config"
	"github.com/NishimuraTakuya-nt/go-rest-clean-plane/internal/infrastructure/logger"
)

// @title Go REST Clean API
// @version 1.0
// @description This is a sample server for a Go REST API using clean architecture.
// @host localhost:8081
// @BasePath /api/v1
func main() {
	if err := run(); err != nil {
		logger.GetLogger().Error("Application failed to run", "error", err)
		os.Exit(1)
	}
}

func run() error {
	log := logger.GetLogger()
	cfg := config.Load()

	tokenService := auth.NewTokenService(cfg.JWTSecretKey)
	graphQLClient := piyographql.NewClient()
	userUseCase := usecases.NewUserUseCase(graphQLClient)

	router := routes.SetupRouter(cfg, tokenService, userUseCase)

	srv := &http.Server{
		Addr:    cfg.ServerAddress,
		Handler: router,
	}

	// シグナルを受け取るためのコンテキストを設定
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt, os.Kill)
	defer stop()

	// サーバーをゴルーチンで起動
	go func() {
		log.Info("Server started", slog.String("address", cfg.ServerAddress))
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error("Server listen failed", slog.String("error", err.Error()))
		}
	}()

	// シグナルを待機
	<-ctx.Done()
	log.Info("Shutdown signal received")

	// シャットダウンのためのコンテキストを作成
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// サーバーをシャットダウン
	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Error("Server forced to shutdown", slog.String("error", err.Error()))
		return err
	}

	// その他のリソースのクリーンアップ
	//if err := graphQLClient.Close(); err != nil {
	//	log.Error("Error closing GraphQL client", slog.String("error", err.Error()))
	//}

	log.Info("Server exited properly")
	return nil
}
