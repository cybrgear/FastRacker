package models

// Structure pour les Relations (Lieux + Dates par artiste)
type RelationIndex struct {
	Index []Relation `json:"index"`
}

type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"` // Clé = Lieu, Valeur = Liste des dates
}
