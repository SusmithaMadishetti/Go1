package controllers

import (
	"context"
	"fmt"
	"log"

	"gopkg.in/mgo.v2/bson"
)

func insert() {
	pId := bson.NewObjectId()
	cookie := Product{"cookies", 15, pId}
	fmt.Println(cookie)
	insertResult, err := collection.InsertOne(context.TODO(), cookie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult, pId)
}
