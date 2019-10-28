package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
		name := query.Get("name")
	if name == "" {
		name = "Guest"
	}
	log.Printf("Received request for %s\n", name)
	w.Write([]byte(fmt.Sprintf("Hello, %s\n", name)))


}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	http.ListenAndServe(":8000", r)

}
