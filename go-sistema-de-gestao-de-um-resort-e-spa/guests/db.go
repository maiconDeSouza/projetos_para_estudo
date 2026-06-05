package guests

import "slices"

var guestsList []guest

func GetGuests() []guest {
	return guestsList
}

func NewGuests(g guest) {
	guestsList = append(guestsList, g)
}

func ExitGuests(i int) {
	slices.Delete(guestsList, i, i+1)
}
