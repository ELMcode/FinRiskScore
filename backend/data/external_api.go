package data

import (
	"encoding/json" // Package pour l'encodage et le décodage des données JSON
	"fmt"           // Package pour afficher des messages
	"net/http"      // Package pour effectuer des requêtes HTTP
)

// FinancialData représente les données financières reçues depuis l'API.
type FinancialData struct {
	Date                                    string  `json:"date"`
	Symbol                                  string  `json:"symbol"`
	ReportedCurrency                        string  `json:"reportedCurrency"`
	Cik                                     string  `json:"cik"`
	FillingDate                             string  `json:"fillingDate"`
	AcceptedDate                            string  `json:"acceptedDate"`
	CalendarYear                            string  `json:"calendarYear"`
	Period                                  string  `json:"period"`
	Revenue                                 float64 `json:"revenue"`
	CostOfRevenue                           float64 `json:"costOfRevenue"`
	GrossProfit                             float64 `json:"grossProfit"`
	GrossProfitRatio                        float64 `json:"grossProfitRatio"`
	ResearchAndDevelopmentExpenses          float64 `json:"researchAndDevelopmentExpenses"`
	SellingGeneralAndAdministrativeExpenses float64 `json:"sellingGeneralAndAdministrativeExpenses"`
	OperatingExpenses                       float64 `json:"operatingExpenses"`
	OperatingIncome                         float64 `json:"operatingIncome"`
	OperatingIncomeRatio                    float64 `json:"operatingIncomeRatio"`
	NetIncome                               float64 `json:"netIncome"`
	Eps                                     float64 `json:"eps"`
	Epsdiluted                              float64 `json:"epsdiluted"`
	WeightedAverageShsOut                   float64 `json:"weightedAverageShsOut"`
	WeightedAverageShsOutDil                float64 `json:"weightedAverageShsOutDil"`
	Link                                    string  `json:"link"`
	FinalLink                               string  `json:"finalLink"`
}

// FetchFinancialData récupère les données financières depuis une API externe
func FetchFinancialData(apiKey string) ([]FinancialData, error) {

	apiURL := fmt.Sprintf("https://financialmodelingprep.com/api/v3/income-statement/AAPL?limit=120&apikey=%s", apiKey)

	// requête GET à l'API pour récupérer les données
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la récupération des données : %v", err) // Retourne une erreur en cas d'échec de la requête
	}
	defer resp.Body.Close() // Ferme le corps de la réponse HTTP après l'exécution de la fonction

	var financialData []FinancialData // Crée une variable pour stocker les données financières

	// Décode les données JSON à partir du corps de la réponse HTTP et les stocke dans la variable financialData
	err = json.NewDecoder(resp.Body).Decode(&financialData) // si Decode retourne données valides alors ca passe dans financialData, si retourne une erreur alors ca ne passe pas dans financialData et reste dans la variable err
	if err != nil {
		return nil, fmt.Errorf("erreur lors du décodage JSON : %v", err) // Retourne une erreur en cas d'échec du décodage JSON
	}

	return financialData, nil // Retourne les données financières décodées
}
