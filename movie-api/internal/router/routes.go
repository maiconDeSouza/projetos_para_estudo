package router

import (
	"fmt"
	"movies-api/internal/handlers"
	"movies-api/internal/middleware"
	"net/http"
)

type Routes struct {
	mux           *http.ServeMux
	handlersMovie handlers.HandlersMovieInterface
	HandlersUser  handlers.HandlersUserInterface
	versionAPI    string
}

func NewRoutes(
	mux *http.ServeMux,
	handlersMovie handlers.HandlersMovieInterface,
	HandlersUser handlers.HandlersUserInterface,
	versionAPI string) *Routes {
	return &Routes{mux: mux, handlersMovie: handlersMovie, HandlersUser: HandlersUser, versionAPI: versionAPI}
}

func (r *Routes) InitRoutes() {
	r.mux.HandleFunc(fmt.Sprintf("POST %s/user", r.versionAPI), r.HandlersUser.NewUser)
	r.mux.HandleFunc(fmt.Sprintf("POST %s/auth/login", r.versionAPI), r.HandlersUser.Login)

	r.mux.HandleFunc(fmt.Sprintf("GET %s/movies/all", r.versionAPI), middleware.AuthMiddleware(r.handlersMovie.AllMovies))
	r.mux.HandleFunc(fmt.Sprintf("GET %s/movies/{imdbID}", r.versionAPI), middleware.AuthMiddleware(r.handlersMovie.Movie))
	r.mux.HandleFunc(fmt.Sprintf("GET %s/movies/", r.versionAPI), middleware.AuthMiddleware(r.handlersMovie.SearchOMDB))
	r.mux.HandleFunc(fmt.Sprintf("POST %s/movies", r.versionAPI), middleware.AuthMiddleware(r.handlersMovie.NewMovie))
}
