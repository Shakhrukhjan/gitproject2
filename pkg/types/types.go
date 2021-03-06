package types

type Money int

type dirams Money

type Currency string
const (
	EUR Currency = "EUR"
)
type Transaction struct {
	Amount Money
	Type   string
	//Time   time.Time
}

type Card struct {
	Balance  Money
	Currency Currency
	Activity bool
}

var LocalCard = Card{
	Balance:  9000_00,
	Currency: Dollars,
	Type:     Gold, 
	Activity: Active,
	cvv:      112,
}

const (
	Platinum = "Platinum"
	Gold     = "Gold"
	Silver   = "Silver"
)

const (
	Somoni  = "TJS"
	Rubls   = "RUB"
	Dollars = "USD"
)

const (
	Active   = true
	InActive = false
)
