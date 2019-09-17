package models

import (
	"gopkg.in/mgo.v2/bson"
)

//Userinfo struct
type Userinfo struct {
	Name   string        `json:"name" bson:"name"`
	Gender string        `json:"gender" bson:"gender"`
	Age    int           `json:"age" bson:"age"`
	Id     bson.ObjectId `json:"id" bson:"_id"`
}

//Userintrest struct
type Userintrest struct {
	Name  string `json:"name" bson:"name"`
	Food  string `json:"food" bson:"food"`
	Hobby string `json:"hobby" bson:"hobby"`
}
