package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model

	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	// migrate schema
	db.AutoMigrate(&Product{})

	// create a new product
	db.Create(&Product{Code: "D42", Price: 100})

	// read from products list
	var product Product

	// finding first product with row 'code' equal to 'D42'
	db.First(&product, "code = ?", "D42")

	// update product price to 200
	db.Model(&product).Update("Price", 121)

	db.Model(&product).Updates(Product{Price: 121, Code: "F42"})
	db.Model(&product).Updates(map[string]interface{}{
		"Price": 121, "Code": "F42",
	})
}
