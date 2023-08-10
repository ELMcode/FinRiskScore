package handlers

import (
	"encoding/json"
	"net/http"

	"FinRiskScore/backend/algorithm"
	"FinRiskScore/backend/api/validators"
	"FinRiskScore/backend/config"
	"FinRiskScore/backend/data"
)

type Response struct {
	Symbol    string  `json:"symbol"`
	RiskScore float64 `json:"riskScore"`
}

func CalculateRiskScoreHandler(w http.ResponseWriter, r *http.Request) {

	apiKey := config.APIKey

	symbol := r.URL.Query().Get("symbol")

	params := validators.CalculateRiskScoreParams{Symbol: symbol}

	if err := validators.ValidateCalculateRiskScoreParams(params); err != nil {
		http.Error(w, "Invalid params", http.StatusBadRequest)
		return
	}

	financialData, err := data.FetchFinancialData(apiKey, symbol)
	if err != nil {
		http.Error(w, "Error fetching data", http.StatusInternalServerError)
		return
	}

	if len(financialData) == 0 {
		http.Error(w, "No data found", http.StatusNotFound)
		return
	}

	riskScore := algorithm.CalculateRiskScore(financialData)

	resp := Response{
		Symbol:    symbol,
		RiskScore: riskScore,
	}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)

}
