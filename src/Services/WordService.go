package Services

import (
	Models "WordMixBack/src/Model"
	"cloud.google.com/go/firestore"
	"context"
	"google.golang.org/api/iterator"
)

func AddNewWords(client *firestore.Client, words []Models.Word) ([]Models.Word, error) {
	var addedWords []Models.Word
	ctx := context.Background()
	for _, word := range words {
		query := client.Collection("Words").Where("Word", "==", word.Word).Limit(1)
		docSnap, err := query.Documents(ctx).GetAll()
		if err != nil {
			continue
		}

		if len(docSnap) == 0 {
			_, _, err = client.Collection("Words").Add(ctx, word)
			if err != nil {
				return nil, err
			}
			addedWords = append(addedWords, word)
		}
	}
	return addedWords, nil
}

func GetWords(client *firestore.Client) ([]Models.Word, error) {
	ctx := context.Background()
	var words []Models.Word
	iter := client.Collection("Words").Documents(ctx)
	defer iter.Stop()
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var word Models.Word
		err = doc.DataTo(&word)
		if err != nil {
			return nil, err
		}
		words = append(words, word)
	}
	return words, nil
}
