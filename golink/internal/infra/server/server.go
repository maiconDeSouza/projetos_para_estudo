package server

import (
	"fmt"
	"golink/internal/handler"
	"golink/internal/infra/env"
	"log"
	"net/http"
)

type Server struct {
	mux     *http.ServeMux
	env     *env.ENV
	handler handler.HandlerInterface
}

func NewServer(env *env.ENV, handler handler.HandlerInterface) *Server {
	return &Server{
		mux:     http.NewServeMux(),
		env:     env,
		handler: handler,
	}
}

func (s *Server) Run() {
	s.routes()
	fmt.Println("🚀 Servidor iniciado com sucesso na porta " + s.env.SERVER_PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", s.env.SERVER_PORT), s.mux))
}

func (s *Server) routes() {
	s.mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Maicon")
	})

	s.mux.HandleFunc("POST /admin", s.handler.CreateUserAdmin)
}
