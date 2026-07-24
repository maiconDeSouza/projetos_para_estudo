package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, tasks chan int) {
	for t := range tasks {
		time.Sleep(time.Duration(100) * time.Millisecond)
		fmt.Printf("Trabalhador %d fez a tarefa %d\n", id, t)
	}
}

func main() {
	tasks := make(chan int, 20)
	var wg sync.WaitGroup

	for i := 1; i <= 20; i++ {
		tasks <- i
	}
	close(tasks)

	const numWorkers = 3
	for w := 1; w <= numWorkers; w++ {
		wg.Go(func() {
			worker(w, tasks)
		})
	}
	wg.Wait()
}
