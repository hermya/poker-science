package model

import "poker_science/internal/app/domain/model/game"

type PlayerState struct {
	RoomUuid    string      // Room with which the player is associated
	Status      int8        // tells whether player has matched, raised, has been matched, or fold
	TotalEntry  float64     // left with player since buyIn
	Role        int8        // is Small-blind, big blind, or other
	CardsInHand []game.Card // the two cards dealt to player
	PlayingPot  int8        // Refers to PotValue.PotPriority. Useful for capturing All-in
}
