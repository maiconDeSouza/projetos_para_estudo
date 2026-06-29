package models

type PatientRequest struct {
	Name string `json:"name"`
	Data string `json:"data"`
	Time string `json:"time"`
}

type Agenda struct {
	Patient string `json:"patient"`
	Data    string `json:"data"`
	Time    string `json:"time"`
}

type GetMedical struct {
	Crm string `json:"crm"`
}

type Medical struct {
	Crm       string `json:"crm"`
	Name      string `json:"name"`
	Specialty string `json:"specialty"`
	Agenda    []*Agenda
}

type MedicalRequest struct {
	Crm       string `json:"crm"`
	Name      string `json:"name"`
	Specialty string `json:"specialty"`
}

type Scheduling struct {
	Patient string `json:"patient"`
	Medical string `json:"medical"`
	Data    string `json:"data"`
}
