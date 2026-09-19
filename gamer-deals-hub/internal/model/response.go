package model

type GameResponse struct {
	External string `json:"external"`
	Cheapest string `json:"cheapest"`
	GameID   string `json:"gameID"`
	Thumb    string `json:"thumb"`
}

type PriceResponse struct {
	CheapestPriceEver CheapestPriceEver `json:"cheapestPriceEver"`
	Deals             []Deals           `json:"deals"`
}

type CheapestPriceEver struct {
	Price string `json:"price"`
	Date  uint   `json:"date"`
}

type Deals struct {
	Price       string `json:"price"`
	RetailPrice string `json:"retailPrice"`
}
