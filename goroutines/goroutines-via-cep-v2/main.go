package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type ViaCEPResponse struct {
	CEP        string `json:"cep,omitempty"`
	Logradouro string `json:"logradouro,omitempty"`
	Bairro     string `json:"bairro,omitempty"`
	Localidade string `json:"localidade,omitempty"`
	Uf         string `json:"uf,omitempty"`
	Erro       string `json:"erro,omitempty"`
}

type CEP struct{}

func (c CEP) get(cep string) string {
	dados := ViaCEPResponse{}
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)

	res, err := http.Get(url)
	if err != nil {
		return fmt.Sprintf("falha de conexão na API ViaCEP: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusBadRequest {
		return fmt.Sprintf("formato de CEP [%s] icorreto", cep)
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Sprintf("erro inesperado: status %d", res.StatusCode)
	}

	err = json.NewDecoder(res.Body).Decode(&dados)
	if err != nil {
		return fmt.Sprintf("erro ao decodificar JSON: %v", err)
	}

	if dados.Erro == "true" || dados.Erro != "" {
		return fmt.Sprintf("CEP [%s] nao encontrado na base de dados", cep)
	}

	return fmt.Sprintf("CEP: %s - %s - Bairro: %s - %s/%s",
		dados.CEP, dados.Logradouro, dados.Bairro, dados.Localidade, dados.Uf)
}

func (c CEP) Worker(ceps <-chan string, results chan<- string) {
	for cep := range ceps {
		results <- c.get(cep)
	}
}

var cepsList = []string{
	"01311000", "55555555", "70070000", "01002000", "30110001", "33333333", "04001001", "24020000",
	"21040360", "79002000", "36010000", "66013000", "13010000", "06010000", "32600000", "99999999",
	"74110010", "12210000", "22020002", "30140010", "88888888", "38010000", "22640100", "01310200",
	"09010000", "41830000", "69010000", "57020000", "77015000", "65010000", "60010000", "14400000",
	"22410003", "80020000", "29100000", "60060390", "11010000", "49010000", "40015000", "03102000",
	"11600000", "77777777", "35160000", "14010000", "20211110", "11700000", "17010000", "74003010",
	"90020000", "29010001", "88010000", "69900000", "05001000", "01001000", "76801000", "95010000",
	"22222222", "66666666", "08010000", "37130000", "38400000", "77001000", "25010000", "11111111",
	"96010000", "20020010", "44444444", "04101300", "02011000", "40020160", "29050000", "50030230",
	"06401000", "07110000", "31270901", "89010000", "68900000", "78015000", "00000000", "90010000",
	"32010000", "29160000", "15010000", "57030000", "58010000", "91010000", "18010000", "80010000",
	"16010000", "13400000", "50050000", "27910000", "04571010", "78900000", "30310000", "66040000",
	"89201000", "20040002", "07010000", "19010000", "41150000", "64000000", "69005000", "09710000",
}

func main() {
	wg := sync.WaitGroup{}
	cep := CEP{}
	ceps := make(chan string, len(cepsList))
	results := make(chan string, len(cepsList))
	workers := 6

	for _, cep := range cepsList {
		ceps <- cep
	}
	close(ceps)

	for i := range workers {
		fmt.Printf("Trabalhador %d rodando\n", i+1)
		wg.Go(func() {
			cep.Worker(ceps, results)
		})
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for r := range results {
		fmt.Println(r)
	}
}
