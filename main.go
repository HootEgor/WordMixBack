package main

import (
	"fmt"
	"log"
)

type config struct {
	port int
}

type app struct {
	config   config
	infolog  *log.Logger
	errorlog *log.Logger
}

type user struct {
	id       int
	login    string
	password string
	language int
}

type score struct {
	id       int
	score    int
	language int
	idUser   int
}

var idUser = 0
var users = make([]user, 0)

func main() {

	var cfg config

	addUser("Nikita", "123")
	addUser("Egor", "321")

	fmt.Print(users)
}

func addUser(login string, password string) {
	newUser := user{login: login, password: password}
	newUser.id = idUser
	idUser++
	newUser.language = 0
	users = append(users, newUser)
}
