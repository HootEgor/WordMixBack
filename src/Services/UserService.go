package Services

import (
	Models "WordMixBack/src/Model"
	"cloud.google.com/go/firestore"
	"context"
	"google.golang.org/api/iterator"
	"log"
	"sort"
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

func LoginUser(client *firestore.Client, user Models.User) ([]Models.User, error) {
	ctx := context.Background()

	collection := client.Collection("Users")

	query := collection.Where("Login", "==", user.Login).Where("Password", "==", user.Password)

	var users []Models.User
	iter := query.Documents(ctx)
	defer iter.Stop()
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		// Convert the Firestore document to a User struct
		var user Models.User
		err = doc.DataTo(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if len(users) != 1 {
		return nil, nil
	}

	return users, nil
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

//func GetLeaders(client *firestore.Client) (Models.Score, error) {
//	ctx := context.Background()
//	id := "LZpcuQMrXkFiL72QhHkZ"
//	dsnap, err := client.Collection("Score").Doc(id).Get(ctx)
//	newScore := Models.Score{}
//	if err != nil {
//		return newScore, err
//	}
//	m := dsnap.Data()
//	newScore.Language = m["Language"].(int64)
//	newScore.Score = m["Score"].(int64)
//	newScore.UserID = m["UserID"].(string)
//
//	return newScore, nil
//}

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
