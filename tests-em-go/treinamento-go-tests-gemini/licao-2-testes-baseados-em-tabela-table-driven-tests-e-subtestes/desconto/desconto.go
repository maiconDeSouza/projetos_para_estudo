package desconto

// CalcularDesconto retorna o valor do desconto aplicado:
// - Compras a partir de 100.0 recebem 10% de desconto.
// - Compras a partir de 500.0 recebem 20% de desconto.
// - Compras abaixo de 100.0 não recebem desconto (0.0).
func CalcularDesconto(valor float64) float64 {
	if valor >= 500.0 {
		return valor * 0.20
	}
	if valor >= 100.0 {
		return valor * 0.10
	}
	return 0.0
}
