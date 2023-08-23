// Package handlers calculate_risk_score_handler.go
package handlers

import (
	"encoding/json"
	"math"
	"net/http"

	"FinRiskScore/backend/algorithm"
	"FinRiskScore/backend/api/validators"
	"FinRiskScore/backend/config"
	"FinRiskScore/backend/data"
)

// Response Structure de réponse pour endpoint CalculateRiskScoreHandler
type Response struct {
	Symbol    string  `json:"symbol"`
	RiskScore float64 `json:"riskScore"`

	CompanyName       string  `json:"companyName"`
	Industry          string  `json:"industry"`
	FullTimeEmployees string  `json:"fullTimeEmployees"`
	Description       string  `json:"description"`
	Price             float64 `json:"price"`
	Website           string  `json:"website"`
	Image             string  `json:"image"`
	Exchange          string  `json:"exchange"`
	Changes           float64 `json:"changes"`
	IsActivelyTrading bool    `json:"isActivelyTrading"`
}

// CalculateRiskScoreHandler gère la demande de calcul du score de risque financier
func CalculateRiskScoreHandler(w http.ResponseWriter, r *http.Request) {

	// Récupére la clé API à partir du dossier/fichier config
	apiKey := config.APIKey

	// Extraction du symbole de l'URL
	symbol := r.URL.Query().Get("symbol")

	// Valide les paramètres de calcul du score de risque financier
	// Si param valides, la fonction de validation validators.ValidateCalculateRiskScoreParams() renvoie nil.
	params := validators.CalculateRiskScoreParams{Symbol: symbol}
	if err := validators.ValidateCalculateRiskScoreParams(params); err != nil {
		http.Error(w, "Invalid params", http.StatusBadRequest)
		return
	}

	// Récup les données financières à partir de l'API externe
	financialData, err := data.FetchFinancialData(apiKey, symbol)
	if err != nil {
		http.Error(w, "Error fetching data", http.StatusInternalServerError)
		return
	}

	// Vérifie les données financières récupérées
	if len(financialData) == 0 {
		http.Error(w, "No data found", http.StatusNotFound)
		return
	}

	// Utilise le calcul du score de risque financier du dossier algorithm
	riskScore := algorithm.CalculateRiskScore(financialData)
	// limite à 2 décimales le résultat du risque financier
	formattedRiskScore := math.Round(riskScore*100) / 100

	// Récupére les données de description de l'entreprise de l'api externe de description
	description, err := data.FetchCompanyDescription(apiKey, symbol)
	if err != nil {
		http.Error(w, "Error fetching description", http.StatusInternalServerError)
		return
	}

	// Création de la réponse JSON
	resp := Response{
		Symbol:    symbol,
		RiskScore: formattedRiskScore,

		CompanyName:       description[0].CompanyName,
		Industry:          description[0].Industry,
		FullTimeEmployees: description[0].FullTimeEmployees,
		Description:       description[0].Description,
		Price:             description[0].Price,
		Website:           description[0].Website,
		Image:             description[0].Image,
		Exchange:          description[0].Exchange,
		Changes:           description[0].Changes,
		IsActivelyTrading: description[0].IsActivelyTrading,
	}

	// Encodage de la réponse en JSON
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}

	// Définition de l'en-tête header de la réponse JSON
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}
