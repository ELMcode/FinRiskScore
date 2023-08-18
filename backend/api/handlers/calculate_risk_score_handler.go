// calculate_risk_score_handler.go
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
	formattedRiskScore := float64(int(riskScore*100)) / 100

	description, err := data.FetchCompanyDescription(apiKey, symbol)
	if err != nil {
		http.Error(w, "Error fetching description", http.StatusInternalServerError)
		return
	}

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

	jsonResp, err := json.Marshal(resp)

	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)

}
