package main

import (
	"database/sql"
	"fmt"
	_ "fmt"
	"log"
	"math/rand"
	"net/http"
	"net/smtp"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

//Student datatype
type Student struct {
	ID, Number  int
	Name, Email string
}

//BasicResponse structure
type BasicResponse struct {
	Statuscode int
	Msg        string
	Response   []Student
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
	fmt.Println(emailid)
	fmt.Println(password)
	db, err := sql.Open("mysql", "root:Alb@ny1995@tcp(localhost:3306)/demodb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	result, err := db.Query("SELECT * FROM student WHERE email='susmithamadishetti54@gmail.com'")
	fmt.Println(result)
	var name, email, passwordhash string
	for result.Next() {
		result.Scan(&name, &email, &passwordhash)
		fmt.Println(name)
		fmt.Println(email)
		fmt.Println(passwordhash)
		//fmt.Println(password)
	}

	answer := CheckPasswordHash(password, passwordhash)
	fmt.Println(answer)
	if answer == true {
		w.Header().Set("Content-Type", "text/html")
		renderTemplate(w, "welcome", &BasicResponse{})
	}
	if answer == false {
		w.Header().Set("Content-Type", "text/html")
		renderTemplate(w, "error", &BasicResponse{})
	}
}

//CheckPasswordHash checking password with database
func CheckPasswordHash(password, passwordhash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordhash), []byte(password))
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

	password := name + string(rand.Intn(1000))
	passwordhash, _ := HashPassword(password)
	stmtIns, err := db.Query("INSERT INTO student (name, email, passwordhash) VALUES ('" + name + "','" + email + "','" + passwordhash + "')")

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close()
	//p := &Pagedata{Datasql: resultsRow, Userdata: results}
	send(password, email)

	responsedata := &BasicResponse{Statuscode: 1, Msg: "registered successfully!!Please check your mail for credentials."}
	// TODO   generate password, insert it into DB, send an email.
	w.Header().Set("Content-Type", "text/html")
	renderTemplate(w, "login", responsedata)
	//fmt.Println(results)
}

// HashPassword Utils
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
