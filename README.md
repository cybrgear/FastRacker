Groupie Tracker

°Présentation du Projet

Groupie Tracker est une application web développée en Go qui consomme une API RESTful externe pour afficher, filtrer et géolocaliser des informations relatives à des groupes de musique et artistes (membres, dates et lieux de concerts, relations entre ces données).

Le projet met l'accent sur la robustesse du serveur (gestion stricte des erreurs sans crash), la modularité de l'architecture et une ergonomie soignée respectant les bonnes pratiques IHM.

°Stack Technique

Backend : Go (Standard Library uniquement — net/http, html/template, encoding/json, etc.)

Frontend : HTML5, CSS3, JavaScript (Vanilla JS)

Cartographie : API de géocodage & Leaflet / OpenStreetMap (Bonus B-4)

Architecture : Clean Architecture / MVC adaptée à Go

°Architecture du Projet

L'application est structurée de manière modulaire au sein du répertoire internal/ afin de séparer la logique métier, la gestion HTTP et la couche de présentation.

FastRacker/
├── docs/                     # Documentation du projet
│   └── architecture.md
├── internal/                 # Logiciels et packages internes non-exportables
│   ├── API/                  # Entités et communication avec l'API externe
│   │   ├── models/           # Structures Go (Data Models & ViewModels)
│   │   ├── fetch.go          # Client HTTP pour récupérer le JSON
│   │   └── client.go
│   ├── filter/               # Algorithmes de filtrage (dates, membres, lieux)
│   ├── geocode/              # Service de conversion d'adresses en coordonnées (Lat/Lng)
│   ├── handlers/             # Contrôleurs HTTP (Home, Artist, Search, Filter, Geo, Errors)
│   ├── middleware/           # Intercepteurs HTTP (Logging, Recover anti-crash)
│   ├── search/               # Algorithme de recherche avec suggestions en temps réel
│   └── store/                # Couche d'accès et de mise en cache des données
├── web/                      # Ressources Frontend
│   ├── static/               # Assets statiques (CSS, JS, Images)
│   └── templates/            # Templates HTML (Layout, Vues et Composants Partials)
├── go.mod                    # Fichier de module Go
├── main.go                   # Point d'entrée de l'application
└── README.md


°Fonctionnalités Principales

Core (FT-1 à FT-3)

Récupération & Manipulation JSON : Consommation des 4 endpoints de l'API (artists, locations, dates, relation).

Affichage dynamique : Présentation claire sous forme de cartes d'artistes et pages détaillées.

Communication Client-Serveur : Gestion des requêtes/réponses dynamiques via Go.

Gestion d'erreurs robuste : Middleware recover.go et handlers dédiés pour empêcher tout crash du serveur.

Extensions & Bonus (B-1 à B-4)

UI/UX & Design (B-1) : Interface conforme aux 8 règles d'or de Shneiderman.

Barre de Recherche Avancée (B-2) : Recherche insensible à la casse avec auto-complétion et typage des suggestions (ex: "Freddie Mercury - membre").

Système de Filtres (B-3) : Filtrage combiné par intervalle de dates, nombre de membres et zones géographiques.

Géolocalisation & Carte Interactive (B-4) : Géocodage des lieux de concert et placement de marqueurs sur une carte interactive.

°Lancement du Projet

Prérequis

Go 1.20 ou supérieur installe.

Installation & Exécution

Cloner le dépôt :

git clone <url-du-repo>
cd FastRacker


Lancer le serveur :

go run main.go


Accéder à l'application dans votre navigateur :

http://localhost:8080


°Lancement des Tests

Le projet inclut une suite de tests unitaires pour valider les fonctions métier et la manipulation des données :

go test ./internal/...
