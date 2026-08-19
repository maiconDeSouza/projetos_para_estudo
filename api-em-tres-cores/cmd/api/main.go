package main

import (
	"fmt"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome a API em três cores 🇾🇪")
}

const PORT = 2005

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", homeHandler)

	fmt.Printf("Servidor rodando na porta: %d ... Bora Tricolor ⚽", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), mux))
}
