package pedido

import "testing"

// 2. Criamos o Mock que satisfaz a interface ProvedorPagamento
type MockPagamento struct {
	DevePassar bool // Controlamos o comportamento do mock aqui
}

func (m MockPagamento) Pagar(valor float64) bool {
	return m.DevePassar
}

// 3. Teste unitário isolado de rede
func TestFecharPedido_Sucesso(t *testing.T) {
	mock := MockPagamento{DevePassar: true}
	processador := ProcessadorPedido{Provedor: mock}

	err := processador.FecharPedido(100.0)
	if err != nil {
		t.Errorf("Esperava sucesso, mas obteve erro: %v", err)
	}
}

func TestFecharPedido_Falha(t *testing.T) {
	mock := MockPagamento{DevePassar: false}
	processador := ProcessadorPedido{Provedor: mock}

	if err := processador.FecharPedido(100.0); err == nil {
		t.Errorf("Esperava um erro, mas a função retornou nil: %v", err)
	}
}
