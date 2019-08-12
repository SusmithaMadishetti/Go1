package main

import (
	"database/sql"
	_ "fmt"
	"log"
	"net/http"
	"net/smtp"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type Student struct {
	ID, Number  int
	Name, Email string
}

type BasicResponse struct {
	Status_code   int
	Msg           string
	Response_data []Student
}

var templates = template.Must(template.ParseFiles("templates/register.html", "templates/login.html", "templates/welcome.html", "templates/error.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, br *BasicResponse) {
	// p.Datasql = template.HTML()

	err := templates.ExecuteTemplate(w, tmpl+".html", br)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func registerHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")
	renderTemplate(w, "register", &BasicResponse{})
	//fmt.Println(results)
	//http.Redirect(w,r,/login,http.StatusFound)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	emailid := r.FormValue("email")
	password := r.FormValue("password")
	db, err := sql.Open("mysql", "root:Alb@ny1995@tcp(localhost:3306)/demodb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	result, err := db.Query("select * from student where email= '" + emailid + "' ")
	var name, email, password_hash string
	for result.Next() {
		result.Scan(&name, &email, &password_hash)
	}
	answer := CheckPasswordHash(password, password_hash)
	if answer == true {
		w.Header().Set("Content-Type", "text/html")
		renderTemplate(w, "welcome", &BasicResponse{})
	}
	w.Header().Set("Content-Type", "text/html")
	renderTemplate(w, "error", &BasicResponse{})

}
func CheckPasswordHash(password, password_hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password_hash), []byte(password))
	return err == nil
}

func registerUserHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	email := r.FormValue("email")

	// Validation

	db, err := sql.Open("mysql", "root:Alb@ny1995@tcp(localhost:3306)/demodb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	password := name + "test"
	password_hash, _ := HashPassword(password)
	stmtIns, err := db.Query("INSERT INTO student (name, email, password_hash) VALUES ('" + name + "','" + email + "','" + password_hash + "')")

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close()
	//p := &Pagedata{Datasql: resultsRow, Userdata: results}
	send(password, email)

	response_data := &BasicResponse{Status_code: 1, Msg: "registered successfully!!Please check your mail for credentials."}
	// TODO   generate password, insert it into DB, send an email.
	w.Header().Set("Content-Type", "text/html")
	renderTemplate(w, "login", response_data)
	//fmt.Println(results)
}

// Utils
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func send(body, to string) {
	from := "susmithamadishetti54@gmail.com"
	pass := "Alb@ny1995"
	// to := "pavanjuttada@gmail.com"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Go lang master minders\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	log.Print("sent, It's time to go ..!")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", registerHandler)
	r.HandleFunc("/save", registerUserHandler)
	r.HandleFunc("/loginsubmit", loginHandler)
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	http.ListenAndServe(":8001", r)
}
