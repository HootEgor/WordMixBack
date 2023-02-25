package Handlers

import (
	Models "WordMixBack/src/Model"
	"encoding/json"
)

func ParseUserToJSON(user Models.User) (string, error) {
	jsonData, err := json.Marshal(user)
	if err != nil {
		return "", err
	}

	jsonString := string(jsonData)

	return jsonString, nil
}
