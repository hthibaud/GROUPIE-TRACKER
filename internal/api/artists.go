package artists

import (
	"encoding/json"
	"net/http"
)

func getArtists([]Artist, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var Artists []Artists
	if err := json.NewDecoder(resp.Body).Decode(&artists); err != nil {
		return nil, err
	}
	return artists, nil
}
