package API

// Récupération des données d'artistes depuis l'API distante

import (
	"net/http"
	"fmt"
	"encoding/json"
	"FastRacker/internal/API/models"
)

// GetLocations récupère la liste des lieux de concert
func GetLocations() ([]models.Location, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var locIndex models.LocationIndex
	err = json.NewDecoder(resp.Body).Decode(&locIndex)
	if err != nil {
		return nil, err
	}

	return locIndex.Index, nil
}

// GetDates récupère la liste des dates de concert
func GetDates() ([]models.Date, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var dateIndex models.DateIndex
	err = json.NewDecoder(resp.Body).Decode(&dateIndex)
	if err != nil {
		return nil, err
	}

	return dateIndex.Index, nil
}

// GetRelations récupère les relations (lieux et dates associés par artiste)
func GetRelations() ([]models.Relation, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var relIndex models.RelationIndex
	err = json.NewDecoder(resp.Body).Decode(&relIndex)
	if err != nil {
		return nil, err
	}

	return relIndex.Index, nil
}

/* ici le code go utilise mon client HTTP existant
pour faire une requête GET sur chaque endpoint de l'API Groupie Tracker.
Il décode le JSON reçu pour le ranger bien proprement dans les structures
(models) qui viennent d'être définies.
Il retourne directement la liste ([]models.Location, etc.)
prête à être utilisée par les handlers.*/

// GetRelationByArtistID récupère les relations (lieux et dates) d'un artiste via son ID
func GetRelationByArtistID(id int) (models.Relation, error) {
	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/relation/%d", id)
	resp, err := http.Get(url)
	if err != nil {
		return models.Relation{}, err
	}
	defer resp.Body.Close()

	var rel models.Relation
	err = json.NewDecoder(resp.Body).Decode(&rel)
	if err != nil {
		return models.Relation{}, err
	}

	return rel, nil
}
