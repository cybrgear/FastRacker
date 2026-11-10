package models

// Structure pour les Dates
type DateIndex struct {
	Index []Date `json:"index"`
}

type Date struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}
