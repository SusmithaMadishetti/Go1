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

var u = models.Userinfo{}
var ui = models.Userintrest{}
var uc = UserController{}

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
	go userinfo(u, w, r)
	go userintrest(ui, w, r, p)

}

func userinfo(u models.Userinfo, w http.ResponseWriter, r *http.Request) {
	//	u := models.Userinfo{}
	json.NewDecoder(r.Body).Decode(&u)
	u.Id = bson.NewObjectId()

	uc.session.DB("User").C("userinfo").Insert(u)
	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}
func userintrest(ui models.Userintrest, w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("name")
	//verify id is objectid
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	//object id
	oid := bson.ObjectIdHex(id)
	fmt.Println(oid)
	u := models.Userintrest{}
	json.NewDecoder(r.Body).Decode(&u)

	uid := bson.M{"_id": oid}
	fmt.Println(uid)
	//uc.session.DB("test").C("users").UpdateId(u.Id, bson.M{"name": u.Name, "age": u.Age})
	uc.session.DB("local").C("users").Update(uid, models.Userinfo{})

	uc.session.DB("User").C("userintrest").Insert(u)
	uj, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}
