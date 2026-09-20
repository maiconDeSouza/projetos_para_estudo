package router

import (
	"fmt"
	"movies-api/internal/handlers"
	"net/http"
)

type Routes struct {
	mux        *http.ServeMux
	handlers   handlers.HandlersInterface
	versionAPI string
}

func NewRoutes(mux *http.ServeMux, handlers handlers.HandlersInterface, versionAPI string) *Routes {
	return &Routes{mux: mux, handlers: handlers, versionAPI: versionAPI}
}

func (r *Routes) InitRoutes() {
	r.mux.HandleFunc(fmt.Sprintf("GET %s/movies/all", r.versionAPI), r.handlers.AllMovies)
	r.mux.HandleFunc(fmt.Sprintf("GET %s/movies/{imdbID}", r.versionAPI), r.handlers.Movie)
	r.mux.HandleFunc(fmt.Sprintf("GET %s/movies/", r.versionAPI), r.handlers.SearchOMDB)
	r.mux.HandleFunc(fmt.Sprintf("POST %s/movies", r.versionAPI), r.handlers.NewMovie)
}
