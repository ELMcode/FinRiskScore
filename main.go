package main

import (
	"FinRiskScore/backend/api"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	// routes définies dans routes.go
	r := api.RegisterRoutes()

	// Créer un routeur pour gérer les routes
	router := http.NewServeMux()

	// gestionnaires de routes
	router.Handle("/", r)

	// middleware CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	})

	// Demarrage serveur HTTP port 8080 avec le middleware CORS
	addr := ":8080"
	log.Printf("Serveur en ligne : http://localhost%s", addr)
	err := http.ListenAndServe(addr, c.Handler(router))
	if err != nil {
		log.Fatalf("Erreur lors du démarrage du serveur: %v", err)
	}
}
