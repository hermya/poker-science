package model

import "poker_science/internal/app/domain/model/game"

type TableState struct {
	CardsOnTable []game.Card // 5 cards on the table
	Progress     int8        // number of open cards
	Pots         []PotValue  // Different Pots for current game
	Complete     bool
}
