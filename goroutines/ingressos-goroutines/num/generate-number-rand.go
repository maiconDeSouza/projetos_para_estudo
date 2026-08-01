package num

import "math/rand/v2"

func GenerateNumberRand(min, max int) int {
	return min + rand.IntN(max-min+1)
}
