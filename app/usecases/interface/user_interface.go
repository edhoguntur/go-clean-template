package usecases

import (
	"context"
	"go-clean-template/app/entities/domain"
)

type UserRepository interface {
	SignUp(ctx context.Context, user domain.User) error
	SignIn(ctx context.Context, email, password string) (*domain.User, error)
}
