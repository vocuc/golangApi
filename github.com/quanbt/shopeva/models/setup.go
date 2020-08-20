package models

import (
	_ "github.com/go-sql-driver/mysql" //...
	"github.com/jinzhu/gorm"
)

//DB ...
var DB *gorm.DB

//ConnectDataBase ...
func ConnectDataBase() {
	database, err := gorm.Open("mysql", "root:agrilavin@/shopeva?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("Failed to connect to database!")
	}

	DB = database
}
