package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var signUpRequest signUpRequest
	err := decoder.Decode(&signUpRequest)
	if err != nil {
		panic(err)
	}
	log.Println(signUpRequest)

	w.WriteHeader(200)
	w.Write([]byte("{}"))
	return
}

type signUpRequest struct {
	Username string
	Email string
	EmailConfirm string
	Password string
	PasswordConfirm string
}