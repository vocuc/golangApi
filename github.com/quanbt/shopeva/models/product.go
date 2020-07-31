package models

import (
	"database/sql"
)

//Product ...
type Product struct {
	ID     uint           `json:"id" gorm:"primary_key"`
	Name   sql.NullString `gorm:"column:product_name" json:"name"`
	Price  int            `gorm:"column:price_lv0" json:"price"`
	Images string         `gorm:"column:product_avarta" json:"images"`
}

//TableName Product table
func (Product) TableName() string {
	return "products"
}
