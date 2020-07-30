package models

import (
	"driver"
	"entities"
	"fmt"
)

//ProductModel ...
type ProductModel struct {
}

//FindAll ..
func (*ProductModel) FindAll() ([]entities.Product, error) {
	db, err := driver.ConnectSQL()

	if err != nil {
		return nil, err
	}

	rows, err2 := db.Query("select id, price_lv0, product_avarta from products limit 10")

	if err2 != nil {
		return nil, err2
	}

	var products []entities.Product

	for rows.Next() {
		var product entities.Product
		rows.Scan(
			&product.ID,
			&product.PriceLv0,
			&product.ProductAvarta,
		)
		products = append(products, product)
	}

	fmt.Println(products)

	defer db.Close()
	return products, nil
}
