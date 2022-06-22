package usecases

import (
	"context"
	"go-clean-template/app/domain/model"
)

type BookRepository interface {
	Create(ctx context.Context, book model.Book) error
	Edit(ctx context.Context, book model.Book) error
	Delete(ctx context.Context, id int64) error
	FindByID(ctx context.Context, id int64) (*model.Book, error)
	FindAll(ctx context.Context) (*model.Books, error)
}
