package converter

import (
	"encoding/json"
	"errors"
	"fmt"
)

func FormatJSON(input string) (string, error) {
	var jsonData interface{}
	err := json.Unmarshal([]byte(input), &jsonData)
	if err != nil {
		return "", errors.New("неверный JSON")
	}

	formattedJSON, err := json.MarshalIndent(jsonData, "", "  ")
	if err != nil {
		return "", fmt.Errorf("ошибка форматирования: %v", err)
	}

	return string(formattedJSON), nil
}
