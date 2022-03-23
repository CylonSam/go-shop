package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func connectToDB() *sql.DB {
	connection := "user=sam dbname=goshop password=test123 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)

	if err != nil {
		panic(err)
	}

	return db
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":4000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	db := connectToDB()

	allProducts, err := db.Query("select * from products")

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

		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)
	}

	templates.ExecuteTemplate(w, "Index", products)
}
