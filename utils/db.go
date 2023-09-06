package utils

import (
	"gopkg.in/mgo.v2"
)

var DB *mgo.Database

const (
	hostNmae       string = "localhost:27017"
	dbName         string = "todo_list_go"
	CollectionName string = "todo"
)

func init() {
	sess, err := mgo.Dial(hostNmae)
	CheckError(err)
	sess.SetMode(mgo.Monotonic, true)
	DB = sess.DB(dbName)
}
