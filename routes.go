package main

import "net/http"

var routes = Routes{
	{
		Pattern: "/sign-up",
		Handler: SignUpHandler,
		Method: "POST",
	},
}

type Routes []Route
type Route struct {
	Pattern string
	Handler http.HandlerFunc
	Method string
}
