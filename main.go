package main

import (
	"fmt"
	"net/http"

	"github.com/groozin/sailor/routes"
)

func main() {
	fmt.Println("Starting API on port 8080...")
	router := routes.RegisterRoutes()
	http.ListenAndServe(":8080", router)
}
