package controllers

import (
	"net/http"
	"text/template"

	"github.com/CylonSam/go-shop/models"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.GetAllProducts()
	templates.ExecuteTemplate(w, "Index", products)
}
