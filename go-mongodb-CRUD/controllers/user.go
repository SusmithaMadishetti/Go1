package controllers

import (
	"encoding/json"
	"fmt"
	"html/template"
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

//Index export
func (uc UserController) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	tm := template.Must(template.New("index").ParseGlob("templates/*.html"))
	tm.ExecuteTemplate(w, "sush.html", nil)

	// s := `<!DOCTYPE html>
	// <html>
	// <head>
	// 	<meta charset='utf-8'>
	// 	<meta http-equiv='X-UA-Compatible' content='IE=edge'>
	// 	<title>sample</title>
	// 	<meta name='viewport' content='width=device-width, initial-scale=1'>
	// 	<link rel='stylesheet' type='text/css' media='screen' href='main.css'>
	// 	<script src='main.js'></script>
	// </head>
	// <body>
	// 	<a href=/user/9"> GO TO: http://localhost:8080/user/9</a>
	// </body>
	// </html>`
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// w.Write([]byte(s))
}

//CreateUser export
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{}
	json.NewDecoder(r.Body).Decode(&u)
	u.Id = bson.NewObjectId()

	uc.session.DB("local").C("users").Insert(u)
	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}

//GetUser export
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	fmt.Println(id)
	oid := bson.ObjectIdHex(id)
	fmt.Println(oid)
	u := models.User{}
	if err := uc.session.DB("local").C("users").FindId(oid).One(&u); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s\n", uj)
}

//GetUserByName
func (uc UserController) GetUserByName(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	name := p.ByName("name")
	// if !bson.IsObjectIdHex(id) {
	// 	w.WriteHeader(http.StatusNotFound)
	// 	return
	// }
	// fmt.Println(id)
	// oid := bson.ObjectIdHex(id)
	// fmt.Println(oid)
	u := []models.User{}
	if err := uc.session.DB("local").C("users").Find(bson.M{"name": name}).All(&u); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s\n", uj)
}

//GetUser export
func (uc UserController) GetAllUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	u := []models.User{}
	if err := uc.session.DB("local").C("users").Find(nil).All(&u); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s\n", uj)
}

//DeleteUser export
func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	oid := bson.ObjectIdHex(id)
	if err := uc.session.DB("local").C("users").RemoveId(oid); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "deleted user", oid, "\n")
}

//UpdateUser export
func (uc UserController) UpdateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//grab id
	id := p.ByName("id")
	//verify id is objectid
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	//object id
	oid := bson.ObjectIdHex(id)
	fmt.Println(oid)
	// json to u type
	u := models.User{}
	json.NewDecoder(r.Body).Decode(&u)

	uid := bson.M{"_id": oid}
	fmt.Println(uid)

	//uc.session.DB("test").C("users").UpdateId(u.Id, bson.M{"name": u.Name, "age": u.Age})
	uc.session.DB("local").C("users").Update(uid, bson.M{"name": u.Name, "age": u.Age})

	//object id
	// oid := bson.ObjectIdHex(id)
	// fmt.Println(oid)
	//	u.Id = bson.NewObjectId()

	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)

}
