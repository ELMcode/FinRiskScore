package main

import (
	"FinRiskScore/backend/api"
	"fmt"
	"net/http"
)

func main() {
	// routes définies dans routes.go
	r := api.RegisterRoutes()

	// gestionnaire racine pour les routes
	http.Handle("/", r)

	// Démarrage serveur HTTP port 8080
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Erreur lors du démarrage du serveur:", err)
	}
}
