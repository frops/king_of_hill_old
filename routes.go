package main

import (
	"net/http"

	"github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

type Route struct {
	Name         string
	Method       string
	Pattern      string
	HandlerFunc  http.HandlerFunc
	AuthRequired bool
}

// Routes is map for requests
type Routes []Route

// NewRouter creates maps of routes
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var h http.Handler

		if route.AuthRequired {
			h = jwtMiddleware.Handler(route.HandlerFunc)
		} else {
			h = route.HandlerFunc
		}

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(h)
	}

	return router
}

var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	},

	SigningMethod: jwt.SigningMethodHS256,
})

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
		false,
	},
	Route{
		"TodoIndex",
		"GET",
		"/todos",
		TodoIndex,
		true,
	},
	Route{
		"AuthGetToken",
		"POST",
		"/auth/get-token",
		AuthGetToken,
		false,
	},
}
