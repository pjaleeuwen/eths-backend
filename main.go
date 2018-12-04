package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", IndexHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8081", router))
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("true"))
	return
}