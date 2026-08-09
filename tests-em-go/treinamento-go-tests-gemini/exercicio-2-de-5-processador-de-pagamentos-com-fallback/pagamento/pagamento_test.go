package pagamento

import (
	"errors"
	"testing"
)

type MockPrincipal struct {
	err error
}

func (mp *MockPrincipal) Pagar(valor float64) error {
	return mp.err
}

type MockSecundario struct {
	err error
}

func (ms *MockSecundario) Pagar(valor float64) error {
	return ms.err
}

func TestProcessar(t *testing.T) {
	t.Run("Sucesso no gateway Principal", func(t *testing.T) {
		mp := MockPrincipal{err: nil}
		ms := MockSecundario{err: errors.New("Erro!!!!")}
		p := ProcessadorPagamento{Principal: &mp, Secundario: &ms}
		valor := 14.99

		err := p.Processar(valor)
		if err != nil {
			t.Error("Era esperado que gateway Principal não falhasse")
		}
	})

	t.Run("Sucesso no gateway Secundário", func(t *testing.T) {
		mp := MockPrincipal{err: errors.New("Erro!!!!")}
		ms := MockSecundario{err: nil}
		p := ProcessadorPagamento{Principal: &mp, Secundario: &ms}
		valor := 14.99

		err := p.Processar(valor)
		if err != nil {
			t.Error("Era esperado que gateway Principal falhasse e gateway Segundário Funcionasse")
		}
	})

	t.Run("Falha em ambos os gateways", func(t *testing.T) {
		mp := MockPrincipal{err: errors.New("Erro!!!!")}
		ms := MockSecundario{err: errors.New("Erro!!!!")}
		p := ProcessadorPagamento{Principal: &mp, Secundario: &ms}
		valor := 14.99

		err := p.Processar(valor)
		if !errors.Is(err, ErrAmbos) {
			t.Error("Era esperado que ambos gateways falhassem")
		}
	})
}
