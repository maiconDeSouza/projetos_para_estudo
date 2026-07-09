package services

import "api-sistema/internal/interfaces"

type Service struct {
	db interfaces.Database
}

func NewServices(db interfaces.Database) *Service {
	services := &Service{
		db: db,
	}
	return services
}
