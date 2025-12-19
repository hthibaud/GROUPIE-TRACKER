package api

import (
	"encoding/json"
	"net/http"
)

type Artist struct {
	ID           int         `json:"id"`
	Name         string      `json:"name"`
	Image        string      `json:"image"`
	Members      []string    `json:"members"`
	CreationDate int         `json:"creationDate"`
	FirstAlbum   string      `json:"firstAlbum"`
	Locations    interface{} `json:"locations"`
	ConcertDates interface{} `json:"concertDates"`
	Relations    interface{} `json:"relations"`
}

func GetArtists() ([]Artist, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var artists []Artist
	if err := json.NewDecoder(resp.Body).Decode(&artists); err != nil {
		return nil, err
	}
	return artists, nil
}

func GetArtist() ([]Artist, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var artist []Artist
	if err := json.NewDecoder(resp.Body).Decode(&artist); err != nil {
		return nil, err
	}
	return artist, nil
}
