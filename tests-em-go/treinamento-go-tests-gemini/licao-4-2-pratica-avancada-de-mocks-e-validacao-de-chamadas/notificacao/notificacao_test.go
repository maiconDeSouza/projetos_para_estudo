package notificacao

import (
	"fmt"
	"testing"
)

// Mock para RepositorioUsuario
type MockRepo struct {
	EmailRetorno string
	ErroRetorno  error
}

func (m MockRepo) BuscarEmailPorID(id int) (string, error) {
	return m.EmailRetorno, m.ErroRetorno
}

// Mock para EnviadorEmail
type MockEmail struct {
	Foienviado  bool
	ErroRetorno error
}

func (m *MockEmail) Enviar(email, mensagem string) error {
	if m.ErroRetorno == nil {
		m.Foienviado = true
	}
	return m.ErroRetorno
}

func TestNotificarUsuario(t *testing.T) {
	t.Run("Sucesso", func(t *testing.T) {
		mr := MockRepo{"dev@exemplo.com", nil}
		me := &MockEmail{true, nil}

		s := ServicoNotificacao{
			Repo:  mr,
			Email: me,
		}

		err := s.NotificarUsuario(1, "teste")
		if err != nil {
			t.Errorf("esperava erro nil, obteve %v", err)
		}

		if !me.Foienviado {
			t.Error("esperava que o e-mail tivesse sido enviado, mas não foi")
		}

	})

	t.Run("Usuario_Nao_Encontrado", func(t *testing.T) {
		mr := MockRepo{EmailRetorno: "", ErroRetorno: fmt.Errorf("db error")}
		me := &MockEmail{Foienviado: false}

		s := ServicoNotificacao{Repo: mr, Email: me}

		err := s.NotificarUsuario(2, "Relatório pronto!")
		if err == nil {
			t.Error("esperava erro ao buscar usuário, mas obteve nil")
		}

		if me.Foienviado {
			t.Error("não deveria ter enviado e-mail quando o usuário não é encontrado")
		}
	})
}
