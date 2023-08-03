package algorithm

import (
	"FinRiskScore/backend/data"
	"fmt"
)

func CalculateRiskScore(financialData []data.FinancialData) float64 {

	// Sélectionne la dernière entrée des données financières (la plus récente)
	if len(financialData) == 0 {
		fmt.Println("No financial data found.")
		return 0.0
	}

	fdata := financialData[0]

	// ratio de rentabilité (Net Profit Margin)
	netProfitMargin := fdata.NetIncome / fdata.Revenue

	// ratio de marge brute (Gross Profit Margin)
	grossProfitMargin := fdata.GrossProfit / fdata.Revenue

	// ratio de dépenses opérationnelles (Operating Expense Ratio)
	operatingExpenseRatio := fdata.OperatingExpenses / fdata.Revenue

	// score de risque financier
	riskScore := (netProfitMargin + grossProfitMargin + operatingExpenseRatio) / 3

	return riskScore
}
