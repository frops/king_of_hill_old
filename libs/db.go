package libs

import "gopkg.in/mgo.v2"

func Open() *mgo.Session {
    session, err := mgo.Dial("localhost:27017")

    if err != nil {
        panic(err)
    }

    return session
}
