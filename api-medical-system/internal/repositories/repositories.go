package repositories

import (
	"api-medical-system/internal/models"
	"encoding/json"
	"errors"
	"log"
	"maps"
	"os"
)

type DBResotirories interface {
	WriteJson()
	AddMedical(medical *models.Medical) error
	AddAgenda(crm string, agenda []*models.Agenda) error
	GetMedical(crm string) (*models.Medical, error)
	GetDB() map[string]*models.Medical
}

type DB struct {
	FilePath string                     `json:"filePath"`
	DB       map[string]*models.Medical `json:"db"`
}

func (d *DB) WriteJson() {
	db := DB{
		FilePath: d.FilePath,
		DB:       make(map[string]*models.Medical),
	}

	maps.Copy(db.DB, d.DB)

	jsonDB, err := json.MarshalIndent(db, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(d.FilePath, jsonDB, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func (d *DB) ReadJson() {
	db := DB{}

	jsonDB, err := os.ReadFile(d.FilePath)
	if err != nil {
		if os.IsNotExist(err) {
			d.WriteJson()
			return
		}
		log.Fatal(err)
	}

	err = json.Unmarshal(jsonDB, &db)
	if err != nil {
		log.Fatal(err)
	}

	maps.Copy(d.DB, db.DB)
}

func (d *DB) AddMedical(medical *models.Medical) error {
	_, ok := d.DB[medical.Crm]
	if ok {
		return errors.New("CRM já existente!")
	}

	d.DB[medical.Crm] = medical
	d.WriteJson()

	_, ok = d.DB[medical.Crm]
	if !ok {
		return errors.New("Médico não salvo no banco de dados, tente novamente")
	}

	return nil
}

func (d *DB) GetMedical(crm string) (*models.Medical, error) {
	medical, ok := d.DB[crm]
	if !ok {
		return nil, errors.New("CRM não existente!")
	}
	return medical, nil
}

func (d *DB) GetDB() map[string]*models.Medical {
	return d.DB
}

func (d *DB) AddAgenda(crm string, agenda []*models.Agenda) error {
	medical, ok := d.DB[crm]
	if !ok {
		return errors.New("CRM Inválido")
	}

	medical.Agenda = append(medical.Agenda, agenda...)

	return nil
}

func NewRepositories() *DB {
	db := &DB{
		FilePath: "./db.json",
		DB:       make(map[string]*models.Medical),
	}

	db.ReadJson()

	return db
}
