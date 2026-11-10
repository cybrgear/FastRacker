package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"

	API "FastRacker/internal/API"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	artists, err := API.GetArtists()
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Chemins des templates (Layout + Index)
	tmplFiles := []string{
		filepath.Join("web", "templates", "layout.html"),
		filepath.Join("web", "templates", "index.html"),
	}

	tmpl, err := template.ParseFiles(tmplFiles...)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "layout", artists)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
	}
}
/*HomeHandler gère la route racine ("/"). 
 Il vérifie si le chemin est exact, 
 récupère la liste des artistes via l'API, 
charge le design global (layout.html) combiné avec la page d'accueil (index.html), 
 puis injecte les données (artists)*/
