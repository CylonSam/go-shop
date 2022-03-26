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

func Edit(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	product := models.GetProduct(productId)

	templates.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		idConvInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Id conversion error: ", err)
		}

		priceConvFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Price conversion error: ", err)
		}

		quantityConvInt, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Quantity conversion error: ", err)
		}

		models.UpdateProduct(idConvInt, quantityConvInt, name, description, priceConvFloat)
	}

	http.Redirect(w, r, "/", 301)
}
