package handlers

import (
	"fmt"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	fmt.Fprintf(w, "Erreur %d : Une erreur est survenue sur le serveur.", status)
}

/*Une fonction centralisée (ErrorHandler)
 permet de renvoyer un code HTTP propre
  (404 Not Found, 400 Bad Request, 500 Internal Server Error)
 accompagné d'un message d'erreur.*/