package Services

import (
	Models "WordMixBack/src/Model"
	"cloud.google.com/go/firestore"
	"context"
	"google.golang.org/api/iterator"
	"log"
	"sort"
)

type Service interface {
	AddNewUser(client *firestore.Client, user Models.User) (string, error)
	LoginUser(client *firestore.Client, user Models.User) (string, error)
	GetUserInfo(client *firestore.Client, id string) (Models.User, error)
	GetLeaders(client *firestore.Client) ([]Models.Score, error)
	GetUserHistory(client *firestore.Client, id string) ([]Models.Score, error)
	NewUserScore(client *firestore.Client, score Models.Score) error
}

type handler struct {
	service Service
}

func AddNewUser(client *firestore.Client, user Models.User) (string, error) {
	ctx := context.Background()
	userRef, _, err := client.Collection("Users").Add(ctx, user)
	//var handler handler
	//userRef, err := handler.service.GetUser()

	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
		return "", err
	}

	return userRef.ID, nil
}

func LoginUser(client *firestore.Client, user Models.User) (string, error) {
	ctx := context.Background()

	collection := client.Collection("Users")

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

	return newUser, nil
}

func GetLeaders(client *firestore.Client) ([]Models.Score, error) {
	var scores []Models.Score
	ctx := context.Background()
	iter := client.Collection("Score").OrderBy("Score", firestore.Desc).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var score Models.Score
		if err := doc.DataTo(&score); err != nil {
			return nil, err
		}
		scores = append(scores, score)
	}

	sort.Slice(scores, func(i, j int) bool {
		return scores[i].Score > scores[j].Score
	})

	return scores, nil
}

func GetUserHistory(client *firestore.Client, id string) ([]Models.Score, error) {
	var scores []Models.Score
	ctx := context.Background()
	iter := client.Collection("Score").Where("UserID", "==", id).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var score Models.Score
		if err := doc.DataTo(&score); err != nil {
			return nil, err
		}
		scores = append(scores, score)
	}

	return scores, nil
}

func NewUserScore(client *firestore.Client, score Models.Score) error {
	ctx := context.Background()
	_, _, err := client.Collection("Score").Add(ctx, score)

	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
		return err
	}

	return nil
}
