package services

import (
	"api-medical-system/internal/models"
	"api-medical-system/internal/repositories"
	"errors"
)

type ServicesInterface interface {
	CreateMedical(medical *models.MedicalRequest) (*models.Medical, error)
}

type Services struct {
	repositories repositories.DBResotirories
}

func (s *Services) CreateMedical(medical *models.MedicalRequest) (*models.Medical, error) {
	if len(medical.Crm) != 6 {
		return nil, errors.New("CRM Inválido")
	}

	if len(medical.Name) <= 3 {
		return nil, errors.New("Nome Inválido")
	}

	if len(medical.Specialty) == 0 {
		return nil, errors.New("Especialidade Inválido")
	}

	newMedical := models.Medical{
		Crm:       medical.Crm,
		Name:      medical.Name,
		Specialty: medical.Specialty,
		Agenda:    make([]*models.Agenda, 0),
	}

	err := s.repositories.AddMedical(&newMedical)
	if err != nil {
		return nil, err
	}

	return &newMedical, nil
}

func (s *Services) CreateAgenda(crm *models.GetMedical) error {
	addCrm := crm.Crm

	medical, err := s.repositories.GetMedical(addCrm)
	if err != nil {
		return err
	}

	medical.Agenda

	return nil
}

func NewServices(repositories repositories.DBResotirories) *Services {
	return &Services{
		repositories: repositories,
	}
}
