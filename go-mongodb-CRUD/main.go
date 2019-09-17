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
	r.GET("/", uc.Index)
	r.GET("/user/:id", uc.GetUser)
	r.GET("/users/:name", uc.GetUserByName)
	r.GET("/users", uc.GetAllUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	r.PUT("/users/:id", uc.UpdateUser)
	http.ListenAndServe(":8000", r)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	return s
}
