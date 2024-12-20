package routes

import (
	"net/http"

	"github.com/groozin/sailor/internal/handlers"
)

func RegisterRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/is-prime", handlers.IsPrimeHandler)
	mux.HandleFunc("/add", handlers.AddHandler)
	mux.HandleFunc("/multiply", handlers.MultiplyHandler)
	mux.HandleFunc("/subtract", handlers.SubtractHandler)

	return mux
}
