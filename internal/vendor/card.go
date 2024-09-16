package vendor

type Card struct {
	value Value
	house House
}

func cardToString(card Card) string{
	return toValueString(card.value) + " of " + toHouseString(card.house);
}