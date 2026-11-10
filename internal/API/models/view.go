package models

// stucture combinee
type ArtistDetail struct {
	Artist    Artist
	Relation  Relation // Contient les dates et lieux associés à cet artiste
}

// ArtistPageData regroupe l'artiste et ses concerts pour l'affichage
type ArtistPageData struct {
	Artist   Artist
	Relation Relation
}
