package main

import "github.com/gorilla/mux"

func main() {
	srv := &server{
		port: ":8080",
		db: database(),
		router: mux.NewRouter(),
		email: nil,
	}

	srv.start()
}