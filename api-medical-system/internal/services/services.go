package services

import (
	"api-medical-system/internal/models"
	"api-medical-system/internal/repositories"
	"errors"
)

type ServicesInterface interface {
	CreateMedical(medical *models.MedicalRequest) (*models.Medical, error)
	CreateAgenda(agenda *models.MedicalAgenda) error
	GetMedical(crm string) (*models.Medical, error)
	DeleteMedical(crm string) (*models.Medical, error)
	PatientScheduling(patient *models.PatientRequest, crm string) (*models.Scheduling, []*models.Agenda, error)
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

func (s *Services) CreateAgenda(agenda *models.MedicalAgenda) error {
	addCrm := agenda.Crm

	medical, err := s.repositories.GetMedical(addCrm)
	if err != nil {
		return err
	}

	if len(medical.Agenda) > 0 {
		return errors.New("Agenda de 30 dias já criada")
	}

	newAgenda := GenerateAgendaNow(agenda.Hours, agenda.Minutes)
	medical.Agenda = append(medical.Agenda, newAgenda...)

	s.repositories.WriteJson()

	return nil
}

func (s *Services) GetMedical(crm string) (*models.Medical, error) {
	if len(crm) != 6 {
		return nil, errors.New("CRM Inválido")
	}

	medical, err := s.repositories.GetMedical(crm)
	if err != nil {
		return nil, err
	}

	return medical, nil
}

func (s *Services) DeleteMedical(crm string) (*models.Medical, error) {
	if len(crm) != 6 {
		return nil, errors.New("CRM Inválido")
	}

	medical, err := s.repositories.GetMedical(crm)
	if err != nil {
		return nil, err
	}

	delete(s.repositories.GetDB(), medical.Crm)

	s.repositories.WriteJson()

	return medical, nil
}

func (s *Services) PatientScheduling(patient *models.PatientRequest, crm string) (*models.Scheduling, []*models.Agenda, error) {
	if len(crm) != 6 {
		return nil, nil, errors.New("CRM Inválido")
	}

	medical, err := s.repositories.GetMedical(crm)
	if err != nil {
		return nil, nil, err
	}

	ag := []*models.Agenda{}
	sch := models.Scheduling{}

	for _, v := range medical.Agenda {
		if v.Data == patient.Data && v.Time == patient.Time && v.Patient == "" {
			v.Patient = patient.Name
			s.repositories.WriteJson()
			sch.Data = patient.Data
			sch.Medical = medical.Name
			sch.Patient = patient.Name
			sch.Time = patient.Time
			return &sch, nil, nil
		}
	}

	for _, v := range medical.Agenda {
		if v.Data == patient.Data && v.Patient == "" {
			ag = append(ag, v)
		}
	}

	if len(ag) > 0 {
		return nil, ag, nil
	}

	return nil, nil, errors.New("O médico não atende nesta data.")
}

func NewServices(repositories repositories.DBResotirories) *Services {
	return &Services{
		repositories: repositories,
	}
}
