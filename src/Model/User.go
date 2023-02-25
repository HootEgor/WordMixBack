package Models

type User struct {
	Login    string `firestore:"Login"`
	Password string `firestore:"Password"`
}
