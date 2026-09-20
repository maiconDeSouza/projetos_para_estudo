package repositories

import (
	"movies-api/internal/models"

	"gorm.io/gorm"
)

type RepoInterface interface {
	CreateMovie(movie models.Movie) error
	GetAllMovies() ([]models.Movie, error)
	GetMovie(imdbID string) (*models.Movie, error)
}

type Repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) CreateMovie(movie models.Movie) error {
	result := r.db.Create(&movie)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *Repo) GetAllMovies() ([]models.Movie, error) {
	movies := []models.Movie{}

	result := r.db.Find(&movies)
	if result.Error != nil {
		return nil, result.Error
	}

	return movies, nil
}

func (r *Repo) GetMovie(imdbID string) (*models.Movie, error) {
	movie := models.Movie{}

	err := r.db.Find(&movie, "imdb_id = ?", imdbID).First(&movie).Error
	if err != nil {
		return nil, err
	}

	return &movie, nil
}
