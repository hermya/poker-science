package model

type Room struct {
	RoomUuid    string     // roomIdentifier
	Players     []Player   // Players on the table
	BuyIn       float64    // Minimum buy-in price ($500)
	BasePrice   float64    // Small-blind bet ($10)
	CurrentGame TableState // information about TableState
}
