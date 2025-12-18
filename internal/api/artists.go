package artists

import (
	"encoding/json"
	"net/http"
)

type Artist struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Image        string   `json:"image"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

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
