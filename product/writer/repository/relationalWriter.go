package repository

import (
	"gorm.io/gorm"
	"simpleecommerceproductapi/product"
)

type (
	RelationalWriter struct {
		db *gorm.DB
	}
)

func NewRelationalWriter(db *gorm.DB) RelationalWriter {
	return RelationalWriter{db: db}
}

func (w RelationalWriter) Create(entity product.Entity) (string, error) {
	err := w.db.Create(&entity).Error
	if err != nil {
		return "", err
	}
	return entity.ID, err
}
