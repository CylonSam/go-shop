package routes

import (
	"net/http"

	"github.com/CylonSam/go-shop/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
}
