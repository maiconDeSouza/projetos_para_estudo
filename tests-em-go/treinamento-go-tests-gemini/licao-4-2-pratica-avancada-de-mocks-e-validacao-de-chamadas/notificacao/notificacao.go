package notificacao

import "errors"

// Interface para o repositório de usuários
type RepositorioUsuario interface {
	BuscarEmailPorID(id int) (string, error)
}

// Interface para o serviço de envio de e-mails
type EnviadorEmail interface {
	Enviar(email, mensagem string) error
}

// Struct principal do serviço
type ServicoNotificacao struct {
	Repo  RepositorioUsuario
	Email EnviadorEmail
}

// Função a ser testada
func (s *ServicoNotificacao) NotificarUsuario(id int, mensagem string) error {
	email, err := s.Repo.BuscarEmailPorID(id)
	if err != nil {
		return errors.New("usuario nao encontrado")
	}

	err = s.Email.Enviar(email, mensagem)
	if err != nil {
		return errors.New("falha ao enviar email")
	}

	return nil
}
