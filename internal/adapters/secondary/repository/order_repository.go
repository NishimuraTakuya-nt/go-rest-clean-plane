package repository

import (
	"context"

	"github.com/NishimuraTakuya-nt/go-rest-clean-plane/internal/core/domain/models"
)

type OrderRepository interface {
	Get(ctx context.Context, id string) (*models.Order, error)
	List(ctx context.Context, offset, limit *int) ([]*models.Order, error)
}
