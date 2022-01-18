package product

import "gorm.io/gorm"

type (
	Entity struct {
		ID          string
		Description string
	}
)

func (Entity) TableName() string {
	return "products"
}
func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&Entity{})
}
