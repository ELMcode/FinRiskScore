package main

import (
	"FinRiskScore/backend/api"
	"net/http"
)

func main() {
	// routes définies dans routes.go
	r := api.RegisterRoutes()

	// gestionnaire racine pour les routes
	http.Handle("/", r)

	// Démarrage serveur HTTP port 8080
	http.ListenAndServe(":8080", nil)
}
