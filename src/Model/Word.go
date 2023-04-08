package Models

type Word struct {
	Language int64  `firestore:"Language"`
	Word     string `firestore:"Word"`
}
