package main

import "fmt"

func sequence(start, end int) <-chan int {
	s := make(chan int)

	go func() {
		for i := start; i <= end; i++ {
			s <- i
		}
		close(s)
	}()

	return s
}

func multiply(in <-chan int, factor int) <-chan int {
	m := make(chan int)

	go func() {
		for n := range in {
			m <- n * factor
		}
		close(m)
	}()

	return m
}

func keepEven(in <-chan int) <-chan int {
	k := make(chan int)

	go func() {
		for n := range in {
			if n%2 == 0 {
				k <- n
			}
		}
		close(k)
	}()

	return k
}

func main() {
	s := sequence(23, 187)
	m := multiply(s, 2)
	k := keepEven(m)

	for n := range k {
		fmt.Println(n)
	}
}
