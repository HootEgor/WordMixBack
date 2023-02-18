package Models

type User struct {
	Login    string `firestore:"Login"`
	Password string `firestore:"Password"`
	Language int64  `firestore:"Language"`
}

//func GetUserByID(id int) (User, error) {
//	var errUser User
//	for i := 0; i < len(Users); i++ {
//		if Users[i].ID == id {
//			return Users[i], nil
//		}
//	}
//	return errUser, errors.New("user not found")
//}
//
//var user1 = User{Login: "ro1ot", Password: "ro1ot", Language: 0}
//var user2 = User{Login: "ro2ot", Password: "root4", Language: 3}
//
//var Users = []User{user1, user2}
