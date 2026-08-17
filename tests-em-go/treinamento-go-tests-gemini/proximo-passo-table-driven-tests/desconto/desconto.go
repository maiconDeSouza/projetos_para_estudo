package desconto

type ValidadorVip interface {
	EhhVip(clienteID string) bool
}

type CalculadoraDesconto struct {
	VipService ValidadorVip
}

func (c *CalculadoraDesconto) Calcular(clienteID string, valorOriginal float64) float64 {
	if c.VipService.EhhVip(clienteID) {
		return valorOriginal * 0.90
	}

	return valorOriginal
}
