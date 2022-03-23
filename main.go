package main

import (
	"html/template"
	"net/http"
)

type Product struct {
	Name        string
	Description string
	Price       float64
	Quantity    int
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":4000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	products := []Product{
		{Name: "T-Shirt", Description: "Blue and beautiful", Price: 50, Quantity: 3},
		{"Sneakers", "Comfortable and stylish", 299, 5},
		{"Headphones", "Good!", 49, 10},
	}
	templates.ExecuteTemplate(w, "Index", products)
}
