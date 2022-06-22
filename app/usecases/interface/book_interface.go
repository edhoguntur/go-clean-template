package usecases

import (
	"context"
	"go-clean-template/app/entities/domain"
)

type BookRepository interface {
	Create(ctx context.Context, book domain.Book) error
	Edit(ctx context.Context, book domain.Book) error
	Delete(ctx context.Context, id int64) error
	FindByID(ctx context.Context, id int64) (*domain.Book, error)
	FindAll(ctx context.Context) (*domain.Books, error)
}
