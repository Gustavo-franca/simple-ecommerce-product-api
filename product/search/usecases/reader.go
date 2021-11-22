package usecases

import (
	"simpleecommerceproductapi/product"
	"simpleecommerceproductapi/product/search"
)

type (
	Reader struct {
		repository RepositoryReader
	}
	RepositoryReader interface {
		GetByID(id string) (product.Entity, error)
		GetByParams(params search.Params) ([]product.Entity, error)
	}
)

func NewReader(repository RepositoryReader) Reader {
	return Reader{
		repository: repository,
	}
}
func (r Reader) SearchByFilters(params search.Params) ([]product.Entity, error) {
	return r.repository.GetByParams(params)
}

func (r Reader) SearchByID(id string) (product.Entity, error) {
	return r.repository.GetByID(id)
}
