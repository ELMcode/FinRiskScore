package api

import (
	"FinRiskScore/backend/algorithm"
	"FinRiskScore/backend/data"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Page d'accueil"))
}

func CalculateRiskScoreHandler(w http.ResponseWriter, r *http.Request) {
	apiKey := "XXX"
	financialData, err := data.FetchFinancialData(apiKey)
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des données", http.StatusInternalServerError)
		return
	}

	// Vérifie si données ont été récupérées
	if len(financialData) == 0 {
		http.Error(w, "Aucune donnée financière récupérée", http.StatusNotFound)
		return
	}

	// Calcul le risk score en utilisant l'algo
	riskScore := algorithm.CalculateRiskScore(financialData)

	// Renvoi le risk score dans la réponse HTTP
	fmt.Fprintf(w, "Risk Score: %.2f", riskScore)
}

func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/calculate-risk-score", CalculateRiskScoreHandler).Methods("GET")

	return r
}
