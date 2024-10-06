package config

import (
	"os"
	"strings"
)

type Config struct {
	ServerAddress  string
	AllowedOrigins []string
	JWTSecretKey   string
	// その他の設定項目
}

func Load() *Config {
	return &Config{
		ServerAddress:  getEnv("SERVER_ADDRESS", ":8081"),
		AllowedOrigins: getEnvAsSlice("ALLOWED_ORIGINS", []string{"*"}),
		JWTSecretKey:   getEnv("JWT_SECRET_KEY", "jwt-secret"),
		// その他の設定項目の読み込み
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvAsSlice(key string, fallback []string) []string {
	value := getEnv(key, "")
	if value == "" {
		return fallback
	}
	return strings.Split(value, ",")
}
