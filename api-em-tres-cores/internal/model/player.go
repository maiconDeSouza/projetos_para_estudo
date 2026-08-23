package model

type Player struct {
	ID       [16]byte `json:"id"`
	Name     string   `json:"name"`
	Position string   `json:"position"`
	Number   uint     `json:"number"`
	Stats    Stats    `json:"stats"`
}

type Stats struct {
	Games         uint `json:"games"`
	MinutesPlayed uint `json:"minutes_played"`
	Goals         uint `json:"goals"`
	Assistis      uint `json:"asssistis"`
	YellowCards   uint `json:"yellow_cards"`
	RedCards      uint `json:"red_cards"`
}
