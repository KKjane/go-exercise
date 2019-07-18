package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	Title  string `json:"title"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
}

type Persons []Person

func returnPersons(responseWriter http.ResponseWriter, request *http.Request) {
	persons := Persons{
		Person{Title: "Mr.", Name: "Wachi", Gender: "M"},
		Person{Title: "Msr.", Name: "Hanako", Gender: "F"},
	}
	json.NewEncoder(responseWriter).Encode(persons)
}

func index(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(responseWriter, "Welcome to GO Restful API")
	fmt.Println("Access to index.")
}

func handleRequests() {
	mainRouter := mux.NewRouter().StrictSlash(true)
	mainRouter.HandleFunc("/", index)
	mainRouter.HandleFunc("/persons", returnPersons).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", mainRouter))
}

func main() {
	fmt.Println("API Server started.")
	handleRequests()
}
