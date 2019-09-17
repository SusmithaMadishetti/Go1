package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"../models"
	"github.com/julienschmidt/httprouter"
)

//UserInfoController export
type UserController struct {
	session *mgo.Session
}

//NewUserInfoController export
func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

//CreateUser export
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.Userinfo{}
	json.NewDecoder(r.Body).Decode(&u)
	u.Id = bson.NewObjectId()

	uc.session.DB("User").C("userinfo").Insert(u)
	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}

//CreateUserIntrest export
func (uc UserController) CreateUserIntrest(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	u := models.Userintrest{}
	json.NewDecoder(r.Body).Decode(&u)
	uc.session.DB("User").C("userintrest").Find(Userinfo.name)
	uj, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}
