package repositories

import (
	"errors"
	"movies-api/internal/models"
	"uuid"

	"gorm.io/gorm"
)

type RepoInterface interface {
	CreateMovie(id uuid.UUID, movie models.Movie) error
	GetAllMovies(id uuid.UUID) ([]models.Movie, error)
	GetMovie(id uuid.UUID, imdbID string) (*models.Movie, error)
	CreateUser(user models.User) (*models.User, error)
	GetUser(login models.Login) (*models.User, error)
	DeleteMovie(id uuid.UUID, imdbID string) error
}

type Repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) CreateMovie(id uuid.UUID, movie models.Movie) error {
	user := models.User{}
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return err
	}

	existingMovie := models.Movie{}
	err = r.db.Where("imdb_id = ?", movie.ImdbID).First(&existingMovie).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		if err = r.db.Create(&movie).Error; err != nil {
			return err
		}
	} else if err != nil {
		return err
	} else {
		movie = existingMovie
	}

	err = r.db.Where("user_id = ? AND movie_id = ?", user.ID, movie.ID).First(&models.UserMovie{}).Error
	if err == nil {
		return errors.New("filme já cadastrado na sua lista")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	err = r.db.Model(&user).Association("Movie").Append(&movie)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) GetAllMovies(id uuid.UUID) ([]models.Movie, error) {
	movies := []models.Movie{}
	user := models.User{}

	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}

	err = r.db.Model(&user).Association("Movie").Find(&movies)
	if err != nil {
		return nil, err
	}

	return movies, nil
}

func (r *Repo) GetMovie(id uuid.UUID, imdbID string) (*models.Movie, error) {
	movie := models.Movie{}
	user := models.User{}

	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}

	err = r.db.Model(&user).Where("imdb_id = ?", imdbID).Association("Movie").Find(&movie)
	if err != nil {
		return nil, err
	}

	return &movie, nil
}

func (r *Repo) DeleteMovie(userID uuid.UUID, imdbID string) error {
	var movie models.Movie
	if err := r.db.Where("imdb_id = ?", imdbID).First(&movie).Error; err != nil {
		return err // Retorna erro se o filme nem existir no banco
	}

	result := r.db.Where("user_id = ? AND movie_id = ?", userID, movie.ID).Delete(&models.UserMovie{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *Repo) CreateUser(user models.User) (*models.User, error) {
	result := r.db.Create(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (r *Repo) GetUser(login models.Login) (*models.User, error) {
	user := models.User{}

	err := r.db.Where("email = ?", login.Email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
