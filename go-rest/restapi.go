package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Expanse is struct
type Expanse struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Anount      float64 `json:"amount"`
}

var Expanses []Expanse

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welocome to home Page..!")
}

func getAllExpanses(w http.ResponseWriter, r *http.Request) {
	// TODO get it from DB
	json.NewEncoder(w).Encode(Expanses)
}

func getSingleExapnse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	for _, expanse := range Expanses {
		if expanse.Id == id {
			json.NewEncoder(w).Encode(expanse)
		}
	}
}
func deleteExpanse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	for index, expanse := range Expanses {
		if expanse.Id == id {
			//json.NewEncoder(w).Encode(expanse)
			Expanses = append(Expanses[:index], Expanses[index+1:]...)

		}
	}

}

func createExpanse(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)
	var expanse Expanse
	json.Unmarshal(reqBody, &expanse)
	Expanses = append(Expanses, expanse)
	json.NewEncoder(w).Encode(Expanses)

}

func updateExpanse(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)
	var expanse Expanse
	json.Unmarshal(reqBody, &expanse)
	//Expanses = append(Expanses, expanse)
	for index, expanse_old := range Expanses {
		if expanse.Id == expanse_old.Id {
			Expanses[index] = expanse
			json.NewEncoder(w).Encode(Expanses)
		}
	}

}

func requestHandler() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/expanses", getAllExpanses)
	router.HandleFunc("/expanse/{id}", getSingleExapnse).Methods("GET")
	router.HandleFunc("/expanse", createExpanse).Methods("POST")
	router.HandleFunc("/expanse", updateExpanse).Methods("PUT")
	router.HandleFunc("/expans/{id}", deleteExpanse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
func main() {
	fmt.Println("Starting the server @8000")

	Expanses = []Expanse{Expanse{Id: "1", Name: "Rent", Description: "For the mont Aug", Anount: 36000},
		Expanse{Id: "2", Name: "Power Bill", Description: "For the mont Aug", Anount: 3000.0}}
	requestHandler()
}
