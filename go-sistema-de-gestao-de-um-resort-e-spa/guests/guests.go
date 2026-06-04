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
	name    string
	isAdult bool
}

type service struct {
	name  string
	price float64
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

func (g Guest) GetTier() string {
	return g.tier.String()
}

func (g Guest) GetDiscount() float64 {
	return g.discount
}

func RegisterGuest(name string, tier Tier) Guest {
	return Guest{name: name, tier: tier, discount: tier.discount(), totalDependents: 4, days: 0}
}
