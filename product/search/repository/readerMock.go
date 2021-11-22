package repository

import (
	"simpleecommerceproductapi/product"
	"simpleecommerceproductapi/product/search"
)

type (
	ReaderMock struct {
	}
)

func NewReaderMock() ReaderMock {
	return ReaderMock{}
}

func (r ReaderMock) GetByParams(params search.Params) ([]product.Entity, error) {
	return []product.Entity{
		{
			ID:          "sadalskdaj-asdkasdasd=askads=laskdasd",
			Description: "Celular Sansung Galax EX",
		},
		{
			ID:          "sadalskdaj-asdkasdasd=askads=laskdasd",
			Description: "Celular Sansung Galax S",
		},
		{
			ID:          "sadalskdaj-asdkasdasd=askads=laskdasd",
			Description: "Celular Sansung Galax S 10",
		},
		{
			ID:          "sadalskdaj-asdkasdasd=askads=laskdasd",
			Description: "Celular Sansung Galax XPTO",
		},
		{
			ID:          "sadalskdaj-asdkasdasd=askads=laskdasd",
			Description: "Celular LG EX",
		},
		{
			ID:          "sadalskdaj-asdkasdasd=askads=laskdasd",
			Description: "Celular Motorola EX",
		},
	}, nil
}

func (r ReaderMock) GetByID(id string) (product.Entity, error) {
	return product.Entity{
		ID:          id,
		Description: "Celular Motorola EX",
	}, nil
}
