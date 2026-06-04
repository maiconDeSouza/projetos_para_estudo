package guests

import "slices"

var guestsList []Guest

func GetGuests() []Guest {
	return guestsList
}

func NewGuests(g Guest) {
	guestsList = append(guestsList, g)
}

func ExitGuests(i int) {
	slices.Delete(guestsList, i, i+1)
}
