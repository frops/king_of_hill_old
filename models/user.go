package models

import (
    "gopkg.in/mgo.v2/bson"
)

type User struct {
    Id bson.ObjectId `db:"_id" json:"-" bson:"_id,omitempty"`
    Username string `db:"username" json:"username"`
    Password string `db:"password" json:"-"`
    State string `db:"state" json:"state"`
    Started uint64 `db:"started" json:"started"`
}
