package vendor

type House int8

const (
	CLUBS House = iota
	SPADES 
	HEARTS
	DIAMONDS
)

var(
	names = [...]string{"CLUBS", "SPADE", "HEARTS", "DIAMONDS"}
)

func toHouseString(house House) string {
	return names[house];
}