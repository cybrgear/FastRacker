package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"

	"FastRacker/internal/API"
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