package main

import (
	"net/http"

	"gopkg.in/mgo.v2"

	"./controllers"
	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	//uc := controllers.NewUserIntrestController(getSession())
	go r.POST("/user", uc.CreateUser)

	http.ListenAndServe(":8088", r)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	return s
}
