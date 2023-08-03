package main

import (
	"FinRiskScore/backend/algorithm"
	"FinRiskScore/backend/data"
	"fmt"
)

func main() {

	apiKey := "XXX"

	financialData, err := data.FetchFinancialData(apiKey)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}

	// Vérifier que des données ont bien été récupérées
	if len(financialData) == 0 {
		fmt.Println("No financial data fetched")
		return
	}

	// Récupérer les dernières données
	fdata := financialData[0]

	// Afficher les données utilisées pour le risk score
	fmt.Println("Data used for risk score:")
	fmt.Printf("Company: %s\n", fdata.Symbol)
	fmt.Printf("Date: %s\n", fdata.Date)
	fmt.Printf("Revenue: %.2f\n", fdata.Revenue)
	fmt.Printf("Net Income: %.2f\n", fdata.NetIncome)
	fmt.Printf("Gross Profit: %.2f\n", fdata.GrossProfit)
	fmt.Printf("Operating Expenses: %.2f\n", fdata.OperatingExpenses)

	// Calculer le risk score
	riskScore := algorithm.CalculateRiskScore(financialData)

	// Afficher le risk score
	fmt.Printf("Financial risk score: %.2f\n", riskScore)

}
