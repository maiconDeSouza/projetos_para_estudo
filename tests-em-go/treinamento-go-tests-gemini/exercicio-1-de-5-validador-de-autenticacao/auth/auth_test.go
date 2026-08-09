package auth

import (
	"errors"
	"testing"
)

type MockProvedorAuth struct {
	valido    bool
	usuarioID string
	err       error
}

func (m *MockProvedorAuth) ValidarToken(token string) (valido bool, usuarioID string, err error) {
	valido = m.valido
	usuarioID = m.usuarioID
	err = m.err
	return
}

func TestAutenticador(t *testing.T) {
	t.Run("Token Vazio", func(t *testing.T) {
		m := MockProvedorAuth{valido: false, usuarioID: "", err: nil}
		v := ValidadorSessao{Auth: &m}

		usuarioID, err := v.Autenticar("")
		if err == nil {
			t.Error("Era esperado um erro para token vazio")
		}

		if usuarioID != "" {
			t.Error("Era esperado um usuárioID vazio para token vazio")
		}

		if !errors.Is(err, ErrTokenAusente) {
			t.Errorf("Era esperado esse erro: [%v] para token vazio", ErrTokenAusente)
		}
	})

	t.Run("Token expirado/inválido", func(t *testing.T) {
		m := MockProvedorAuth{valido: false, usuarioID: "", err: nil}
		v := ValidadorSessao{Auth: &m}

		usuarioID, err := v.Autenticar("token_expirado")

		if usuarioID != "" {
			t.Error("Era esperado um usuárioID vazio")
		}

		if !errors.Is(err, ErrToken) {
			t.Errorf("Era esperado o erro [%v], mas retornou [%v]", ErrToken, err)
		}
	})

	t.Run("Token expirado/inválido", func(t *testing.T) {
		m := MockProvedorAuth{valido: false, usuarioID: "", err: ErrToken}
		v := ValidadorSessao{Auth: &m}

		usuarioID, err := v.Autenticar("user1")
		if errors.Is(err, ErrTokenAusente) {
			t.Errorf("Era esperado o erro: [%v] - mas retornou [%v]", ErrToken, ErrTokenAusente)
		}

		if usuarioID != "" {
			t.Error("Era esperado um usuárioID vazio para token vazio")
		}

		if !errors.Is(err, ErrToken) {
			t.Errorf("Era esperado esse erro: [%v] para token vazio", ErrToken)
		}
	})

	t.Run("Autenticação com sucesso", func(t *testing.T) {
		u := "user1"
		m := MockProvedorAuth{valido: true, usuarioID: u, err: nil}
		v := ValidadorSessao{Auth: &m}

		usuarioID, err := v.Autenticar(u)
		if err != nil {
			t.Errorf("Não era esperado nenhum erro, mas foi recebido [%v]", err)
		}

		if usuarioID != u {
			t.Errorf("O usuário devolvido foi [%s], mas era esperado [%s]", usuarioID, u)
		}
	})
}
