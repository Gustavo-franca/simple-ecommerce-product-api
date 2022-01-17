package usecases

import "simpleecommerceproductapi/product"

type (
	Writer struct {
		repository RepositoryWriter
	}
	RepositoryWriter interface {
		Create(product product.Entity) (string, error)
	}
)

func NewCreate(repository RepositoryWriter) Writer {
	return Writer{
		repository: repository,
	}
}
func (r Writer) Create(params product.Entity) (string, error) {
	return r.repository.Create(params)
}
