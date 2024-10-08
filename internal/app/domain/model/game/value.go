package game

type Value int8

const (
	TWO Value = iota
	THREE
	FOUR
	FIVE
	SIX
	SEVEN
	EIGHT
	NINE
	TEN
	JACK
	QUEEN
	KING
	ACE
)

var (
	cardValues  = [...]int8{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	cardStrings = [...]string{"TWO", "THREE", "FOUR", "FIVE", "SIX", "SEVEN", "EIGHT", "NINE", "TEN", "JACK", "QUEEN", "KING", "ACE"}
	Cards       = [...]Value{TWO, THREE, FOUR, FIVE, SIX, SEVEN, EIGHT, NINE, TEN, JACK, QUEEN, KING, ACE}
)

func GetOperationalValue(v Value) int8 {
	return (int8(v) + 2)
}

func ToValueString(value Value) string {
	return cardStrings[value]
}
