package calculation

import (
	"fmt"
	"poker_science/internal/app/domain/model/game"
	"poker_science/internal/app/domain/service/calculation"
)

func dummyCalculateFFScore(cards []game.Card) game.FFScore {
	ffHouseKeeper := calculation.GetNewHouseKeeper()
	for _, card := range cards {
		calculation.AddToHouseKeeper(card, ffHouseKeeper)
	}
	fmt.Println(ffHouseKeeper)
	return calculation.EvaluateFFScore(ffHouseKeeper)
}

func getSampleCardsForNoFlush() []game.Card {
	var cards []game.Card
	cards = append(cards, game.Card{Value: game.QUEEN, House: game.HEARTS})
	cards = append(cards, game.Card{Value: game.ACE, House: game.SPADES})
	cards = append(cards, game.Card{Value: game.THREE, House: game.DIAMONDS})
	cards = append(cards, game.Card{Value: game.FIVE, House: game.CLUBS})
	cards = append(cards, game.Card{Value: game.SEVEN, House: game.HEARTS})
	cards = append(cards, game.Card{Value: game.TEN, House: game.SPADES})
	cards = append(cards, game.Card{Value: game.KING, House: game.HEARTS})
	return cards
}

func getSampleCardsTieBreakingNoFlush() []game.Card {
	var cards []game.Card
	cards = append(cards, game.Card{Value: game.SEVEN, House: game.HEARTS})
	cards = append(cards, game.Card{Value: game.NINE, House: game.CLUBS})
	cards = append(cards, game.Card{Value: game.QUEEN, House: game.HEARTS})
	cards = append(cards, game.Card{Value: game.JACK, House: game.CLUBS})
	cards = append(cards, game.Card{Value: game.SEVEN, House: game.CLUBS})
	cards = append(cards, game.Card{Value: game.TEN, House: game.CLUBS})
	cards = append(cards, game.Card{Value: game.KING, House: game.HEARTS})
	return cards
}

func getSampleCardsForHardFlush() []game.Card {
	var cards []game.Card
	cards = append(cards, game.Card{Value: game.TWO, House: game.HEARTS})
	cards = append(cards, game.Card{Value: game.THREE, House: game.HEARTS})
	cards = append(cards, game.Card{Value: game.SIX, House: game.DIAMONDS})
	cards = append(cards, game.Card{Value: game.FIVE, House: game.CLUBS})
	cards = append(cards, game.Card{Value: game.SEVEN, House: game.HEARTS})
	cards = append(cards, game.Card{Value: game.TEN, House: game.HEARTS})
	cards = append(cards, game.Card{Value: game.FOUR, House: game.HEARTS})
	return cards
}

func getSampleCardsForMidFlush() []game.Card {
	var cards []game.Card
	cards = append(cards, game.Card{Value: game.EIGHT, House: game.HEARTS})
	cards = append(cards, game.Card{Value: game.TEN, House: game.DIAMONDS})
	cards = append(cards, game.Card{Value: game.QUEEN, House: game.HEARTS})
	cards = append(cards, game.Card{Value: game.ACE, House: game.HEARTS})
	cards = append(cards, game.Card{Value: game.KING, House: game.HEARTS})
	cards = append(cards, game.Card{Value: game.NINE, House: game.HEARTS})
	cards = append(cards, game.Card{Value: game.JACK, House: game.HEARTS})
	return cards
}

func getSampleCardsForPlatterFlush() []game.Card {
	var cards []game.Card
	cards = append(cards, game.Card{Value: game.TWO, House: game.HEARTS})
	cards = append(cards, game.Card{Value: game.THREE, House: game.HEARTS})
	cards = append(cards, game.Card{Value: game.FIVE, House: game.HEARTS})
	cards = append(cards, game.Card{Value: game.KING, House: game.HEARTS})
	cards = append(cards, game.Card{Value: game.QUEEN, House: game.HEARTS})
	cards = append(cards, game.Card{Value: game.SEVEN, House: game.HEARTS})
	cards = append(cards, game.Card{Value: game.FOUR, House: game.HEARTS})
	return cards
}
