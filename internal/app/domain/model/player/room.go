package player

import "poker_science/internal/app/domain/model/game"

type RoomMeta struct {
	RoomUuid     string
	Balance      float64
	CardsInHand  [2]game.Card
	CardsOnTable [5]game.Card
}
