package pedido

import "errors"

// 1. Interface que define o contrato do serviço de pagamento
type ProvedorPagamento interface {
	Pagar(valor float64) bool
}

type ProcessadorPedido struct {
	Provedor ProvedorPagamento
}

// FecharPedido usa a interface para processar a cobrança
func (p *ProcessadorPedido) FecharPedido(valor float64) error {
	sucesso := p.Provedor.Pagar(valor)
	if !sucesso {
		return errors.New("falha no pagamento")
	}
	return nil
}
