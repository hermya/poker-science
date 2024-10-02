package model

type House int8

const (
	CLUBS House = iota
	SPADES
	HEARTS
	DIAMONDS
)

var (
	names  = [...]string{"CLUBS", "SPADE", "HEARTS", "DIAMONDS"}
	Houses = [...]House{CLUBS, SPADES, HEARTS, DIAMONDS}
)

func toHouseString(house House) string {
	return names[house]
}
