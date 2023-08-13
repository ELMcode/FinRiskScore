package data

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// CompanyDescription Struct
type CompanyDescription struct {
	Symbol            string  `json:"symbol"`
	Price             float64 `json:"price"`
	Beta              float64 `json:"beta"`
	VolAvg            int     `json:"volAvg"`
	MktCap            int64   `json:"mktCap"`
	LastDiv           float64 `json:"lastDiv"`
	Range             string  `json:"range"`
	Changes           float64 `json:"changes"`
	CompanyName       string  `json:"companyName"`
	Currency          string  `json:"currency"`
	Cik               string  `json:"cik"`
	Isin              string  `json:"isin"`
	Cusip             string  `json:"cusip"`
	Exchange          string  `json:"exchange"`
	ExchangeShortName string  `json:"exchangeShortName"`
	Industry          string  `json:"industry"`
	Website           string  `json:"website"`
	Description       string  `json:"description"`
	Ceo               string  `json:"ceo"`
	Sector            string  `json:"sector"`
	Country           string  `json:"country"`
	FullTimeEmployees string  `json:"fullTimeEmployees"`
	Phone             string  `json:"phone"`
	Address           string  `json:"address"`
	City              string  `json:"city"`
	State             string  `json:"state"`
	Zip               string  `json:"zip"`
	DcfDiff           float64 `json:"dcfDiff"`
	Dcf               float64 `json:"dcf"`
	Image             string  `json:"image"`
	IpoDate           string  `json:"ipoDate"`
	DefaultImage      bool    `json:"defaultImage"`
	IsEtf             bool    `json:"isEtf"`
	IsActivelyTrading bool    `json:"isActivelyTrading"`
	IsAdr             bool    `json:"isAdr"`
	IsFund            bool    `json:"isFund"`
}

// FetchCompanyDescriptions Fonction pour récupérer les descriptions de l'API externe
func FetchCompanyDescriptions(apiKey, symbol string) ([]CompanyDescription, error) {

	// Construction de l'URL avec le symbole
	apiURL := fmt.Sprintf("https://financialmodelingprep.com/api/v3/profile/%s?apikey=%s", symbol, apiKey)

	// Requête HTTP GET
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("erreur API: %v", err)
	}

	defer resp.Body.Close()

	// Décoder la réponse JSON
	var descriptionData []CompanyDescription // Crée une variable pour stocker les données de description
	if err := json.NewDecoder(resp.Body).Decode(&descriptionData); err != nil {
		return nil, fmt.Errorf("erreur décodage: %v", err)
	}

	return descriptionData, nil
}
