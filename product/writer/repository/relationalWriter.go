package repository

import (
	"context"
	uuid "github.com/satori/go.uuid"
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

func (w RelationalWriter) Create(ctx context.Context, entity product.Entity) (string, error) {
	entity.ID = uuid.NewV4().String()
	err := w.db.WithContext(ctx).Create(&entity).Error
	if err != nil {
		return "", err
	}
	return entity.ID, err
}
