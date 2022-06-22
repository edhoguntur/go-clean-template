package usecases

import (
	"context"
	"go-clean-template/app/entities/domain"
)

type UserRepositoryImpl struct {
}

func (u *UserRepositoryImpl) SignUp(ctx context.Context, user domain.User) error {
	return nil
}

func (u *UserRepositoryImpl) SignIn(ctx context.Context, email, password string) (*domain.Book, error) {
	return nil, nil
}
