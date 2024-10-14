package usecases

import (
	"context"

	"github.com/NishimuraTakuya-nt/go-rest-clean-plane/internal/adapters/secondary/piyographql"
	"github.com/NishimuraTakuya-nt/go-rest-clean-plane/internal/core/domain/models"
	"github.com/NishimuraTakuya-nt/go-rest-clean-plane/internal/infrastructure/logger"
)

type UserUseCase interface {
	Get(ctx context.Context, ID string) (*models.User, error)
	List(ctx context.Context, offset, limit *int) ([]models.User, error)
}

type userUseCase struct {
	graphqlClient piyographql.Client
}

func NewUserUseCase(client piyographql.Client) UserUseCase {
	return &userUseCase{
		graphqlClient: client,
	}
}

func (uc *userUseCase) Get(ctx context.Context, ID string) (*models.User, error) {
	// todo trace log

	n := lenID(ID)
	logger.GetLogger().Info("lenID", "n", n)
	return uc.graphqlClient.GetUser(ctx, ID)
}

func lenID(ID string) int {
	return len(ID)
}

func (uc *userUseCase) List(ctx context.Context, offset, limit *int) ([]models.User, error) {
	return uc.graphqlClient.ListUser(ctx, offset, limit)
}
