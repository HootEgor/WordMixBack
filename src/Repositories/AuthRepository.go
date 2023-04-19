package Repositories

import (
	Models "WordMixBack/src/Model"
	"cloud.google.com/go/firestore"
	"context"
	"google.golang.org/api/iterator"
	"log"
)

type Repository struct {
	c *firestore.Client
}

func NewRepository(c *firestore.Client) *Repository {
	return &Repository{c}
}

func (o *Repository) AddNewUser(ctx context.Context, user Models.User) (string, error) {
	userRef, _, err := o.c.Collection("Users").Add(ctx, user)

	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
		return "", err
	}

	return userRef.ID, nil
}

func (o *Repository) GetUserByInfo(ctx context.Context, user Models.User) (string, error) {
	collection := o.c.Collection("Users")
	query := collection.Where("Login", "==", user.Login).Where("Password", "==", user.Password)
	userID := ""
	iter := query.Documents(ctx)
	defer iter.Stop()
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return "", err
		}
		userID = doc.Ref.ID
		break
	}
	return userID, nil
}
