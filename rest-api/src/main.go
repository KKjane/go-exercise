package main

import (
	"fmt"
	"net/http"
	person_api "rest-api/rest-api/apis/person_api"

	"github.com/gorilla/mux"
)

func index(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(responseWriter, "Welcome to GO Restful API")
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", index)
	router.HandleFunc("/api/persons/findall", person_api.FindAll)

	err := http.ListenAndServe(":5000", router)
	if err != nil {
		fmt.Println(err)
	}

}
