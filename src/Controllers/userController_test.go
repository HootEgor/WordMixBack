package Handlers

import (
	Models "WordMixBack/src/Model"
	"WordMixBack/src/Repositories"
	Services "WordMixBack/src/Services"
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"net/http"
	"net/http/httptest"
	"testing"
)

func initializeFirestoreClient() (*firestore.Client, error) {
	ctx := context.Background()
	serviceAcc := option.WithCredentialsFile("../ServiceAccount/wordmixdatabase-firebase-adminsdk-b3nak-f88292c0fe.json")
	config := &firebase.Config{
		ProjectID: "wordmixdatabase",
	}
	app, err := firebase.NewApp(ctx, config, serviceAcc)
	if err != nil {
		return nil, err
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func TestHandler_GetUserInfo(t *testing.T) {
	client, err := initializeFirestoreClient()
	assert.NoError(t, err)
	defer client.Close()

	router := mux.NewRouter()

	userRepository := Repositories.NewRepository(client)
	authRepository := Repositories.NewRepository(client)
	wordRepository := Repositories.NewRepository(client)
	userService := Services.NewUserService(authRepository, userRepository)
	wordService := Services.NewWordService(wordRepository)
	handler := NewHttpHandler(userService, wordService)

	router.HandleFunc("/users/{id}", handler.GetUserInfo)

	req, err := http.NewRequest("GET", "/users/GBoBWtIC51Agb7fIqXlP", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	expected := `{"Login":"pop","Password":"pop"}`
	assert.Equal(t, expected, rr.Body.String())
}

func TestHandler_NewUserScore(t *testing.T) {
	client, err := initializeFirestoreClient()
	assert.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	score := Models.Score{
		Language: 3,
		Score:    100,
		UserID:   "test-user-id",
	}

	userRepository := Repositories.NewRepository(client)
	authRepository := Repositories.NewRepository(client)
	userService := Services.NewUserService(authRepository, userRepository)

	err = userService.NewUserScore(ctx, score)
	if err != nil {
		t.Errorf("Error adding new score to Firestore: %v", err)
	}

	iter := client.Collection("Score").Where("UserID", "==", score.UserID).Where("Language", "==", score.Language).Where("Score", "==", score.Score).Documents(ctx)
	var docRef *firestore.DocumentRef
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			break
		}

		docRef = doc.Ref
	}
	doc, err := docRef.Get(ctx)
	if err != nil {
		t.Errorf("Error getting new score from Firestore: %v", err)
	}
	var savedScore Models.Score
	err = doc.DataTo(&savedScore)
	if err != nil {
		t.Errorf("Error parsing new score from Firestore: %v", err)
	}
	if savedScore.UserID != score.UserID {
		t.Errorf("Saved score has wrong UserID: expected %v, got %v", score.UserID, savedScore.UserID)
	}
	if savedScore.Score != score.Score {
		t.Errorf("Saved score has wrong Score: expected %v, got %v", score.Score, savedScore.Score)
	}

	_, err = docRef.Delete(ctx)
	if err != nil {
		t.Errorf("Error deleting new score from Firestore: %v", err)
	}
}

func TestHandler_GetUserHistory(t *testing.T) {
	client, err := initializeFirestoreClient()
	assert.NoError(t, err)
	defer client.Close()

	router := mux.NewRouter()

	userRepository := Repositories.NewRepository(client)
	authRepository := Repositories.NewRepository(client)
	wordRepository := Repositories.NewRepository(client)
	userService := Services.NewUserService(authRepository, userRepository)
	wordService := Services.NewWordService(wordRepository)
	handler := NewHttpHandler(userService, wordService)

	router.HandleFunc("/user/score/{id}", handler.GetUserHistory)

	req, err := http.NewRequest("GET", "/user/score/recWU1AGad27OhZ9FVIo", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	expected := `null`
	assert.Equal(t, expected, rr.Body.String())
}

func TestHandler_AddNewWords(t *testing.T) {
	client, err := initializeFirestoreClient()
	assert.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	words := []Models.Word{}
	words = append(words, Models.Word{
		Language: 3,
		Word:     "test-word",
	})

	wordRepository := Repositories.NewRepository(client)
	wordService := Services.NewWordService(wordRepository)

	_, err = wordService.AddNewWords(ctx, words)
	if err != nil {
		t.Errorf("Error adding new words to Firestore: %v", err)
	}

	iter := client.Collection("Words").Where("Language", "==", words[0].Language).Where("Word", "==", words[0].Word).Documents(ctx)
	var docRef *firestore.DocumentRef
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			break
		}

		docRef = doc.Ref
	}
	doc, err := docRef.Get(ctx)
	if err != nil {
		t.Errorf("Error getting new words from Firestore: %v", err)
	}
	var savedWord Models.Word
	err = doc.DataTo(&savedWord)
	if err != nil {
		t.Errorf("Error parsing new words from Firestore: %v", err)
	}
	if savedWord.Language != words[0].Language {
		t.Errorf("Saved word has wrong Language: expected %v, got %v", words[0].Language, savedWord.Language)
	}
	if savedWord.Word != words[0].Word {
		t.Errorf("Saved word has wrong Word: expected %v, got %v", words[0].Word, savedWord.Word)
	}

	_, err = docRef.Delete(ctx)
	if err != nil {
		t.Errorf("Error deleting new word from Firestore: %v", err)
	}
}
