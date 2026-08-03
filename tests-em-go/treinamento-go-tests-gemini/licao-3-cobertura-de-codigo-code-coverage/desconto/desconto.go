package desconto

func CalcularDesconto(valor float64) float64 {
	if valor < 0 {
		return 0.0 // Regra para valores inválidos
	}
	if valor >= 500.0 {
		return valor * 0.20
	}
	if valor >= 100.0 {
		return valor * 0.10
	}
	return 0.0
}
