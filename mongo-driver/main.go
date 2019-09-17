package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

//Product structure
type Product struct {
	Name string `json:"name" bson:"name"`

	Price int           `json:"price" bson:"price"`
	Id    bson.ObjectId `json:"id" bson:"_id"`
}

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)

	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	collection := client.Database("local").Collection("products")
	//insert one
	pid := bson.NewObjectId()
	insertone := Product{"cookies", 15, pid}
	fmt.Println(insertone)
	insertResult, err := collection.InsertOne(context.TODO(), insertone)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult, pid)

	//insertmany
	qid := bson.NewObjectId()
	rid := bson.NewObjectId()

	icecream := Product{"icecream", 20, qid}
	nutella := Product{"nutella", 25, rid}
	insertmany := []interface{}{icecream, nutella}
	insertManyResult, err := collection.InsertMany(context.TODO(), insertmany)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)

	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connection to MongoDB closed.!")

}
