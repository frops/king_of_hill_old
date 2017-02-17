package mongo

import (
	"gopkg.in/mgo.v2"
	"../config"
	"log"
)

var Mongo *mgo.Session

func Connection() (*mgo.Session, *mgo.Database) {
	var err error

	if Mongo == nil {
		Mongo, err = mgo.Dial(config.Config.Mongo.Host + ":" + config.Config.Mongo.Port)

		if err != nil {
			log.Fatal(err)
		}
		Mongo.SetMode(mgo.Monotonic, true)
	}

    sess := Mongo.Copy()

	return sess, sess.DB(config.Config.Mongo.Database)
}
