package usecases

import (
	"context"
	"go-clean-template/app/domain/model"
)

type UserRepositoryImpl struct {
}

func (u *UserRepositoryImpl) SignUp(ctx context.Context, user model.User) error {
	return nil
}

func (u *UserRepositoryImpl) SignIn(ctx context.Context, email, password string) (*model.Book, error) {
	return nil, nil
}
