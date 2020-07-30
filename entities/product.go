package entities

import "database/sql"

//Product Định nghĩ một kiểu dữ liệu trong golang
type Product struct {
	ID            int            `json:"id" db:"id"`
	ProductName   sql.NullString `json:"name"`
	PriceLv0      int            `json:"price"`
	ProductAvarta string         `json:"images"`
}
