package main

import (
	"context"
	"fmt"
	"igressos-goroutines/name"
	"igressos-goroutines/num"
	"igressos-goroutines/show"
	"sync"
	"time"
)

func Pay(ctx context.Context) error {
	chErr := make(chan error, 1)

	go func() {
		t := num.GenerateNumberRand(100, 500)
		time.Sleep(time.Duration(t) * time.Millisecond)
		chErr <- nil
	}()

	select {
	case err := <-chErr:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

func TicketOffice(orders <-chan *show.Order, results chan<- string, show *show.Show) {
	for order := range orders {
		if order.GetErr() != nil {
			show.AddPedidos(order)
			results <- order.GetErr().Error()
			continue
		}

		if err := show.DecrementTicket(order.GetAmount()); err != nil {
			results <- fmt.Sprintf("[%s]  [%s]", order.GetName(), err.Error())
			continue
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(300)*time.Millisecond)
		err := Pay(ctx)
		cancel()
		if err != nil {
			order.SetErr(err)
			show.AddPedidos(order)
			show.AddTicket(order.GetAmount())
			results <- fmt.Sprintf("[%s] erro no pagamento [%s]", order.GetName(), order.GetErr().Error())
			continue
		}

		results <- fmt.Sprintf("[%s] comprou com sucesso o ingresso para o show [%s]", order.GetName(), show.GetName())

	}
}

func main() {
	start := time.Now()
	newShow := show.NewShow("Belo", 1000)
	namesList := name.GenerateNames(num.GenerateNumberRand(500, 2500))

	order := make(chan *show.Order, len(namesList))
	results := make(chan string, len(namesList))

	for _, name := range namesList {
		nameCopy := name
		o := show.NewOrder(nameCopy, num.GenerateNumberRand(1, 6))

		order <- o
	}
	close(order)

	wg := sync.WaitGroup{}

	for i := 1; i <= 5; i++ {
		wg.Go(func() {
			TicketOffice(order, results, newShow)
		})
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for r := range results {
		fmt.Println(r)
	}

	fmt.Printf("Total de ingressos: %d\n", newShow.GetTotalTicket())
	fmt.Printf("%d tentaram comprar ingressos\n", len(namesList))
	fmt.Printf("Os ingressos acabaram em %.2f\n", time.Since(start).Seconds())
}
