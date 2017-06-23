package services

import (
    "github.com/frops/king_of_hill/models"
    "github.com/frops/king_of_hill/storage"
)

func Register(login string, password string) models.User {
    user := models.User{
        Username: login,
        Password: password,
        State: "new",
    }

    if userStorage.Exists(user) {

    }

    err := userStorage.Create(&user)

    if err != nil {
        panic(err)
    }

    return user
}

func IsValidUserToRegister (login string, password string) {
    user := models.User{
        Username: login,
        Password: password,
    }

    if userStorage.Exists(user) {
        return
    }


}
