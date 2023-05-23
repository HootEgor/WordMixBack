package Handlers

import (
	Models "WordMixBack/src/Model"
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

type HttpHandler struct {
	authService AuthService
	userService UserService
	wordService WordService
}

type AuthService interface {
	AddNewUser(ctx context.Context, user Models.User) (string, error)
	LoginUser(ctx context.Context, user Models.User) (string, error)
}

func NewHttpHandler(userService UserService, wordService WordService, authService AuthService) *HttpHandler {
	return &HttpHandler{
		authService: authService,
		userService: userService,
		wordService: wordService,
	}
}

func (h *HttpHandler) Register(w http.ResponseWriter, r *http.Request) {
	var newUser Models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		return
	}

	userID, err := h.authService.AddNewUser(r.Context(), newUser)
	if err != nil {
		return
	}

	token, err := GenerateJWT(userID)
	if err != nil {
		return
	}

	fmt.Fprintf(w, "%+v", token)
}

func (h *HttpHandler) Login(w http.ResponseWriter, r *http.Request) {
	var newUser Models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		return
	}

	userID, err := h.authService.LoginUser(r.Context(), newUser)
	if err != nil || userID == "" {
		http.Error(w, "user not found", 404)
		return
	}

	token, err := GenerateJWT(userID)
	if err != nil {
		return
	}

	fmt.Fprintf(w, "%+v", token)
}

func GenerateJWT(userID string) (string, error) {

	claims := jwt.MapClaims{}
	claims["UserID"] = userID
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKey := []byte("my-secret-key")
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	jsonData, err := json.Marshal(tokenString)
	if err != nil {
		return "", err
	}

	jsonString := string(jsonData)

	return jsonString, nil
}
