package card
import ""
func Issue() types.Card {
	return types.Card{
		Balance:  0,
		Currency: types.Currency("EUR"),
	}
}

func Withdraw(card types.Card) types.Card {
	card.Balance -= 100
	return card
}
