package main

import (
	"fmt"
	"sync"
	"time"
)

type Participants struct {
	name string
	pub  []*PubSub
}

func (p *Participants) Create(name string) {
	p.name = name
}

func (p *Participants) SignUp(pub *PubSub) {
	p.pub = append(p.pub, pub)

	pub.Subscribe(*p)
}

type PubSub struct {
	nameCanal   string
	mu          sync.Mutex
	subscribers []*Participants
}

func (p *PubSub) Create(name string) {
	p.nameCanal = name
}

func (p *PubSub) Subscribe(participants Participants) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.subscribers = append(p.subscribers, &participants)
}

func (p *PubSub) Publish(nameVideo string) <-chan string {
	var wg sync.WaitGroup
	msg := make(chan string)
	for _, sub := range p.subscribers {
		wg.Go(func() {
			msg <- fmt.Sprintf("O vídeo [%s] está sendo enviado pelo criador [%s] para o inscrito [%s]", nameVideo, p.nameCanal, sub.name)
		})
	}
	go func() {
		wg.Wait()
		close(msg)
	}()
	return msg
}

func main() {
	channel := PubSub{}
	channel.Create("Aprenda Go")

	names := []string{
		"Ana", "Bruno", "Carlos", "Daniela", "Eduardo",
		"Fernanda", "Gabriel", "Helena", "Igor", "Juliana",
		"Lucas", "Mariana", "Nicolas", "Olivia", "Pedro",
		"Renata", "Samuel", "Tatiane", "Vitor", "Yasmin",
	}

	for _, name := range names {
		sub := Participants{}
		sub.Create(name)
		sub.SignUp(&channel)
	}

	msg := channel.Publish("Go - primeira aula")
	for m := range msg {
		time.Sleep(time.Duration(2) * time.Second)
		fmt.Printf("%s -> %s\n", time.Now().Format("15:04:05.000"), m)
	}

	fmt.Println()
	fmt.Println("-----------------Com go Routines e todos recebendo ao mesmo tempo-----------------")
	var wg sync.WaitGroup
	channel2 := PubSub{}
	channel2.Create("Praticando Logica")

	for _, name := range names {
		sub := Participants{}
		sub.Create(name)
		sub.SignUp(&channel2)
	}

	msg2 := channel2.Publish("Primeira aula de Logíca")

	for m := range msg2 {
		wg.Go(func() {
			time.Sleep(time.Duration(2) * time.Second)
			fmt.Printf("%s -> %s\n", time.Now().Format("15:04:05.000"), m)
		})
	}
	wg.Wait()
}
