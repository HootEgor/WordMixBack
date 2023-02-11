package Models

import "errors"

type User struct {
	ID       int
	Login    string
	Password string
	Language int
}

func AddNewUser(login string, password string, language int) {
	newUser := User{
		ID:       len(Users),
		Login:    login,
		Password: password,
		Language: language,
	}
	Users = append(Users, newUser)
}

func GetUserByID(id int) (User, error) {
	var errUser User
	for i := 0; i < len(Users); i++ {
		if Users[i].ID == id {
			return Users[i], nil
		}
	}
	return errUser, errors.New("user not found")
}

var user1 = User{ID: 0, Login: "ro1ot", Password: "ro1ot", Language: 0}
var user2 = User{ID: 1, Login: "ro2ot", Password: "root4", Language: 3}

var Users = []User{user1, user2}
