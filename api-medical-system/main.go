package main

import (
	"fmt"
	"time"
)

func main() {
	layout := "2006-01-02 15:04"
	data := time.Now().Local().Format(layout)

	hr, err := time.Parse(layout, data)
	if err != nil {
		fmt.Println("Erro aoo converter:", err)
		return
	}

	fmt.Println(hr)

	hj := time.Now()

	agendamento := hj.AddDate(0, 0, 30)
	fmt.Println(agendamento.Month())
}
