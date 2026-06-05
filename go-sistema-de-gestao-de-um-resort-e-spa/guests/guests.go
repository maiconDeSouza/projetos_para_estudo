package guests

type Tier int

const (
	Standard Tier = iota
	Intermediate
	Premium
)

func (t Tier) String() string {
	switch t {
	case Standard:
		return "Standard"
	case Intermediate:
		return "Intermediate"
	case Premium:
		return "Premium"
	default:
		return "Unknown"
	}
}

func (t Tier) discount() float64 {
	switch t {
	case Standard:
		return 0.00
	case Intermediate:
		return 0.10
	case Premium:
		return 0.20
	default:
		return 0.00
	}
}

type dependents struct {
	Name    string
	IsAdult bool
}

type service struct {
	Name  string
	Price float64
}

type Guest struct {
	name            string
	tier            Tier
	discount        float64
	totalDependents uint
	dependents      []dependents
	pay             float64
	days            uint
	service         []service
}

func (g Guest) GetName() string {
	return g.name
}

func (g Guest) GetTier() string {
	return g.tier.String()
}

func (g Guest) GetDiscount() float64 {
	return g.discount
}

func (g Guest) GetDependent() []dependents {
	return g.dependents
}

func (g *Guest) AddDependent(name string, isAdult bool) {
	d := dependents{Name: name, IsAdult: isAdult}

	g.dependents = append(g.dependents, d)
}

func (g *Guest) GetService() []service {
	return g.service
}

func (g *Guest) AddService(name string, price float64) {
	s := service{Name: name, Price: price}
	g.service = append(g.service, s)
}

func (g Guest) GetDay() uint {
	return g.days
}

func (g *Guest) AddDay() {
	g.days = g.days + 1
}

func RegisterGuest(name string, tier Tier) Guest {
	return Guest{name: name, tier: tier, discount: tier.discount(), totalDependents: 4, days: 0}
}
