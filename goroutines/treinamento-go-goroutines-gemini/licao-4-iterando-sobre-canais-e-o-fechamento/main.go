package main

import (
	"fmt"
	"math/rand"
	"time"
)

func enviarEmails(ch chan string, emails ...string) {
	for _, email := range emails {
		time.Sleep(time.Duration(rand.Intn(5)+1) * time.Second)
		ch <- fmt.Sprintf("E-mail para cliente: %s", email)
	}

	close(ch)
}

func main() {
	emails := []string{
		"usuario1@ficticio.com",
		"contato@exemplo.org",
		"teste.a@dominio.net",
		"fulano@teste.com",
	}

	ch := make(chan string)
	start := time.Now()

	go enviarEmails(ch, emails...)

	for msg := range ch {
		fmt.Println(msg)
	}

	fmt.Printf("Todos os e-mails foram enviados com sucesso! em %.2f s\n", time.Since(start).Seconds())
}
