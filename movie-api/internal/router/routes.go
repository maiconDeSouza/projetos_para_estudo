package router

import (
	"fmt"
	"net/http"
)

type Routes struct {
	mux *http.ServeMux
}

func NewRoutes(mux *http.ServeMux) *Routes {
	return &Routes{mux: mux}
}

func (r *Routes) InitRoutes() {
	r.mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Rodou 🌐​")
	})
}
