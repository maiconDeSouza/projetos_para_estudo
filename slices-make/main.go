package main

import (
	"fmt"
)

func main() {
	sl := make([]int, 0, 3)
	fmt.Println(sl)

	sl = append(sl, 2, 3, 4, 5)

	fmt.Println(sl)
	fmt.Println(cap(sl))
	fmt.Println(len(sl))

}
