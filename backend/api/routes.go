package api

import (
	"FinRiskScore/backend/api/handlers" // Importe les handlers depuis le dossier "handlers"
	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", handlers.HomeHandler)
	r.HandleFunc("/calculate-risk-score", handlers.CalculateRiskScoreHandler).Methods("GET")

	return r
}
