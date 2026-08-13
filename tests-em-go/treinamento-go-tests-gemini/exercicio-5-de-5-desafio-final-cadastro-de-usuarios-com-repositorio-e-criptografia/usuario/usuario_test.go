package usuario

import (
	"errors"
	"testing"
)

type MockUsuarioRepository struct {
	email     string
	errSalvar error
	usuario   Usuario
}

func (mr *MockUsuarioRepository) BuscarPorEmail(email string) (*Usuario, error) {
	if email == mr.email {
		return &mr.usuario, ErrEmailJaCadastrado
	}
	return nil, nil
}

func (mr *MockUsuarioRepository) Salvar(usuario *Usuario) error {
	if mr.errSalvar != nil {
		return mr.errSalvar
	}

	return mr.errSalvar
}

type MockHasherSenha struct {
	funcHash bool
	err      error
}

func (mh *MockHasherSenha) GerarHash(senha string) (string, error) {
	mh.funcHash = true

	if mh.err != nil {
		return "", mh.err
	}

	return "#$%¨&", mh.err
}

func TestCadastrar(t *testing.T) {
	t.Run("Dados de entrada inválidos (campos vazios) 🚫", func(t *testing.T) {
		mr := MockUsuarioRepository{email: "", errSalvar: nil, usuario: Usuario{}}
		mk := MockHasherSenha{funcHash: false, err: nil}
		su := ServicoUsuario{Repo: &mr, Hasher: &mk}

		_, err := su.Cadastrar("", "", "")
		if !errors.Is(err, ErrCamposInvalidos) {
			t.Errorf("Era esperado [%v], mas recebemos [%v]", ErrCamposInvalidos, err)
		}
	})

	t.Run("E-mail já cadastrado 🛑", func(t *testing.T) {
		u := Usuario{ID: "a1", Nome: "Fulano", Email: "teste@email.com", SenhaHash: "#$%"}
		mr := MockUsuarioRepository{email: "teste@email.com", errSalvar: nil, usuario: u}
		mk := MockHasherSenha{funcHash: false, err: nil}
		su := ServicoUsuario{Repo: &mr, Hasher: &mk}

		_, err := su.Cadastrar("Jorge", "teste@email.com", "123")
		if !errors.Is(err, ErrEmailJaCadastrado) {
			t.Errorf("Era esperado [%v], mas recebemos [%v]", ErrEmailJaCadastrado, err)
		}
	})

	t.Run("Falha ao gerar hash da senha 🔐", func(t *testing.T) {
		u := Usuario{ID: "a1", Nome: "Fulano", Email: "teste@email.com", SenhaHash: "#$%"}
		mr := MockUsuarioRepository{email: "novo@email.com", errSalvar: nil, usuario: u}
		mk := MockHasherSenha{funcHash: true, err: ErrCriptografia}
		su := ServicoUsuario{Repo: &mr, Hasher: &mk}

		_, err := su.Cadastrar("Jorge", "teste@email.com", "123")
		if !errors.Is(err, ErrCriptografia) {
			t.Errorf("Era esperado [%v], mas recebemos [%v]", ErrCriptografia, err)
		}
	})

	t.Run("Sucesso no cadastro (verificando se o usuário foi salvo corretamente) 🟢", func(t *testing.T) {
		u := Usuario{ID: "a1", Nome: "Fulano", Email: "teste@email.com", SenhaHash: "#$%"}
		mr := MockUsuarioRepository{email: "novo@email.com", errSalvar: nil, usuario: u}
		mk := MockHasherSenha{funcHash: true, err: nil}
		su := ServicoUsuario{Repo: &mr, Hasher: &mk}

		newUsuario, err := su.Cadastrar("Jorge", "teste@email.com", "123")

		if err != nil {
			t.Errorf("Não era esperado o erro [%v]", err)
		}
		if newUsuario == nil {
			t.Errorf("Era esperado o usuário [%v]", newUsuario)
		}
	})

	t.Run("Erro ao salvar", func(t *testing.T) {
		u := Usuario{ID: "a1", Nome: "Fulano", Email: "teste@email.com", SenhaHash: "#$%"}
		mr := MockUsuarioRepository{email: "novo@email.com", errSalvar: ErrRepo, usuario: u}
		mk := MockHasherSenha{funcHash: true, err: nil}
		su := ServicoUsuario{Repo: &mr, Hasher: &mk}

		_, err := su.Cadastrar("Jorge", "teste@email.com", "123")
		if !errors.Is(err, ErrRepo) {
			t.Errorf("Era esperado [%v], mas recebemos [%v]", ErrRepo, err)
		}
	})
}
