package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func createRouter() *mux.Router {
	router := mux.NewRouter()

	for _, route := range routes {
		router.HandleFunc(route.Pattern, withCORS(route.Handler)).Methods(route.Method, "OPTIONS")
	}

	return router
}

func withCORS(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		setupResponse(&w)
		if r.Method == "OPTIONS" {
			return
		}

		h(w,r)
	}
}

func setupResponse(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "application/json; charset=utf-8")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}