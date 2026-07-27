package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type ViaCEPResponse struct {
	CEP        string `json:"cep,omitempty"`
	Logradouro string `json:"logradouro,omitempty"`
	Bairro     string `json:"bairro,omitempty"`
	Localidade string `json:"localidade,omitempty"`
	Uf         string `json:"uf,omitempty"`
	Erro       string `json:"erro,omitempty"`
}

func PegarCEP(cep string) (*ViaCEPResponse, error) {
	dados := ViaCEPResponse{}
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)

	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("falha de conexao na API ViaCEP: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusBadRequest {
		return nil, errors.New("formato de CEP incorreto")
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("erro inesperado: status %d", res.StatusCode)
	}

	err = json.NewDecoder(res.Body).Decode(&dados)
	if err != nil {
		return nil, fmt.Errorf("erro ao decodificar JSON: %w", err)
	}

	if dados.Erro == "true" || dados.Erro != "" {
		return nil, errors.New("CEP nao encontrado na base de dados")
	}

	return &dados, nil
}

func ReturnCep(result *ViaCEPResponse, err error) {
	if err != nil {
		fmt.Printf("Erro: %s\n", err)
		return
	}

	if result == nil {
		fmt.Println("Erro: dados da resposta nulos")
		return
	}

	fmt.Printf("CEP: %s - %s - Bairro: %s - %s/%s\n",
		result.CEP, result.Logradouro, result.Bairro, result.Localidade, result.Uf)
}

var ceps = []string{
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

type WaitGroup struct {
	sync.WaitGroup
}

func (wg *WaitGroup) Go(f func()) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		f()
	}()
}

func main() {
	fmt.Println("Processando sequencialmente...")
	inicioSemGoroutines := time.Now()
	for _, cep := range ceps {
		result, err := PegarCEP(cep)
		ReturnCep(result, err)
	}
	duracaoSemGoroutines := time.Since(inicioSemGoroutines)

	fmt.Println("----------------------------------------------------------------------------")
	fmt.Println("Iniciando processamento concorrente com 3 goroutines...")
	fmt.Println("----------------------------------------------------------------------------")

	cepChan := make(chan string, len(ceps))
	for _, cep := range ceps {
		cepChan <- cep
	}
	close(cepChan)

	var wg WaitGroup
	inicioComGoroutines := time.Now()

	for i := 0; i < 3; i++ {
		wg.Go(func() {
			for cep := range cepChan {
				result, err := PegarCEP(cep)
				ReturnCep(result, err)
			}
		})
	}
	wg.Wait()
	duracaoComGoroutines := time.Since(inicioComGoroutines)

	fmt.Println("\n============================================================================")
	fmt.Println("RELATORIO FINAL DE PERFORMANCE")
	fmt.Println("============================================================================")
	fmt.Printf("Tempo total SINCRONO (um por um):   %v\n", duracaoSemGoroutines)
	fmt.Printf("Tempo total ASSINCRONO (3 goroutines): %v\n", duracaoComGoroutines)
	fmt.Println("----------------------------------------------------------------------------")

	ganho := float64(duracaoSemGoroutines) / float64(duracaoComGoroutines)
	fmt.Printf("A concorrencia com 3 goroutines foi %.1fx mais veloz!\n", ganho)
	fmt.Println("============================================================================")
}
