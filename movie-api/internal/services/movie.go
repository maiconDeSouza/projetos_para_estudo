package services

import (
	"encoding/json"
	"fmt"
	"movies-api/internal/models"
	"movies-api/internal/repositories"
	"net/http"
	"net/url"
	"os"
	"time"
	"uuid"
)

type ServicesMovieInterface interface {
	NewMovie(id uuid.UUID, imdb models.ImdbIDRequest) (*models.Movie, error)
	SearchOMDB(q string) (*models.ResultSearchOMDB, error)
	AllMovies(id uuid.UUID) ([]models.Movie, error)
	Movie(id uuid.UUID, imdbID string) (*models.Movie, error)
}

type ServicesMovie struct {
	repo       repositories.RepoInterface
	httpClient *http.Client
}

func NewServicesMovie(repo repositories.RepoInterface) *ServicesMovie {
	return &ServicesMovie{repo: repo, httpClient: &http.Client{Timeout: time.Duration(10) * time.Second}}
}

func (s *ServicesMovie) AllMovies(id uuid.UUID) ([]models.Movie, error) {
	movies, err := s.repo.GetAllMovies(id)
	if err != nil {
		return nil, err
	}

	return movies, nil
}

func (s *ServicesMovie) Movie(id uuid.UUID, imdbID string) (*models.Movie, error) {
	movie, err := s.repo.GetMovie(id, imdbID)
	if err != nil {
		return nil, err
	}

	return movie, nil
}

func (s *ServicesMovie) NewMovie(id uuid.UUID, imdb models.ImdbIDRequest) (*models.Movie, error) {
	movieRequestOMDB := models.MovieRequestOMDB{}
	url := fmt.Sprintf("http://www.omdbapi.com/?apikey=%s&i=%s", os.Getenv("OMDB_APIKEY"), imdb.ImdbID)

	res, err := s.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&movieRequestOMDB); err != nil {
		return nil, err
	}

	if movieRequestOMDB.Response == "False" {
		return nil, fmt.Errorf("filme não encontrado: %s", movieRequestOMDB.Error)
	}

	newMovie := models.Movie{
		Title:       movieRequestOMDB.Title,
		Description: movieRequestOMDB.Plot,
		Year:        movieRequestOMDB.Year,
		Genre:       movieRequestOMDB.Genre,
		ImdbID:      movieRequestOMDB.ImdbID,
	}

	if err := s.repo.CreateMovie(id, newMovie); err != nil {
		return nil, err
	}

	return &newMovie, nil
}

func (s *ServicesMovie) SearchOMDB(q string) (*models.ResultSearchOMDB, error) {
	list := models.ResultSearchOMDB{}
	search := url.PathEscape(q)
	url := fmt.Sprintf("http://www.omdbapi.com/?apikey=%s&s=%s", os.Getenv("OMDB_APIKEY"), search)

	res, err := s.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&list); err != nil {
		return nil, err
	}

	if list.Response == "False" {
		return nil, fmt.Errorf("filme não encontrado: %s", list.Error)
	}

	return &list, nil
}
