package repository

import "simpleecommerceproductapi/product"

type (
	WriterMock struct {
	}
)

func NewWriterMock() WriterMock {
	return WriterMock{}
}

func (w WriterMock) Create(entity product.Entity) (string, error) {
	return entity.ID, nil
}
