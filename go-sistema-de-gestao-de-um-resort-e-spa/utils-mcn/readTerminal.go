package utilsmcn

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var render = bufio.NewReader(os.Stdin)

func ReadTerminal(err *error) string {
	op, errMain := render.ReadString('\n')
	err = &errMain
	if *err != nil {
		PrintMCN(fmt.Sprint("Erro ao ler o prompt:", err))
		SleepMCN(2)
		Clear()
	}

	op = strings.TrimSpace(op)

	return op
}
