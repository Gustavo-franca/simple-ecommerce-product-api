package usecases

import (
	"context"
	"simpleecommerceproductapi/product"
	"simpleecommerceproductapi/product/search"
)

type (
	Reader struct {
		repository RepositoryReader
	}
	RepositoryReader interface {
		GetByID(ctx context.Context, id string) (product.Entity, error)
		GetByParams(ctx context.Context, params search.Params) ([]product.Entity, error)
	}
)

func NewReader(repository RepositoryReader) Reader {
	return Reader{
		repository: repository,
	}
}
func (r Reader) SearchByFilters(ctx context.Context, params search.Params) ([]product.Entity, error) {
	return r.repository.GetByParams(ctx, params)
}

func (r Reader) SearchByID(ctx context.Context, id string) (product.Entity, error) {
	return r.repository.GetByID(ctx, id)
}
