package product

import (
	"database/sql/driver"
	"errors"
	"gorm.io/gorm"
	"strings"
)

type (
	Entity struct {
		ID            string  `json:"id,omitempty"`
		Title         string  `json:"title"`
		Description   string  `json:"description" gorm:"size:5000"`
		Price         float64 `json:"price"`
		PreviousPrice float64 `json:"previousPrice"`
		Rate          float64 `json:"rate"`
		ShippingPrice float64 `json:"shippingPrice"`
		Stock         uint64  `json:"stock"`
		ImgURL        string  `json:"imgUrl"`
		Images        Images  `json:"images" gorm:"size:10000"`
	}
	Images []string
)

func (Entity) TableName() string {
	return "products"
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&Entity{})
}

func (i Images) Value() (driver.Value, error) {
	return strings.Join(i, ";;;"), nil
}

func (i *Images) Scan(data interface{}) error {
	strInBytes, ok := data.([]byte)
	if !ok {
		return errors.New("fail in scan data")
	}
	*i = strings.Split(string(strInBytes), ";;;")
	return nil
}
