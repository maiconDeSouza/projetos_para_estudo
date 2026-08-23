package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// ---------------------------------------------------------------------------
// 1. MOCK (Substituto para o Testify)
// ---------------------------------------------------------------------------
type MockAuth struct {
	mock.Mock
}

func (m *MockAuth) ValidarToken(token string) (valido bool, usuarioID string, err error) {
	args := m.Called(token)
	return args.Bool(0), args.String(1), args.Error(2)
}

// ---------------------------------------------------------------------------
// 2. TESTE EM TABELA (Table-Driven Test)
// ---------------------------------------------------------------------------
func TestAutenticar(t *testing.T) {
	// Definimos a estrutura da tabela e seus cenários
	testes := []struct {
		nome         string
		tokenEntrada string
		chamouMock   bool

		// Respostas simuladas do Mock
		mockValido    bool
		mockUsuarioID string
		mockErr       error

		// Retornos esperados da regra de negócio
		usuarioIDEsperado string
		erroEsperado      error
	}{
		{
			nome:              "Token ausente",
			tokenEntrada:      "",
			chamouMock:        false, // Retorna no Early Return (não chama o mock)
			usuarioIDEsperado: "",
			erroEsperado:      ErrTokenAusente,
		},
		{
			nome:              "Falha no serviço de autenticação",
			tokenEntrada:      "token_conexao_ruim",
			chamouMock:        true,
			mockValido:        false,
			mockUsuarioID:     "",
			mockErr:           ErrFalha,
			usuarioIDEsperado: "",
			erroEsperado:      ErrFalha,
		},
		{
			nome:              "Token inválido/expirado",
			tokenEntrada:      "token_expirado",
			chamouMock:        true,
			mockValido:        false,
			mockUsuarioID:     "",
			mockErr:           nil,
			usuarioIDEsperado: "",
			erroEsperado:      ErrToken,
		},
		{
			nome:              "Sucesso",
			tokenEntrada:      "token_valido",
			chamouMock:        true,
			mockValido:        true,
			mockUsuarioID:     "user123",
			mockErr:           nil,
			usuarioIDEsperado: "user123",
			erroEsperado:      nil,
		},
	}

	// Loop que executa cada linha da tabela
	for _, tt := range testes {
		t.Run(tt.nome, func(t *testing.T) {
			mockAuth := new(MockAuth)

			// Configura a expectativa do mock APENAS se a função realmente for chamá-lo
			if tt.chamouMock {
				mockAuth.On("ValidarToken", tt.tokenEntrada).Return(tt.mockValido, tt.mockUsuarioID, tt.mockErr)
			}

			// Injeta o mock no serviço
			v := ValidadorSessao{Auth: mockAuth}

			// Executa o método
			usuarioID, err := v.Autenticar(tt.tokenEntrada)

			// Asserções dos resultados
			assert.ErrorIs(t, err, tt.erroEsperado)
			assert.Equal(t, tt.usuarioIDEsperado, usuarioID)

			// Valida se as chamadas do mock bateram com a expectativa
			mockAuth.AssertExpectations(t)
		})
	}
}
