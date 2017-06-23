package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/frops/king_of_hill/models"
	"github.com/shagtv/go-api/httperrors"
	"github.com/shagtv/go-api/library/response"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
    "github.com/frops/king_of_hill/services"
)

var mySigningKey = []byte("secret")

var userAccess = map[string]string{
	"frops": "12345",
	"timur": "456",
}

// Index - main controller
var Index = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome")
})

// TodoIndex -list of todos
var TodoIndex = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//var token jwt.Token
	token := r.Context().Value("user").(*jwt.Token)
	//claims := token.Claims.(models.UserClaims)
	cl := token.Claims.(jwt.MapClaims)

	fmt.Println(cl["username"])

	//	fmt.Println(r.Context().Value("Claims"))

	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}

	response.Ok(w, todos)
})

// TodoShow - show scpecify todo
var TodoShow = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID := vars["todoID"]
	fmt.Fprintln(w, "Todo show:", todoID)
})

// AuthGetToken sd sds
var AuthGetToken = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	var login = r.FormValue("login")
	var password = r.FormValue("password")

	if !validateUser(login, password) {
		error := httperrors.New("Неверный логин/пароль", 401)
		response.Error(w, error)
	} else {

		exp := time.Now().Add(time.Hour * 24).Unix()
		claim := createEmptyClaim(login, exp)
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
		tokenString, _ := token.SignedString(mySigningKey)
		json := map[string]string{"token": tokenString}

		response.Ok(w, json)
	}
})

// AuthSignup register user
var AuthSignup = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()

    var login = r.FormValue("login")
    var password = r.FormValue("password")

    user := services.Register(login, password)

    exp := time.Now().Add(time.Hour * 24).Unix()
    claim := createEmptyClaim(user.Username, exp)
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
    tokenString, _ := token.SignedString(mySigningKey)
    json := map[string]string{"token": tokenString}

    response.Ok(w, json)
})

func validateUser(login string, password string) bool {
	if val, ok := userAccess[login]; ok {
		if val == password {
			return true
		}
	}

	return false
}

func createEmptyClaim(login string, expireToken int64) models.UserClaims {
	return models.UserClaims{
		login,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "king_of_hill",
		},
	}
}
