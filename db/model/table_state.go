package model

import (
	"poker_science/internal/domain/cardinfo"
)

type TableState struct {
	CardsOnTable []cardinfo.Card // 5 cards on the table
	Progress     int8            // number of open cards
	Pots         []PotValue      // Different Pots for current cardinfo
	Complete     bool
}
