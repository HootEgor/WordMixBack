package Models

type Score struct {
	Language int64  `firestore:"Language"`
	Score    int64  `firestore:"Score"`
	UserID   string `firestore:"UserID"`
}
