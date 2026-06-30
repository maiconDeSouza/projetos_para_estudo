package services

import (
	"api-medical-system/internal/models"
	"time"
)

func GenerateAgendaNow(h, m int) []*models.Agenda {
	var agendaSlice []*models.Agenda
	now := time.Now()

	baseDate := time.Date(now.Year(), now.Month(), now.Day(), h, m, 0, 0, now.Location())

	for i := 0; i < 30; i++ {
		currentDay := baseDate.AddDate(0, 0, i)

		for x := 0; x < 7; x++ {
			newTimeSlot := currentDay.Add(time.Duration(x) * time.Hour)

			ag := &models.Agenda{
				Data: newTimeSlot.Format("02/01/2006"),
				Time: newTimeSlot.Format("15:04"),
			}
			agendaSlice = append(agendaSlice, ag)
		}
	}
	return agendaSlice
}
