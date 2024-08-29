package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const yandexSpellerURL = "https://speller.yandex.net/services/spellservice.json/checkTexts"

// an error returned by Yandex Speller
type SpellerError struct {
	Code        int      `json:"code"`
	Pos         int      `json:"pos"`
	Len         int      `json:"len"`
	Word        string   `json:"word"`
	Suggestions []string `json:"s"`
}

// checks multiple texts for spelling errors
func CheckSpelling(texts []string) ([][]SpellerError, error) {
	// Prepare the request payload
	query := url.Values{}
	for _, text := range texts {
		query.Add("text", text)  // Add each text fragment individually
	}

	// Send a request to Yandex Speller API
	response, err := http.Get(fmt.Sprintf("%s?%s", yandexSpellerURL, query.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Yandex.Speller: %v", err)
	}
	defer response.Body.Close()

	// Decode the response
	var spellingErrors [][]SpellerError
	if err := json.NewDecoder(response.Body).Decode(&spellingErrors); err != nil {
		return nil, fmt.Errorf("failed to decode Yandex.Speller response: %v", err)
	}

	return spellingErrors, nil
}
