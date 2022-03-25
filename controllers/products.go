package controllers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/CylonSam/go-shop/models"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.GetAllProducts()
	templates.ExecuteTemplate(w, "Index", products)
}

func NewProduct(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "NewProduct", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		priceFloat64, err := strconv.ParseFloat(price, 64)

		if err != nil {
			log.Println("Could not convert price: ", err)
		}

		quantityInt, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Could not convert quantity: ", err)
		}

		models.CreateNewProduct(name, description, priceFloat64, quantityInt)
	}

	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")

	models.DeleteProduct(productId)

	http.Redirect(w, r, "/", 301)
}
