package main

import (
	Handlers "WordMixBack/src/Controllers"
	"WordMixBack/src/Repositories/UserRepository"
	"WordMixBack/src/Repositories/WordRepository"
	"WordMixBack/src/Services/AuthService"
	"WordMixBack/src/Services/UserService"
	"WordMixBack/src/Services/WordService"
	"context"
	firebase "firebase.google.com/go"
	"github.com/gorilla/mux"
	"google.golang.org/api/option"
	"log"
	"net/http"
)

func main() {

	Ctx := context.Background()
	serviceAcc := option.WithCredentialsFile("ServiceAccount/wordmixdatabase-firebase-adminsdk-b3nak-f88292c0fe.json")
	app, err := firebase.NewApp(Ctx, nil, serviceAcc)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(Ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	wordRepository := WordRepository.NewWordRepository(client)
	authRepository := UserRepository.NewUserRepository(client)
	userRepository := UserRepository.NewUserRepository(client)
	authService := AuthService.NewAuthService(authRepository)
	userService := UserService.NewUserService(userRepository)
	wordService := WordService.NewWordService(wordRepository)
	handlers := Handlers.NewHttpHandler(userService, wordService, authService)

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/score/leaders", handlers.GetLeaders).Methods("GET")
	myRouter.HandleFunc("/score", handlers.NewUserScore).Methods("POST")
	myRouter.HandleFunc("/user/score/{id}", handlers.GetUserHistory).Methods("GET")
	myRouter.HandleFunc("/user/{id}", handlers.GetUserInfo).Methods("GET")
	myRouter.HandleFunc("/auth/register", handlers.Register).Methods("POST")
	myRouter.HandleFunc("/auth/login", handlers.Login).Methods("POST")
	myRouter.HandleFunc("/word/add", handlers.AddNewWords).Methods("POST")
	myRouter.HandleFunc("/word/get", handlers.GetWords).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}
