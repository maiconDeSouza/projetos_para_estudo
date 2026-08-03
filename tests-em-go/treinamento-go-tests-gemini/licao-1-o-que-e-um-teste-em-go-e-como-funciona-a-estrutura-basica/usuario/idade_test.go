package usuario

import "testing"

type ResponseTestIdade struct {
	idade int
	res   bool
}

var idades = [...]ResponseTestIdade{
	{23, true},
	{20, true},
	{18, true},
	{17, false},
	{0, false},
	{55, true},
	{-23, false},
	{21, true},
	{15, false},
	{1, false},
}

func TestEhMaiorDeIdade(t *testing.T) {
	for _, i := range idades {
		r := EhMaiorDeIdade(i.idade)

		if r != i.res {
			t.Errorf("EhMaiorDeIdade(%d) = %v; esperado %v", i.idade, r, i.res)
		}
	}
}
