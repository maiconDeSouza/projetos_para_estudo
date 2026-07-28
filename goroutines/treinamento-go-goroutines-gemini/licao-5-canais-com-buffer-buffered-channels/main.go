package main

import "fmt"

var audios = []string{"episodio1.mp3", "episodio2.mp3", "episodio3.mp3"}

func processarFila(ch chan string) {
	for audio := range ch {
		fmt.Printf("Processando [%s]...\n", audio)
	}
}

func main() {
	ch := make(chan string, len(audios))

	for _, audio := range audios {
		ch <- audio
	}
	close(ch)

	processarFila(ch)
}
