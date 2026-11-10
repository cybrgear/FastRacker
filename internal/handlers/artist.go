package handlers

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	API "FastRacker/internal/API"
	"FastRacker/internal/API/models"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Récupérer l'ID depuis l'URL (?id=4)
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	// 2. Récupérer la liste des artistes pour trouver le bon
	artists, err := API.GetArtists()
	if err != nil {
		log.Println("Erreur GetArtists:", err) // <--- Ajoute ceci
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	var foundArtist models.Artist
	found := false
	for _, a := range artists {
		if a.ID == id {
			foundArtist = a
			found = true
			break
		}
	}

	if !found {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	// 3. NOUVEAU : Récupérer les relations (dates/lieux) de cet artiste spécifique
	relation, err := API.GetRelationByArtistID(id)
	if err != nil {
		log.Println("Erreur GetRelationByArtistID:", err) // <--- Ajoute ceci
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// 4. Regrouper les données dans notre structure de vue
	pageData := models.ArtistPageData{
		Artist:   foundArtist,
		Relation: relation,
	}

	// 5. Charger et exécuter le template
	tmplFiles := []string{
		filepath.Join("web", "templates", "layout.html"),
		filepath.Join("web", "templates", "artist.html"),
	}

	tmpl, err := template.ParseFiles(tmplFiles...)
	if err != nil {
		log.Println("Erreur chargement template:", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "layout", pageData)
	if err != nil {
		log.Println("Erreur Template Execution:", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
	}
}

/*L'URL récupère un paramètre d'identification (ex: /artist?id=4).
Le code extrait cet identifiant, cherche l'artiste correspondant dans la liste,
et si tout est ok, charge le template artist.html en lui envoyant
les informations spécifiques de l'artiste.*/

/* à ce niveau, le backend intercepte l'ID demandé dans l'URL,
va chercher les informations de l'artiste, va chercher de manière
ciblée ses dates et lieux de concert via l'API externe,
et empaquète le tout dans un objet pageData prêt à être affiché.*/
