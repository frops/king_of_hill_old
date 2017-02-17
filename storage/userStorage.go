package userStorage

import (
    "../models"
    "../utils/mongo"
)

func Create (user models.User) error  {
    sess, db := mongo.Connection()
    defer sess.Close()

    return db.C("user").Insert(user)
}