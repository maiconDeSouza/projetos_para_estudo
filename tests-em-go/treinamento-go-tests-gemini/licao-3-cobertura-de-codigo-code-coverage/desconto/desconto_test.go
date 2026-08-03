package desconto

import "testing"

func TestCalcularDesconto(t *testing.T) {
	teste := []struct {
		nome        string
		valor       float64
		porcentagem float64
	}{
		{"Valor inválidos", -1, 0.00},
		{"Sem desconto 99.99", 99.99, 0.00},
		{"Sem desconto - baixo de 100", 30.00, 0.00},
		{"10 porcento de desconto - exatamente 100", 100.00, 0.10},
		{"10 porcento de desconto - quase 500", 499.99, 0.10},
		{"20 porcento de desconto - exatamente 500", 500.00, 0.20},
		{"20 porcento de desconto - mais de 500", 987.00, 0.20},
	}

	for _, tt := range teste {
		t.Run(tt.nome, func(t *testing.T) {
			esperado := tt.valor * tt.porcentagem
			resultado := CalcularDesconto(tt.valor)

			if esperado != resultado {
				t.Errorf("TestCalcularDesconto(%.2f) = %.2f; esperado %.2f", tt.valor, resultado, esperado)
			}
		})
	}
}
