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
		Title         string  `json:"title,omitempty"`
		Description   string  `json:"description,omitempty" gorm:"size:5000"`
		Price         float64 `json:"price,omitempty"`
		PreviousPrice float64 `json:"previousPrice,omitempty"`
		Rate          float64 `json:"rate,omitempty"`
		ShippingPrice float64 `json:"shippingPrice,omitempty"`
		Stock         uint64  `json:"stock,omitempty"`
		ImgURL        string  `json:"imgUrl,omitempty"`
		Images        Images  `json:"images,omitempty" gorm:"size:10000"`
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
