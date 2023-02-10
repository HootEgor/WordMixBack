package main

import "fmt"

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

func main() {
	fmt.Print("Hello word")
}
