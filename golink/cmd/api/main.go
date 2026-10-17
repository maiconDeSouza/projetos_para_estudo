package main

import (
	"golink/internal/handler"
	"golink/internal/infra/db"
	"golink/internal/infra/env"
	"golink/internal/infra/server"
	"golink/internal/middleware"
	"golink/internal/repository"
	"golink/internal/service"
)

func main() {
	env := env.NewENV()
	collection := db.NewConnectMongo(env)

	repo := repository.NewRepository(collection)
	service := service.NewService(repo, env)
	middleware := middleware.NewMiddleware(env)
	handler := handler.NewHandler(service, env)

	server := server.NewServer(env, handler, *middleware)
	server.Run()
}
