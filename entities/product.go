package entities

//Product Định nghĩ một kiểu dữ liệu trong golang
type Product struct {
	ID            int    `json:"id" db:"id"`
	ProductName   string `json:"name"`
	PriceLv0      int    `json:"price"`
	ProductAvarta string `json:"images"`
}
