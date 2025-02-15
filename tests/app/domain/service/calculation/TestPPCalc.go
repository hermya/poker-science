package calculation

import (
	"fmt"
	"poker_science/internal/app/domain/model/game"
	"poker_science/internal/app/domain/service/calculation"
)

func dummyCalculatePPScore(cards []game.Card) int8 {
	ppCountKeeper := calculation.GetNewCountKeeper()
	for _, card := range cards {
		calculation.AddToCountKeeper(card, ppCountKeeper)
	}
	fmt.Println(ppCountKeeper)
	return calculation.EvaluatePPScore(ppCountKeeper).Score
}

func getSampleCardsForNoPair() []game.Card {
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

func getSampleCardsForOnePAndNoOtherPair() []game.Card {
	var cards []game.Card
	cards = append(cards, game.Card{Value: game.QUEEN, House: game.HEARTS})
	cards = append(cards, game.Card{Value: game.QUEEN, House: game.SPADES})
	cards = append(cards, game.Card{Value: game.THREE, House: game.DIAMONDS})
	cards = append(cards, game.Card{Value: game.FIVE, House: game.CLUBS})
	cards = append(cards, game.Card{Value: game.SEVEN, House: game.HEARTS})
	cards = append(cards, game.Card{Value: game.TEN, House: game.SPADES})
	cards = append(cards, game.Card{Value: game.KING, House: game.HEARTS})
	return cards
}

func getSampleCardsForTwoPAndNoOtherPair() []game.Card {
	var cards []game.Card
	cards = append(cards, game.Card{Value: game.QUEEN, House: game.HEARTS})
	cards = append(cards, game.Card{Value: game.QUEEN, House: game.SPADES})
	cards = append(cards, game.Card{Value: game.FIVE, House: game.DIAMONDS})
	cards = append(cards, game.Card{Value: game.FIVE, House: game.CLUBS})
	cards = append(cards, game.Card{Value: game.SEVEN, House: game.HEARTS})
	cards = append(cards, game.Card{Value: game.TEN, House: game.SPADES})
	cards = append(cards, game.Card{Value: game.KING, House: game.HEARTS})
	return cards
}

func getSampleCardsForThreePAndNoOtherPair() []game.Card {
	var cards []game.Card
	cards = append(cards, game.Card{Value: game.QUEEN, House: game.HEARTS})
	cards = append(cards, game.Card{Value: game.FIVE, House: game.SPADES})
	cards = append(cards, game.Card{Value: game.FIVE, House: game.DIAMONDS})
	cards = append(cards, game.Card{Value: game.FIVE, House: game.CLUBS})
	cards = append(cards, game.Card{Value: game.SEVEN, House: game.HEARTS})
	cards = append(cards, game.Card{Value: game.TEN, House: game.SPADES})
	cards = append(cards, game.Card{Value: game.KING, House: game.HEARTS})
	return cards
}

func getSampleCardsForFullHouseAndNoOtherPair() []game.Card {
	var cards []game.Card
	cards = append(cards, game.Card{Value: game.QUEEN, House: game.HEARTS})
	cards = append(cards, game.Card{Value: game.FIVE, House: game.SPADES})
	cards = append(cards, game.Card{Value: game.FIVE, House: game.DIAMONDS})
	cards = append(cards, game.Card{Value: game.FIVE, House: game.CLUBS})
	cards = append(cards, game.Card{Value: game.QUEEN, House: game.DIAMONDS})
	cards = append(cards, game.Card{Value: game.TEN, House: game.SPADES})
	cards = append(cards, game.Card{Value: game.KING, House: game.HEARTS})
	return cards
}

func getSampleCardsForFourOfAKindAndNoOtherPair() []game.Card {
	var cards []game.Card
	cards = append(cards, game.Card{Value: game.FIVE, House: game.HEARTS})
	cards = append(cards, game.Card{Value: game.FIVE, House: game.SPADES})
	cards = append(cards, game.Card{Value: game.FIVE, House: game.DIAMONDS})
	cards = append(cards, game.Card{Value: game.FIVE, House: game.CLUBS})
	cards = append(cards, game.Card{Value: game.QUEEN, House: game.DIAMONDS})
	cards = append(cards, game.Card{Value: game.TEN, House: game.SPADES})
	cards = append(cards, game.Card{Value: game.KING, House: game.HEARTS})
	return cards
}
