package repository

import (
	"context"

	"github.com/NishimuraTakuya-nt/go-rest-clean-plane/internal/core/domain/models"
)

type ProductRepository interface {
	Get(ctx context.Context, id string) (*models.Product, error)
	List(ctx context.Context, offset, limit *int) ([]*models.Product, error)
}
