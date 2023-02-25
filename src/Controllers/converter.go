package Handlers

import (
	"encoding/json"
)

func ParseToJSON(data interface{}) (string, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	jsonString := string(jsonData)
	return jsonString, nil
}
