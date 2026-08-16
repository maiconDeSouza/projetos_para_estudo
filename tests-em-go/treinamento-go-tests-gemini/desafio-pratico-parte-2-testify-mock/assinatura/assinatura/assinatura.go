package assinatura

type GatewayPagamento interface {
	CancelarAssinatura(idUsuario string) error
}

func CancelarPlano(gp GatewayPagamento, idUsuario string) error {
	return gp.CancelarAssinatura(idUsuario)
}
