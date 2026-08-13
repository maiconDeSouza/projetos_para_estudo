package usuario

import (
	"errors"
	"strings"
)

type UsuarioRepository interface {
	BuscarPorEmail(email string) (*Usuario, error)
	Salvar(usuario *Usuario) error
}

type HasherSenha interface {
	GerarHash(senha string) (string, error)
}

type Usuario struct {
	ID        string
	Nome      string
	Email     string
	SenhaHash string
}

type ServicoUsuario struct {
	Repo   UsuarioRepository
	Hasher HasherSenha
}

var ErrEmailJaCadastrado = errors.New("Email já cadastrado")
var ErrCriptografia = errors.New("Erro ao Criptografar")
var ErrRepo = errors.New("Erro no Repositorio")
var ErrCamposInvalidos = errors.New("todos os campos (nome, email e senha) são obrigatórios e não podem estar vazios")

func (su *ServicoUsuario) Cadastrar(nome, email, senha string) (*Usuario, error) {
	if strings.TrimSpace(nome) == "" || strings.TrimSpace(email) == "" || strings.TrimSpace(senha) == "" {
		return nil, ErrCamposInvalidos
	}

	_, err := su.Repo.BuscarPorEmail(email)
	if errors.Is(err, ErrEmailJaCadastrado) {
		return nil, err
	}

	senhaCriptografada, err := su.Hasher.GerarHash(senha)
	if err != nil {
		return nil, ErrCriptografia
	}

	newUsuario := Usuario{Nome: nome, Email: email, SenhaHash: senhaCriptografada}

	err = su.Repo.Salvar(&newUsuario)
	if err != nil {
		return nil, ErrRepo
	}

	return &newUsuario, nil
}
