package repository

import (
	"database/sql"
	"gorm.io/gorm"
	"simpleecommerceproductapi/product"
	"simpleecommerceproductapi/product/search"
)

type (
	RelationalReader struct {
		db *gorm.DB
	}
)

const GetByDescriptionConditions = "description LIKE ?"
const GetByTitleConditions = "title LIKE ?"

func NewRelationalReader(db *gorm.DB) RelationalReader {
	return RelationalReader{db: db}
}
func (r RelationalReader) GetByParams(params search.Params) ([]product.Entity, error) {
	var products []product.Entity
	tx := r.db.Begin(&sql.TxOptions{ReadOnly: true})
	if params.Description != "" {
		tx = tx.Or(GetByDescriptionConditions, "%"+params.Description+"%")
	}
	if params.Title != "" {
		tx = tx.Or(GetByTitleConditions, "%"+params.Title+"%")
	}
	err := tx.Find(&products).Error
	return products, err
}

func (r RelationalReader) GetByID(id string) (product.Entity, error) {
	var p product.Entity
	err := r.db.First(&p, product.Entity{
		ID: id,
	}).Error
	return p, err
}
