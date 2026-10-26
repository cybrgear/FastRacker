package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"

	"FastRacker/internal/API"
	"FastRacker/internal/API/models"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	artists, err := API.GetArtists()
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	var foundArtist *models.Artist
	for i := range artists {
		if artists[i].ID == id {
			foundArtist = &artists[i]
			break
		}
	}

	if foundArtist == nil {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	tmplFiles := []string{
		filepath.Join("web", "templates", "layout.html"),
		filepath.Join("web", "templates", "artist.html"),
	}

	tmpl, err := template.ParseFiles(tmplFiles...)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "layout", foundArtist)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
	}
}