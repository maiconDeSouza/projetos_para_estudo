package server

import (
	"fmt"
	"golink/internal/handler"
	"golink/internal/infra/env"
	"golink/internal/middleware"
	"log"
	"net/http"
)

type Server struct {
	mux        *http.ServeMux
	env        *env.ENV
	handler    handler.HandlerInterface
	middleware middleware.Middleware
}

func NewServer(env *env.ENV, handler handler.HandlerInterface, middleware middleware.Middleware) *Server {
	return &Server{
		mux:        http.NewServeMux(),
		env:        env,
		handler:    handler,
		middleware: middleware,
	}
}

func (s *Server) Run() {
	s.routes()
	fmt.Println("🚀 Servidor iniciado com sucesso na porta " + s.env.SERVER_PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", s.env.SERVER_PORT), s.mux))
}

func (s *Server) routes() {
	s.mux.HandleFunc("GET /test", s.middleware.AuthAdmin(s.handler.Test))
	s.mux.HandleFunc("GET /r/{title}", s.handler.Redirect)

	s.mux.HandleFunc("POST /auth/login", s.handler.Login)

	s.mux.HandleFunc("POST /user", s.handler.CreateUser)

	s.mux.HandleFunc("POST /auth/admin", s.middleware.AuthAdmin(s.handler.CreateUserAdmin))
	s.mux.HandleFunc("GET /auth/admin/all", s.middleware.AuthAdmin(s.handler.GetAllUsers))
	s.mux.HandleFunc("POST /auth/admin/link", s.middleware.AuthAdmin(s.handler.AddLinkAdmin))

}
