package email

type EmailService interface {
	Enviar(destinatario string, mensagem string) error
}

func EnviarBoasVindas(es EmailService, email string) error {
	return es.Enviar(email, "Bem-vindo ao sistema!")
}
