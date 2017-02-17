package main

import (
	"./utils/config"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	config.Load()

	err := http.ListenAndServe(":"+config.Config.Server.Port, nil)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/comments/get", handleGetComment)

	fmt.Println("finish")
}
