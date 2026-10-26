package handlers

import (
	"fmt"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	fmt.Fprintf(w, "Erreur %d : Une erreur est survenue sur le serveur.", status)
}