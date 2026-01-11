const dados = {
        mensagem: "Olá do servidor Bun!",
        versao: "1.0",
        ativo: true
      }

const server = Bun.serve({
    port: 1992,
    routes: {
        "/": () => Response.json(dados)
    }
})

console.log(`Listening on ${server.url}`)