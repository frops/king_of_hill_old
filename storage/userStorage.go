package userStorage

import (
    "github.com/frops/king_of_hill/models"
    "github.com/frops/king_of_hill/utils/mongo"
)

func Create (user models.User) error  {
    sess, db := mongo.Connection()
    defer sess.Close()

    return db.C("user").Insert(user)
}