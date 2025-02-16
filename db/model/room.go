package model

type Room struct {
	HostUuid    string
	HostName    string
	RoomUuid    string   // roomIdentifier
	Players     []Player // Players on the table
	TableSize   int8
	TimeOut     int8
	BuyIn       float64    // Minimum buy-in price ($500)
	BasePrice   float64    // Small-blind bet ($10)
	CurrentGame TableState // information about TableState
}
