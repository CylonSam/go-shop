package main

import (
	"net/http"

	"github.com/CylonSam/go-shop/routes"
	_ "github.com/lib/pq"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":4000", nil)
}
