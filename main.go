package main

import (
	"log"
	"net/http"

	"FastRacker/internal/handlers"
)

import _ "FastRacker/internal/API"

func main() {
	// Configuration du routeur et des fichiers statiques (CSS, JS, Images)
	mux := http.NewServeMux()

	// Servir les fichiers CSS/JS/Images
	fileServer := http.FileServer(http.Dir("./web/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	// Routes principales
	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/artist", handlers.ArtistHandler)

	// Démarrage du serveur avec gestion robuste
	port := ":8080"
	log.Printf("Serveur démarré sur http://localhost%s\n", port)
	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatalf("Erreur critique du serveur : %v", err)
	}
}
