package handlers

import (
	"FinRiskScore/backend/algorithm"
	"FinRiskScore/backend/api/validators"
	"FinRiskScore/backend/config"
	"FinRiskScore/backend/data"
	"fmt"
	"net/http"
)

func CalculateRiskScoreHandler(w http.ResponseWriter, r *http.Request) {
	apiKey := config.APIKey

	// Obtient le symbole de l'entreprise à partir des variables de requête
	symbol := r.URL.Query().Get("symbol")

	// Crée une instance de CalculateRiskScoreParams avec le symbole
	params := validators.CalculateRiskScoreParams{Symbol: symbol}

	// Valide les paramètres de la requête
	if err := validators.ValidateCalculateRiskScoreParams(params); err != nil {
		http.Error(w, "Paramètres de requête invalides : paramètre 'symbol' manquant.\nExemple : http://localhost:8080/calculate-risk-score?symbol=GOOGL", http.StatusBadRequest)
		return
	}

	// Récupère les données financières pour le symbole d'entreprise spécifié
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
