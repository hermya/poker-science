package game

type Category int

const (
	RF Category = iota
	SF
	FK
	FH
	FF
	SS
	TK
	TP
	PP
	HC
)

var (
	fullName = [...]string{"Royal Flush", "Straight Flush", "Four Of A Kind",
		"Full House", "Flush", "Straight", "Three Of A Kind", "Two Pair",
		"Pair", "High Card"}
)

func getFullName(category Category) string {
	return fullName[category]
}
