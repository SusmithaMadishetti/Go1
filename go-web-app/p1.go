package main

import (
	"net/http"
	"github.com/gorilla/mux"
)

func someHandler(w http.ResponseWriter, r *http.Request) {
	c := http.Cookie{
		Name:  "theme",
		Value: "dark",
	}
	http.SetCookie(w, &c)
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", someHandler)
	http.ListenAndServe(":8000", r)

}
