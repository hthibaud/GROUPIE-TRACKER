package api

import (
	"encoding/json"
	"net/http"
)

type Relations struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

func GetRelations(url string) (Relations, error) {
	resp, err := http.Get(url)
	if err != nil {
		return Relations{}, err
	}
	defer resp.Body.Close()

	var relations Relations
	err = json.NewDecoder(resp.Body).Decode(&relations)
	return relations, err
}
