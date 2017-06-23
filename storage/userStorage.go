package userStorage

import (
    "github.com/frops/king_of_hill/models"
    "github.com/frops/king_of_hill/libs"
    "log"
    "gopkg.in/mgo.v2/bson"
    "fmt"
)

// Create sdasd
func Create(user *models.User) error {
    session := libs.Open()
    defer session.Close()

    c := session.DB("king_of_hill").C("users")
    err := c.Insert(&user)

    if err != nil {
        log.Fatal(err)
        panic(err)
    }

    result := models.User{}

    err = c.Find(bson.M{"username": user.Username}).One(&result)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("user inserted")

    return nil
}

func Exists(user *models.User) error {
    session := libs.Open()
    defer session.Close()

    c := session.DB("king_of_hill").C("users")

    foundUser := models.User{}
    err := c.Find(bson.M{"username": user.Username}).One(&foundUser)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(foundUser)
}