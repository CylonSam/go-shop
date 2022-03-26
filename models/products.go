package models

import (
	"github.com/CylonSam/go-shop/db"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func GetAllProducts() []Product {
	db := db.ConnectToDB()

	allProducts, err := db.Query("select * from products order by id asc")

	if err != nil {
		panic(err)
	}

	p := Product{}
	products := []Product{}

	for allProducts.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = allProducts.Scan(&id, &name, &description, &price, &quantity)

		if err != nil {
			panic(err)
		}

		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)
	}
	defer db.Close()
	return products
}

func CreateNewProduct(name, description string, price float64, quantity int) {
	db := db.ConnectToDB()

	insertProductIntoDB, err := db.Prepare("insert into products(name, description, price, quantity) values ($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	insertProductIntoDB.Exec(name, description, price, quantity)

	defer db.Close()
}

func DeleteProduct(id string) {
	db := db.ConnectToDB()

	deleteProduct, err := db.Prepare("delete from products where id=$1")

	if err != nil {
		panic(err.Error())
	}

	deleteProduct.Exec(id)
	defer db.Close()
}

func GetProduct(id string) Product {
	db := db.ConnectToDB()

	productFromDB, err := db.Query("select * from products where id=$1", id)

	if err != nil {
		panic(err.Error())
	}

	productToUpdate := Product{}

	for productFromDB.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err := productFromDB.Scan(&id, &name, &description, &price, &quantity)

		if err != nil {
			panic(err.Error())
		}

		productToUpdate.Id = id
		productToUpdate.Name = name
		productToUpdate.Description = description
		productToUpdate.Price = price
		productToUpdate.Quantity = quantity

	}
	defer db.Close()

	return productToUpdate
}

func UpdateProduct(id, quantity int, name, description string, price float64) {
	db := db.ConnectToDB()

	updateProduct, err := db.Prepare("update products set name=$1, description=$2, price=$3, quantity=$4 where id=$5")

	if err != nil {
		panic(err.Error())
	}

	updateProduct.Exec(name, description, price, quantity, id)

	defer db.Close()
}
