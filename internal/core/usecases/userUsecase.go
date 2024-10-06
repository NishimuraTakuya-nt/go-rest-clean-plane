package usecases

import (
	"context"

	"github.com/NishimuraTakuya-nt/go-rest-clean-plane/internal/adapters/secondary/graphql"
	"github.com/NishimuraTakuya-nt/go-rest-clean-plane/internal/core/domain/models"
)

type UserUseCase interface {
	GetUser(ctx context.Context, id string) (*models.User, error)
}

type userUseCase struct {
	graphqlClient graphql.Client
}

func NewUserUseCase(client graphql.Client) UserUseCase {
	return &userUseCase{
		graphqlClient: client,
	}
}

func (uc *userUseCase) GetUser(ctx context.Context, id string) (*models.User, error) {
	// todo trace log
	return uc.graphqlClient.GetUser(ctx, id)
}
