package usecases

import (
	"context"
	"simpleecommerceproductapi/product"
)

type (
	Writer struct {
		repository RepositoryWriter
	}
	RepositoryWriter interface {
		Create(ctx context.Context, product product.Entity) (string, error)
	}
)

func NewCreate(repository RepositoryWriter) Writer {
	return Writer{
		repository: repository,
	}
}
func (r Writer) Create(ctx context.Context, params product.Entity) (string, error) {
	return r.repository.Create(ctx, params)
}
