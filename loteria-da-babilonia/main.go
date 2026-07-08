package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

type Game struct {
	randomNumber int
	chances      int
}

func ClearTerminal() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	g := Game{
		randomNumber: rand.Intn(15) + 1,
		chances:      0,
	}

	fmt.Println("Digite um número entre 1 à 15... Você tem 3 chances!")

	for {
		if g.chances >= 3 {
			fmt.Printf("Você usou suas três tentativas e não acertou! O número sorteado era %d", g.randomNumber)
			break
		}
		numberString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Erro ao ler:", err)
			break
		}

		numberString = strings.TrimSpace(numberString)
		number, err := strconv.Atoi(numberString)
		if err != nil {
			fmt.Println("Erro: A string não é um número válido!")
			break
		}

		if number > g.randomNumber {
			g.chances++
			ClearTerminal()
			fmt.Printf("O número é menor que %d", number)
			fmt.Println()
			continue
		}

		if number < g.randomNumber {
			g.chances++
			ClearTerminal()
			fmt.Printf("O número é maior que %d", number)
			fmt.Println()
			continue
		}
		ClearTerminal()
		fmt.Printf("Você acertou, o número é %d", number)
		break
	}
}
