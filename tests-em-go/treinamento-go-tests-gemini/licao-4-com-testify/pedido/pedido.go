package pedido

import "errors"

type ProvedorPagamento interface {
	Pagar(valor float64) bool
}

type ProcessadorPedido struct {
	Provedor ProvedorPagamento
}

func (p *ProcessadorPedido) FecharPedido(valor float64) error {
	sucesso := p.Provedor.Pagar(valor)
	if !sucesso {
		return errors.New("falha no pagamento")
	}
	return nil
}
