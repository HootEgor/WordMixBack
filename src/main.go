package main

import (
	Handlers "WordMixBack/src/Controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", Handlers.GetUserInfo).Methods("GET")
	myRouter.HandleFunc("/leaders", Handlers.GetLeadersByRegion).Methods("GET")
	myRouter.HandleFunc("/auth/register", Handlers.Register).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	handleRequest()
}
