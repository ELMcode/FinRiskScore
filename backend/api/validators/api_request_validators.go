package validators

import (
	"github.com/go-playground/validator/v10"
)

// Defini un custom validators instance
var validate = validator.New()

// CalculateRiskScoreParams Structure pour la validation des paramètres de la requête "calculate-risk-score"
type CalculateRiskScoreParams struct {
	Symbol string `json:"symbol" validate:"required"`
}

// ValidateCalculateRiskScoreParams Fonction pour valider les paramètres de la requête "calculate-risk-score"
func ValidateCalculateRiskScoreParams(params CalculateRiskScoreParams) error {
	return validate.Struct(params)
}
