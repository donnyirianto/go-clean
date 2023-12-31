package interfaces

import (
	"context"

	domain "github.com/donnyirianto/go-clean/pkg/domain"
)

type UserUseCase interface {
	FindAll(ctx context.Context) ([]domain.User, error)
	FindByID(ctx context.Context, id string) (domain.User, error)
	Create(ctx context.Context, user domain.User) (domain.User, error)
	Delete(ctx context.Context, user domain.User) error
	UpdateByID(ctx context.Context, id string, user domain.User) (domain.User, error)
	GetMatchName(ctx context.Context, text string) ([]domain.User, error)
}
