package repository

import (
	"gorm.io/gorm"
	"simpleecommerceproductapi/product"
	"simpleecommerceproductapi/product/search"
)

type (
	RelationalReader struct {
		db *gorm.DB
	}
)

const GetByParamsConditions = "description LIKE ?"

func NewRelationalReader(db *gorm.DB) RelationalReader {
	return RelationalReader{db: db}
}
func (r RelationalReader) GetByParams(params search.Params) ([]product.Entity, error) {
	var products []product.Entity
	err := r.db.Find(&products, GetByParamsConditions, "%"+params.Description+"%").Error
	return products, err
}

func (r RelationalReader) GetByID(id string) (product.Entity, error) {
	var p product.Entity
	err := r.db.First(&p, product.Entity{
		ID: id,
	}).Error
	return p, err
}
