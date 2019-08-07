package main

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type student struct {
	name   string
	email  string
	mobile int
}

//var temp:=template.Must(template.ParseFiles("welcome.html"))
func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("welcome to page") // console reply
	w.Header().Set("Content-Type","text/html")
	template,err:=parseTemplat
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", welcomeHandler)
	http.ListenAndServe(":8001", r)

}
