package main

import (
	"database/sql"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Student struct {
	ID, Number  int
	Name, Email string
}

var templates = template.Must(template.ParseFiles("templates/register.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, msg string) {
	// p.Datasql = template.HTML()

	err := templates.ExecuteTemplate(w, tmpl+".html", msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func registerHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")
	renderTemplate(w, "register", "")
	//fmt.Println(results)
	//http.Redirect(w,r,/login,http.StatusFound)
}

func registerUserHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	email := r.FormValue("email")

	db, err := sql.Open("mysql", "root:Alb@ny1995@tcp(localhost:3306)/demodb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	stmtIns, err := db.Query("INSERT INTO student (name, email) VALUES ('" + name + "','" + email + "')")

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close()

	// TODO   generate password, insert it into DB, send an email.
	w.Header().Set("Content-Type", "text/html")
	renderTemplate(w, "register", "Your registration is successful Plz chek your mail for the password")
	//fmt.Println(results)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", registerHandler)
	r.HandleFunc("/save", registerUserHandler)

	http.ListenAndServe(":8001", r)
}
