package model

type Transaction struct {
	FromUuid string  // UUID of player or room
	FromType int8    // indicator of who paid
	ToUuid   string  // UUID of player or room
	ToType   int8    // indicator of who was paid
	Amount   float64 // amount in $
	Complete bool
}
