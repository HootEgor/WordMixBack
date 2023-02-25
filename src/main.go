package main

import (
	Handlers "WordMixBack/src/Controllers"
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"github.com/gorilla/mux"
	"google.golang.org/api/option"
	"log"
	"net/http"
)

func handleRequest(client *firestore.Client) {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/leaders", Handlers.GetLeaders(client)).Methods("GET")
	myRouter.HandleFunc("/user/{id}", Handlers.GetUserInfo(client)).Methods("GET")
	myRouter.HandleFunc("/auth/register", Handlers.Register(client)).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func InitFirebase() *firestore.Client {

	ctx := context.Background()
	serviceAcc := option.WithCredentialsFile("ServiceAccount/wordmixdatabase-firebase-adminsdk-b3nak-f88292c0fe.json")
	app, err := firebase.NewApp(ctx, nil, serviceAcc)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	return client
}

func main() {

	client := InitFirebase()
	defer client.Close()

	handleRequest(client)
}
