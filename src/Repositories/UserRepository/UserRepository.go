package UserRepository

import (
	Models "WordMixBack/src/Model"
	"cloud.google.com/go/firestore"
	"context"
	"google.golang.org/api/iterator"
	"log"
	"sort"
)

type UserRepository struct {
	c *firestore.Client
}

func NewUserRepository(c *firestore.Client) *UserRepository {
	return &UserRepository{c}
}

func (o *UserRepository) GetUserInfo(ctx context.Context, id string) (Models.User, error) {
	dsnap, err := o.c.Collection("Users").Doc(id).Get(ctx)
	newUser := Models.User{}
	if err != nil {
		return newUser, err
	}
	m := dsnap.Data()
	newUser.Login = m["Login"].(string)
	newUser.Password = m["Password"].(string)

	return newUser, nil
}

func (o *UserRepository) GetLeaders(ctx context.Context) ([]Models.Score, error) {
	var scores []Models.Score
	iter := o.c.Collection("Score").OrderBy("Score", firestore.Desc).Documents(ctx)
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

func (o *UserRepository) GetUserHistory(ctx context.Context, id string) ([]Models.Score, error) {
	var scores []Models.Score
	iter := o.c.Collection("Score").Where("UserID", "==", id).Documents(ctx)
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

func (o *UserRepository) NewUserScore(ctx context.Context, score Models.Score) error {
	_, _, err := o.c.Collection("Score").Add(ctx, score)
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
		return err
	}
	return nil
}

func (o *UserRepository) AddNewUser(ctx context.Context, user Models.User) (string, error) {
	userRef, _, err := o.c.Collection("Users").Add(ctx, user)

	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
		return "", err
	}

	return userRef.ID, nil
}

func (o *UserRepository) GetUserByInfo(ctx context.Context, user Models.User) (string, error) {
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
