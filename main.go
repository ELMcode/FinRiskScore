package main

import (
	"FinRiskScore/backend/api"
	"log"
	"net/http"
)

func main() {
	// routes définies dans routes.go
	r := api.RegisterRoutes()

	// gestionnaire racine pour les routes
	http.Handle("/", r)

	// Démarrage serveur HTTP port 8080 avec 'log' pour message dans console
	addr := ":8080"
	log.Printf("Serveur en ligne : http://localhost%s", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Erreur lors du démarrage du serveur: %v", err)
	}
}
