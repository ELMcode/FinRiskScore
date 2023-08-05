package handlers

import (
	"FinRiskScore/backend/algorithm"
	"FinRiskScore/backend/data"
	"fmt"
	"net/http"
)

func CalculateRiskScoreHandler(w http.ResponseWriter, r *http.Request) {
	apiKey := "051a81835513220504d1199f85822b5b"

	// Recup le symbole de l'entreprise à partir des variables de requête
	symbol := r.URL.Query().Get("symbol")

	if symbol == "" {
		http.Error(w, "Veuillez fournir le symbole de l'entreprise via le paramètre 'symbol'\nExemple : http://localhost:8080/calculate-risk-score?symbol=?", http.StatusBadRequest)
		return
	}

	// Récupère les données financières du symbole d'entreprise spécifié
	financialData, err := data.FetchFinancialData(apiKey, symbol)
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des données", http.StatusInternalServerError)
		return
	}

	// Vérifie si données ont été récupérées
	if len(financialData) == 0 {
		http.Error(w, "Aucune donnée financière récupérée pour le symbole d'entreprise spécifié", http.StatusNotFound)
		return
	}

	// Calcul le risk score en utilisant l'algo
	riskScore := algorithm.CalculateRiskScore(financialData)

	// Renvoi le risk score dans la réponse HTTP
	fmt.Fprintf(w, "Risk Score for %s: %.2f", symbol, riskScore)
}
