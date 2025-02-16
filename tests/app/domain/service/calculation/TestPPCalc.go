package calculation

import (
	"fmt"
	"poker_science/internal/domain/cardinfo"
	"poker_science/internal/domain/core/calculation"
)

func dummyCalculatePPScore(cards []cardinfo.Card) int8 {
	ppCountKeeper := calculation.GetNewCountKeeper()
	for _, card := range cards {
		calculation.AddToCountKeeper(card, ppCountKeeper)
	}
	fmt.Println(ppCountKeeper)
	return calculation.EvaluatePPScore(ppCountKeeper).Score
}

func getSampleCardsForNoPair() []cardinfo.Card {
	var cards []cardinfo.Card
	cards = append(cards, cardinfo.Card{Value: cardinfo.QUEEN, House: cardinfo.HEARTS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.ACE, House: cardinfo.SPADES})
	cards = append(cards, cardinfo.Card{Value: cardinfo.THREE, House: cardinfo.DIAMONDS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.FIVE, House: cardinfo.CLUBS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.SEVEN, House: cardinfo.HEARTS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.TEN, House: cardinfo.SPADES})
	cards = append(cards, cardinfo.Card{Value: cardinfo.KING, House: cardinfo.HEARTS})
	return cards
}

func getSampleCardsForOnePAndNoOtherPair() []cardinfo.Card {
	var cards []cardinfo.Card
	cards = append(cards, cardinfo.Card{Value: cardinfo.QUEEN, House: cardinfo.HEARTS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.QUEEN, House: cardinfo.SPADES})
	cards = append(cards, cardinfo.Card{Value: cardinfo.THREE, House: cardinfo.DIAMONDS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.FIVE, House: cardinfo.CLUBS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.SEVEN, House: cardinfo.HEARTS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.TEN, House: cardinfo.SPADES})
	cards = append(cards, cardinfo.Card{Value: cardinfo.KING, House: cardinfo.HEARTS})
	return cards
}

func getSampleCardsForTwoPAndNoOtherPair() []cardinfo.Card {
	var cards []cardinfo.Card
	cards = append(cards, cardinfo.Card{Value: cardinfo.QUEEN, House: cardinfo.HEARTS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.QUEEN, House: cardinfo.SPADES})
	cards = append(cards, cardinfo.Card{Value: cardinfo.FIVE, House: cardinfo.DIAMONDS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.FIVE, House: cardinfo.CLUBS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.SEVEN, House: cardinfo.HEARTS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.TEN, House: cardinfo.SPADES})
	cards = append(cards, cardinfo.Card{Value: cardinfo.KING, House: cardinfo.HEARTS})
	return cards
}

func getSampleCardsForThreePAndNoOtherPair() []cardinfo.Card {
	var cards []cardinfo.Card
	cards = append(cards, cardinfo.Card{Value: cardinfo.QUEEN, House: cardinfo.HEARTS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.FIVE, House: cardinfo.SPADES})
	cards = append(cards, cardinfo.Card{Value: cardinfo.FIVE, House: cardinfo.DIAMONDS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.FIVE, House: cardinfo.CLUBS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.SEVEN, House: cardinfo.HEARTS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.TEN, House: cardinfo.SPADES})
	cards = append(cards, cardinfo.Card{Value: cardinfo.KING, House: cardinfo.HEARTS})
	return cards
}

func getSampleCardsForFullHouseAndNoOtherPair() []cardinfo.Card {
	var cards []cardinfo.Card
	cards = append(cards, cardinfo.Card{Value: cardinfo.QUEEN, House: cardinfo.HEARTS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.FIVE, House: cardinfo.SPADES})
	cards = append(cards, cardinfo.Card{Value: cardinfo.FIVE, House: cardinfo.DIAMONDS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.FIVE, House: cardinfo.CLUBS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.QUEEN, House: cardinfo.DIAMONDS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.TEN, House: cardinfo.SPADES})
	cards = append(cards, cardinfo.Card{Value: cardinfo.KING, House: cardinfo.HEARTS})
	return cards
}

func getSampleCardsForFourOfAKindAndNoOtherPair() []cardinfo.Card {
	var cards []cardinfo.Card
	cards = append(cards, cardinfo.Card{Value: cardinfo.FIVE, House: cardinfo.HEARTS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.FIVE, House: cardinfo.SPADES})
	cards = append(cards, cardinfo.Card{Value: cardinfo.FIVE, House: cardinfo.DIAMONDS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.FIVE, House: cardinfo.CLUBS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.QUEEN, House: cardinfo.DIAMONDS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.TEN, House: cardinfo.SPADES})
	cards = append(cards, cardinfo.Card{Value: cardinfo.KING, House: cardinfo.HEARTS})
	return cards
}
