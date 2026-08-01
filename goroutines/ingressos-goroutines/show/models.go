package show

import (
	"errors"
	"fmt"
	"sync"
)

type Order struct {
	id     int
	name   string
	amount int
	err    error
	mu     sync.RWMutex
}

func (o *Order) GetName() string {
	o.mu.RLock()
	defer o.mu.RUnlock()
	return o.name
}

func (o *Order) GetErr() error {
	o.mu.RLock()
	defer o.mu.RUnlock()
	return o.err
}

func (o *Order) SetErr(e error) {
	o.mu.Lock()
	defer o.mu.Unlock()
	o.err = e
}

func NewOrder(name string, amount int) *Order {
	if amount > 5 {
		pedido := &Order{
			name:   name,
			amount: amount,
		}
		pedido.err = fmt.Errorf("[%s] não pode comprar mais de 5 ingressos para o show", pedido.name)
		return pedido
	}

	pedido := &Order{
		name:   name,
		amount: amount,
		err:    nil,
	}
	return pedido
}

type Show struct {
	name    string
	tickets int
	orders  []*Order
	mu      sync.RWMutex
}

func (s *Show) GetName() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.name
}

func (s *Show) AddPedidos(order *Order) {
	s.mu.Lock()
	defer s.mu.Unlock()
	id := len(s.orders) + 1
	order.id = id
	s.orders = append(s.orders, order)
}

func (s *Show) DecrementTicket() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.tickets <= 0 {
		return errors.New("Acabou os ingressos!!!!")
	}
	s.tickets--
	return nil
}

func (s *Show) GetTotalTicket() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.tickets
}

func NewShow(name string, tickets int) *Show {
	show := &Show{
		name:    name,
		tickets: tickets,
		orders:  make([]*Order, 0),
	}

	return show
}
