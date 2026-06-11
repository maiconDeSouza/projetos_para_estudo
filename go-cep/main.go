package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

var ceps = []string{
	"01001000", // SP
	"20040002", // RJ
	"30140071", // MG
	"40010000", // BA
	"50010000", // PE
	"60060080", // CE
	"70040900", // DF
	"69005010", // AM
	"80010000", // PR
	"90010000", // RS
	"88010001", // SC
	"74003010", // GO
	"78005000", // MT
	"79002000", // MS
	"65010000", // MA
	"64000040", // PI
	"58010000", // PB
	"59010000", // RN
	"49010000", // SE
	"57020000", // AL
	"76801059", // RO
	"69301000", // RR
	"68900073", // AP
	"77001014", // TO
	"66010000", // PA

	"01310930",
	"01414000",
	"01501000",
	"02011000",
	"03010000",
	"04003000",
	"05010000",
	"06010000",
	"07010000",
	"08010000",

	"20230010",
	"20550013",
	"22010000",
	"22250040",
	"22410002",
	"22640010",
	"22793082",
	"23013000",
	"24020000",
	"26010000",

	"30112000",
	"30310000",
	"30510000",
	"31010000",
	"31270000",
	"31515000",
	"32010000",
	"34000000",
	"36010000",
	"38400000",

	"40015000",
	"40110000",
	"40210000",
	"40301000",
	"40410000",
	"40510000",
	"40670000",
	"40710000",
	"40810000",
	"40900000",

	"50030000",
	"50710000",
	"51010000",
	"52010000",
	"53020000",
	"54010000",
	"55002000",
	"56000000",
	"57035000",
	"58013000",

	"60025001",
	"60110000",
	"60325000",
	"60410000",
	"60510000",
	"60610000",
	"60740000",
	"60810000",
	"62010000",
	"63010000",

	"70070900",
	"70200900",
	"70390900",
	"70710900",
	"71010900",
	"71900900",
	"72010000",
	"72800000",
	"73010000",
	"74015000",

	"78010000",
	"78110000",
	"79010000",
	"80020000",
	"80510000",
	"81050000",
	"82010000",
	"88015000",
	"89010000",
	"90020000",
}

type Endereco struct {
	CEP         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	UF          string `json:"uf"`
}

func getCep(cep string) (Endereco, error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
	var endereco Endereco

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Erro ao consultar CEP:", err)
		return endereco, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&endereco)
	if err != nil {
		fmt.Println("Erro ao converter JSON:", err)
		return endereco, err
	}

	return endereco, nil
}

func render(end Endereco, i int) {
	fmt.Printf("%d) CEP: %s -> End: %s -> UF: %s\n", i+1, end.CEP, end.Logradouro, end.UF)

}

func main() {
	init := time.Now()
	var wg sync.WaitGroup

	for i, cep := range ceps {
		wg.Go(func() {
			end, err := getCep(cep)
			if err != nil {
				fmt.Println(err)
			}
			render(end, i)
		})
	}

	wg.Wait()
	fmt.Printf("Termonou em %s", time.Since(init))
}
