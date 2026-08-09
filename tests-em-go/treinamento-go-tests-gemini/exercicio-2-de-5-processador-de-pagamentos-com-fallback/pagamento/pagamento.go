package pagamento

import "errors"

type GatewayPagamento interface {
	Pagar(valor float64) error
}

var ErrAmbos = errors.New("ambos os gateways falharam")

type ProcessadorPagamento struct {
	Principal  GatewayPagamento
	Secundario GatewayPagamento
}

func (p *ProcessadorPagamento) Processar(valor float64) error {
	err := p.Principal.Pagar(valor)
	if err == nil {
		return nil
	}

	err = p.Secundario.Pagar(valor)
	if err == nil {
		return nil
	}

	return ErrAmbos
}
