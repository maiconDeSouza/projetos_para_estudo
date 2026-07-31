package numeros

import "math/rand"

func Aleatorios(max int) int {
	return rand.Intn(int(max)) + 1
}
