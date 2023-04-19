package Repositories

import (
	Models "WordMixBack/src/Model"
	"context"
	"google.golang.org/api/iterator"
)

func (o *Repository) AddNewWords(ctx context.Context, words []Models.Word) ([]Models.Word, error) {
	var addedWords []Models.Word
	for _, word := range words {
		query := o.c.Collection("Words").Where("Word", "==", word.Word).Limit(1)
		docSnap, err := query.Documents(ctx).GetAll()
		if err != nil {
			continue
		}

		if len(docSnap) == 0 {
			_, _, err = o.c.Collection("Words").Add(ctx, word)
			if err != nil {
				return nil, err
			}
			addedWords = append(addedWords, word)
		}
	}
	return addedWords, nil
}

func (o *Repository) GetWords(ctx context.Context) ([]Models.Word, error) {
	var words []Models.Word
	iter := o.c.Collection("Words").Documents(ctx)
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
