package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/library/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("oi"))
	})

	fmt.Println("Servidor rodando...")
	http.ListenAndServe(":2005", mux)
}
