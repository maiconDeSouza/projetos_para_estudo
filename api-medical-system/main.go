package main

import (
	"fmt"
	"strconv"
	"time"
)

type FormatData struct {
	day     string
	weekday string
	month   string
	year    string
}

func traduzirDia(d time.Weekday) string {
	dias := map[time.Weekday]string{
		time.Monday:    "Segunda-feira",
		time.Tuesday:   "Terça-feira",
		time.Wednesday: "Quarta-feira",
		time.Thursday:  "Quinta-feira",
		time.Friday:    "Sexta-feira",
	}
	return dias[d]
}

func main() {
	// layout := "2006-01-02 15:04"
	// day := time.Now().Local().Format(layout)

	// hr, err := time.Parse(layout, data)
	// if err != nil {
	// 	fmt.Println("Erro aoo converter:", err)
	// 	return
	// }

	// fmt.Println(hr)

	// hj := time.Now()

	// agendamento := hj.AddDate(0, 0, 30)
	// fmt.Println(agendamento)

	// layout := "2006-01-02 15:04"
	day := time.Now()

	days := []*FormatData{}

	for i := 0; i <= 30; i++ {
		ag := day.AddDate(0, 0, i)

		if ag.Weekday() == time.Saturday || ag.Weekday() == time.Sunday {
			continue
		}

		newData := FormatData{
			day:     strconv.Itoa(ag.Day()),
			weekday: traduzirDia(ag.Weekday()),
			month:   ag.Month().String(),
			year:    strconv.Itoa(ag.Year()),
		}

		days = append(days, &newData)
	}

	for _, v := range days {
		fmt.Println(*v)
	}
}
