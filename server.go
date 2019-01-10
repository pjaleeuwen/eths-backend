package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type server struct {
	port string
	db Database
	router *mux.Router
	email EmailSender
}

func (s *server) start() {
	s.routes()
	log.Fatal(http.ListenAndServe(s.port, s.router))
}

type EmailSender interface{}

