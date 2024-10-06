package graphql

import (
	"context"
	"time"

	"github.com/NishimuraTakuya-nt/go-rest-clean-plane/internal/core/domain/models"
)

type Client interface {
	GetUser(ctx context.Context, id string) (*models.User, error)
}

type client struct {
	// クライアントの設定など
}

func NewClient() Client {
	return &client{}
}

func (c *client) GetUser(_ context.Context, ID string) (*models.User, error) {
	// ここでは簡易的に固定のユーザーを返していますが、
	// 実際には取得する処理を実装します

	return &models.User{
		ID:        ID,
		Name:      "example_user",
		Roles:     []string{"teamA:editor", "teamB:viewer"},
		Email:     "user@example.com",
		CreatedAt: time.Now().Add(-24 * time.Hour),
		UpdatedAt: time.Now(),
	}, nil
}
