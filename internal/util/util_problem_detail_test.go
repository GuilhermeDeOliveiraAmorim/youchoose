package util

import (
	"encoding/json"
	"testing"
)

func TestProblemDetailsSerialization(t *testing.T) {
	pd := ProblemDetails{
		Type:     "validation_error",
		Title:    "Erro de Validação",
		Status:   400,
		Detail:   "Detalhe do erro de validação",
		Instance: "/users/123",
	}

	jsonData, err := json.Marshal(pd)
	if err != nil {
		t.Errorf("Erro ao converter para JSON: %v", err)
	}

	var newPD ProblemDetails
	err = json.Unmarshal(jsonData, &newPD)
	if err != nil {
		t.Errorf("Erro ao converter de JSON: %v", err)
	}

	if pd != newPD {
		t.Errorf("As instâncias de ProblemDetails não são iguais. Original: %+v, Convertida: %+v", pd, newPD)
	}
}
