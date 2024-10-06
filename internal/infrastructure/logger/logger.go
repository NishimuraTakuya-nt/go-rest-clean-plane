package logger

import (
	"log/slog"
	"os"
)

var log *slog.Logger

func init() {
	// 環境に応じてログレベルを設定
	var level slog.Level
	switch os.Getenv("LOG_LEVEL") {
	case "DEBUG":
		level = slog.LevelDebug
	case "INFO":
		level = slog.LevelInfo
	case "WARN":
		level = slog.LevelWarn
	case "ERROR":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	// JSONハンドラーを作成
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	})

	// グローバルロガーを設定
	log = slog.New(handler)
	slog.SetDefault(log)
}

// GetLogger returns the configured logger
func GetLogger() *slog.Logger {
	return log
}
