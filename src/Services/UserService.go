package Services

import (
	Models "WordMixBack/src/Model"
	"cloud.google.com/go/firestore"
	"context"
	"log"
)

func AddNewUser(client *firestore.Client, user Models.User) error {
	ctx := context.Background()
	_, _, err := client.Collection("Users").Add(ctx, user)

	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
		return err
	}

	return nil
}

func GetUserInfo(client *firestore.Client, id string) (Models.User, error) {
	ctx := context.Background()
	dsnap, err := client.Collection("Users").Doc(id).Get(ctx)
	newUser := Models.User{}
	if err != nil {
		return newUser, err
	}
	m := dsnap.Data()
	newUser.Login = m["Login"].(string)
	newUser.Password = m["Password"].(string)
	newUser.Language = m["Language"].(int64)

	return newUser, nil
}
