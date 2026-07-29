package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func buscarPrecoProduto(ctx context.Context, produto string) (string, error) {
	ch := make(chan string, 1)

	go func() {
		time.Sleep(time.Duration(2) * time.Second)
		ch <- fmt.Sprintf("Produto [%s]: R$ 99,00", produto)
	}()

	select {
	case msg := <-ch:
		return msg, nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(rand.Intn(3)+1)*time.Second)
	defer cancel()

	msg, err := buscarPrecoProduto(ctx, "Notebook")
	if err != nil {
		fmt.Println("⏱️ Falha: operação cancelada por timeout!", err)
		return
	}

	fmt.Println("✅ Sucesso:", msg)
}
