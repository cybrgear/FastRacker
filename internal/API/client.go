package API

import (
	"encoding/json"
	"net/http"
	"time"

	"FastRacker/internal/API/models"
)

var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

// GetArtists récupère la liste des artistes depuis l'API externe en toute sécurité
func GetArtists() ([]models.Artist, error) {
	resp, err := httpClient.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var artists []models.Artist
	err = json.NewDecoder(resp.Body).Decode(&artists)
	if err != nil {
		return nil, err
	}

	return artists, nil
}