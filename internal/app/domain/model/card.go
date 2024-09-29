package model

type Card struct {
	Value Value
	House House
}

func cardToString(card Card) string {
	return ToValueString(card.Value) + " of " + toHouseString(card.House)
}
