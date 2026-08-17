package frete

type CalculadorDistancia interface {
	EhCapital(cidade string) bool
}

type ServicoFrete struct {
	Calc CalculadorDistancia
}

func (s *ServicoFrete) CalcularFrete(cidade string) float64 {
	if s.Calc.EhCapital(cidade) {
		return 10.00
	}
	return 30.00
}
