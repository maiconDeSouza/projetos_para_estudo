package quiz // ou o nome que você for dar ao pacote

type Cartao struct {
	Pergunta string
	Resposta string
}

func PegarPerguntas() map[string]Cartao {
	perguntas := make(map[string]Cartao)

	perguntas["Atribuição Curta"] = Cartao{Pergunta: "Qual símbolo é usado para a declaração curta de variáveis com inferência de tipo?", Resposta: ":="}
	perguntas["Repetição"] = Cartao{Pergunta: "Qual é a única estrutura de laço de repetição existente nativamente em Go?", Resposta: "for"}
	perguntas["Condicionais"] = Cartao{Pergunta: "Em um bloco if, os parênteses em volta da condição são obrigatórios? (sim/não)", Resposta: "não"}
	perguntas["Múltiplas Escolhas"] = Cartao{Pergunta: "Qual estrutura substitui o uso de múltiplos 'if else' encadeados?", Resposta: "switch"}
	perguntas["Endereço de Memória"] = Cartao{Pergunta: "Qual caractere é usado antes de uma variável para obter o seu endereço de memória?", Resposta: "&"}
	perguntas["Desreferenciação"] = Cartao{Pergunta: "Qual caractere é usado para acessar o valor apontado por um ponteiro?", Resposta: "*"}
	perguntas["Inicialização Especial"] = Cartao{Pergunta: "Qual função embutida é usada para inicializar slices, maps e channels?", Resposta: "make"}
	perguntas["Listas Dinâmicas"] = Cartao{Pergunta: "Arrays em Go têm tamanho fixo. Qual estrutura similar possui tamanho dinâmico?", Resposta: "slice"}
	perguntas["Adição de Elementos"] = Cartao{Pergunta: "Qual função embutida adiciona elementos ao final de um slice?", Resposta: "append"}
	perguntas["Ordenação"] = Cartao{Pergunta: "Mapas em Go mantêm a ordem de inserção dos itens? (sim/não)", Resposta: "não"}
	perguntas["Valores Descartáveis"] = Cartao{Pergunta: "Qual caractere usamos para ignorar um valor de retorno indesejado?", Resposta: "_"}
	perguntas["Concorrência"] = Cartao{Pergunta: "Qual palavra-chave inicia uma nova thread leve gerida pela linguagem?", Resposta: "go"}
	perguntas["Comunicação Segura"] = Cartao{Pergunta: "Qual estrutura é usada para trocar dados com segurança entre goroutines?", Resposta: "channel"}
	perguntas["Atraso de Execução"] = Cartao{Pergunta: "Qual palavra-chave agenda uma chamada de função para rodar apenas no final da função atual?", Resposta: "defer"}
	perguntas["Tratamento de Falhas"] = Cartao{Pergunta: "Erros em Go são tratados como exceções tradicionais ou como valores retornados? (valores/exceções)", Resposta: "valores"}
	perguntas["Visibilidade Pública"] = Cartao{Pergunta: "Para exportar uma função para outro pacote, a primeira letra do nome deve ser: (maiuscula/minuscula)", Resposta: "maiuscula"}
	perguntas["Modelagem de Dados"] = Cartao{Pergunta: "Qual palavra-chave é usada para definir uma estrutura de dados com múltiplos campos?", Resposta: "struct"}
	perguntas["Contratos e Polimorfismo"] = Cartao{Pergunta: "Qual estrutura define um conjunto de métodos obrigatórios, permitindo interfaces implícitas?", Resposta: "interface"}
	perguntas["Bootstrap"] = Cartao{Pergunta: "Qual função especial roda automaticamente antes da main() em um pacote?", Resposta: "init"}
	perguntas["Módulos"] = Cartao{Pergunta: "Qual sub-comando do 'go mod' inicia o rastreamento de dependências em um projeto novo?", Resposta: "init"}

	return perguntas
}
