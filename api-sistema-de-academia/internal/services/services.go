package services

import (
	"api-sistema/internal/interfaces"
	"api-sistema/internal/models"
	"errors"
)

type Service struct {
	db interfaces.Database
}

func (s *Service) NewUser(newUser *models.UserRequest) error {
	user := models.User{
		NickName: newUser.NickName,
		Password: newUser.Password,
	}

	if newUser.Category != string(models.Basic) || newUser.Category != string(models.Premium) {
		return errors.New("Categoria invalida")
	}

	user.Category = models.Category(newUser.Category)

	return s.db.CreateUser(&user)
}

func NewServices(db interfaces.Database) *Service {
	services := &Service{
		db: db,
	}
	return services
}
