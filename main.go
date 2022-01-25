package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequests() {
	myrouter := mux.NewRouter().StrictSlash(true)

	initMigration()
	myrouter.HandleFunc("/", hello)
	myrouter.HandleFunc("/Get", getplacement).Methods("GET")
	myrouter.HandleFunc("/Post", postplacement).Methods("POST")
	myrouter.HandleFunc("/Update", putplacement).Methods("PUT")
	myrouter.HandleFunc("/Delete", deleteplacement).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", myrouter))
}
func main() {
	handleRequests()
}
