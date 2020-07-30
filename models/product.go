package models

import (
	"driver"
	"entities"
)

//ProductModel ...
type ProductModel struct {
}

//FindAll ..
func (*ProductModel) FindAll(limit, offset string) ([]entities.Product, error) {
	db, err := driver.ConnectSQL()

	if err != nil {
		return nil, err
	}

	rows, err2 := db.Query("select id, product_name, price_lv0, product_avarta from products limit ? offset ?", limit, offset)

	if err2 != nil {
		return nil, err2
	}

	var products []entities.Product

	for rows.Next() {
		var product entities.Product
		rows.Scan(
			&product.ID,
			&product.ProductName,
			&product.PriceLv0,
			&product.ProductAvarta,
		)
		products = append(products, product)
	}

	defer db.Close()
	return products, nil
}
