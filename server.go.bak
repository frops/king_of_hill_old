package main

import (
	"github.com/frops/king_of_hill/models"
	"github.com/frops/king_of_hill/storage"
	"github.com/frops/king_of_hill/utils/config"
	"github.com/frops/king_of_hill/utils/response"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	config.Load()
	rtr := mux.NewRouter()
	rtr.HandleFunc("/app/", VVCount).Methods("GET")
	rtr.HandleFunc("/new", VVCountNew).Methods("GET")

	http.Handle("/", rtr)
	err := http.ListenAndServe(":"+config.Config.Server.Port, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("finish")
}

func VVCount(res http.ResponseWriter, req *http.Request) {
	user := models.User{
		Username: "fedor",
		State:    "new",
		Started:  1000,
	}
	response.Ok(res, user)
}

func VVCountNew(res http.ResponseWriter, req *http.Request) {
	user := models.User{
		Username: "fedor",
		State:    "new",
		Started:  1000,
	}

	err := userStorage.Create(user)

	if err != nil {
		fmt.Println(err)
		response.Ok(res, err)
	} else {
		response.Ok(res, user)
	}
}
