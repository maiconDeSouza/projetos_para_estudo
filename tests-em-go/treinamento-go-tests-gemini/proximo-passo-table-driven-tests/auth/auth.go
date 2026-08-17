package auth

import "errors"

// ---------------------------------------------------------------------------
// ERROS SENTINEL (Constantes de Erro)
// ---------------------------------------------------------------------------
// Criamos variáveis de erro exportadas (com letra maiúscula) para que a nossa
// aplicação e os testes possam comparar erros específicos por referência.

// ErrTokenAusente indica que a requisição nem sequer informou um token.
var ErrTokenAusente = errors.New("token ausente")

// ErrFalha indica um problema de infraestrutura ou rede com o serviço externo.
var ErrFalha = errors.New("falha na conexão com serviço")

// ErrToken indica que o serviço externo respondeu que o token é inválido ou expirou.
var ErrToken = errors.New("token expirado/inválido")

// ---------------------------------------------------------------------------
// INTERFACE (O Contrato)
// ---------------------------------------------------------------------------
// ProvedorAuth abstrai a comunicação com o serviço externo de autenticação.
// Ela exige que qualquer struct que a implemente possua o método ValidarToken.
// O método retorna 3 valores:
// 1. valido (bool): se o token é aceito ou não.
// 2. usuarioID (string): o ID do usuário associado ao token.
// 3. err (error): erros de comunicação/conexão durante a validação.
type ProvedorAuth interface {
	ValidarToken(token string) (valido bool, usuarioID string, err error)
}

// ---------------------------------------------------------------------------
// STRUCT PRINCIPAL (Injeção de Dependência)
// ---------------------------------------------------------------------------
// ValidadorSessao contém a regra de negócio do nosso sistema.
// Em vez de depender de uma implementação concreta, ela depende da interface ProvedorAuth.
type ValidadorSessao struct {
	Auth ProvedorAuth
}

// ---------------------------------------------------------------------------
// MÉTODO DE NEGÓCIO
// ---------------------------------------------------------------------------
// Autenticar executa o fluxo completo de validação da sessão do usuário.
func (v *ValidadorSessao) Autenticar(token string) (string, error) {
	// 1. Validação de Guarda (Early Return):
	// Se o token fornecido for uma string vazia, interrompe a execução imediatamente
	// sem gastar recursos chamando o serviço externo.
	if token == "" {
		return "", ErrTokenAusente
	}

	// 2. Chamada ao Serviço Externo via Interface:
	// Executa o método ValidarToken do ProvedorAuth injetado na struct.
	valido, usuarioID, err := v.Auth.ValidarToken(token)

	// 3. Tratamento de Erro de Conexão/Serviço:
	// Se o provedor retornar algum erro de infraestrutura, repassa esse erro para cima.
	if err != nil {
		return "", err
	}

	// 4. Tratamento de Token Inválido:
	// Se não houve erro de conexão, mas a resposta indicou que o token não é válido.
	if !valido {
		return "", ErrToken
	}

	// 5. Sucesso:
	// Se passou por todas as verificações, retorna o ID do usuário e nil para o erro.
	return usuarioID, nil
}
