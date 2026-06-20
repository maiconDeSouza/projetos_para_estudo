package main

import (
	"api-library-go/internal/storage"
	"fmt"
	"net/http"
)

var Library = storage.Library{}

func init() {
	Library.ReadJSON()
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/library/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("oi"))
	})

	fmt.Println("Servidor rodando...")
	http.ListenAndServe(":2005", mux)
}
