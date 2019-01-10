package main

import "net/http"

func (s *server) handleGetServers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		 _, _ = w.Write([]byte("Hello World!"))
	}
}

func (s *server) handlePostServers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Hello World!"))
	}
}

func (s *server) handleGetServer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Hello World!"))
	}
}

func (s *server) handlePutServer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Hello World!"))
	}
}