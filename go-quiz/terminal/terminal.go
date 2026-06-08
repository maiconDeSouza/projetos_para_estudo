package terminal

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func Sleep(s uint) {
	time.Sleep(time.Duration(s) * time.Second)
}

func PrintTrace() {
	fmt.Println("-----------------------------------------")
}

func Clear() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

func ScanTerminal() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	op := scanner.Text()

	return op
}
