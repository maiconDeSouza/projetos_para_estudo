package hotel

type room struct {
	Room      string
	Price     float64
	View      string
	Busy      bool
	IsItClean bool
}

type service struct {
	Name  string
	Price float64
}

type Hotel struct {
	Name     string
	Rooms    []room
	Services []service
}

func (h Hotel) GetRooms() []room {
	return h.Rooms
}

func NewHotel() Hotel {
	r1 := room{Room: "A", Price: 589, View: "Sacada para a piscina", Busy: false, IsItClean: true}
	r2 := room{Room: "B", Price: 450, View: "Sacada para o jardim", Busy: false, IsItClean: true}
	r3 := room{Room: "C", Price: 450, View: "Sacada para o jardim", Busy: false, IsItClean: true}
	r4 := room{Room: "D", Price: 420, View: "Sacada para a entrada do hotel", Busy: false, IsItClean: true}
	r5 := room{Room: "E", Price: 420, View: "Sacada para a entrada do hotel", Busy: false, IsItClean: true}
	r6 := room{Room: "F", Price: 650, View: "Vista panorâmica para o mar", Busy: false, IsItClean: true}
	r7 := room{Room: "G", Price: 650, View: "Vista panorâmica para o mar", Busy: false, IsItClean: true}
	r8 := room{Room: "H", Price: 520, View: "Vista para a piscina", Busy: false, IsItClean: true}
	r9 := room{Room: "I", Price: 520, View: "Vista para a piscina", Busy: false, IsItClean: true}
	r10 := room{Room: "J", Price: 480, View: "Vista para o jardim interno", Busy: false, IsItClean: true}
	r11 := room{Room: "K", Price: 480, View: "Vista para o jardim interno", Busy: false, IsItClean: true}
	r12 := room{Room: "L", Price: 720, View: "Cobertura com vista para o mar", Busy: false, IsItClean: true}
	r13 := room{Room: "M", Price: 720, View: "Cobertura com vista para o mar", Busy: false, IsItClean: true}
	r14 := room{Room: "N", Price: 390, View: "Janela para o estacionamento", Busy: false, IsItClean: true}
	r15 := room{Room: "O", Price: 390, View: "Janela para o estacionamento", Busy: false, IsItClean: true}

	s1 := service{Name: "Massagem relaxante", Price: 237}
	s2 := service{Name: "acesso à área VIP", Price: 435}
	s3 := service{Name: "Jantar romântico", Price: 320}
	s4 := service{Name: "Passeio de barco", Price: 580}
	s5 := service{Name: "Café da manhã especial no quarto", Price: 145}
	s6 := service{Name: "Aula particular de mergulho", Price: 690}

	h := Hotel{
		Name:     "Hotel Luxo e Lazer",
		Rooms:    []room{r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15},
		Services: []service{s1, s2, s3, s4, s5, s6},
	}

	return h
}
