package usecases

import (
	"context"
	"go-clean-template/app/domain/model"
)

type UserRepository interface {
	SignUp(ctx context.Context, user model.User) error
	SignIn(ctx context.Context, email, password string) (*model.User, error)
}
