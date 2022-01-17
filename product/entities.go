package product

import "gorm.io/gorm"

type (
	Entity struct {
		ID          string
		Description string
	}
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&Entity{})
}
