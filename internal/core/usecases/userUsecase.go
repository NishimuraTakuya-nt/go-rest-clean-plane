package usecases

import (
	"context"

	"github.com/NishimuraTakuya-nt/go-rest-clean-plane/internal/adapters/secondary/graphql"
	"github.com/NishimuraTakuya-nt/go-rest-clean-plane/internal/core/domain/models"
	"github.com/NishimuraTakuya-nt/go-rest-clean-plane/internal/infrastructure/logger"
)

type UserUseCase interface {
	GetUser(ctx context.Context, ID string) (*models.User, error)
}

type userUseCase struct {
	graphqlClient graphql.Client
}

func NewUserUseCase(client graphql.Client) UserUseCase {
	return &userUseCase{
		graphqlClient: client,
	}
}

func (uc *userUseCase) GetUser(ctx context.Context, ID string) (*models.User, error) {
	// todo trace log

	n := lenID(ID)
	logger.GetLogger().Info("lenID", "n", n)
	return uc.graphqlClient.GetUser(ctx, ID)
}

func lenID(ID string) int {
	return len(ID)
}
