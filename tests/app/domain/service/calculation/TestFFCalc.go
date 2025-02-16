package calculation

import (
	"fmt"
	"poker_science/internal/domain/cardinfo"
	"poker_science/internal/domain/core/calculation"
)

func dummyCalculateFFScore(cards []cardinfo.Card) cardinfo.FFScore {
	ffHouseKeeper := calculation.GetNewHouseKeeper()
	for _, card := range cards {
		calculation.AddToHouseKeeper(card, ffHouseKeeper)
	}
	fmt.Println(ffHouseKeeper)
	return calculation.EvaluateFFScore(ffHouseKeeper)
}

func getSampleCardsForNoFlush() []cardinfo.Card {
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

func getSampleCardsTieBreakingNoFlush() []cardinfo.Card {
	var cards []cardinfo.Card
	cards = append(cards, cardinfo.Card{Value: cardinfo.SEVEN, House: cardinfo.HEARTS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.NINE, House: cardinfo.CLUBS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.QUEEN, House: cardinfo.HEARTS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.JACK, House: cardinfo.CLUBS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.SEVEN, House: cardinfo.CLUBS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.TEN, House: cardinfo.CLUBS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.KING, House: cardinfo.HEARTS})
	return cards
}

func getSampleCardsForHardFlush() []cardinfo.Card {
	var cards []cardinfo.Card
	cards = append(cards, cardinfo.Card{Value: cardinfo.TWO, House: cardinfo.HEARTS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.THREE, House: cardinfo.HEARTS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.SIX, House: cardinfo.DIAMONDS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.FIVE, House: cardinfo.CLUBS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.SEVEN, House: cardinfo.HEARTS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.TEN, House: cardinfo.HEARTS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.FOUR, House: cardinfo.HEARTS})
	return cards
}

func getSampleCardsForMidFlush() []cardinfo.Card {
	var cards []cardinfo.Card
	cards = append(cards, cardinfo.Card{Value: cardinfo.EIGHT, House: cardinfo.HEARTS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.TEN, House: cardinfo.DIAMONDS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.QUEEN, House: cardinfo.HEARTS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.ACE, House: cardinfo.HEARTS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.KING, House: cardinfo.HEARTS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.NINE, House: cardinfo.HEARTS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.JACK, House: cardinfo.HEARTS})
	return cards
}

func getSampleCardsForPlatterFlush() []cardinfo.Card {
	var cards []cardinfo.Card
	cards = append(cards, cardinfo.Card{Value: cardinfo.TWO, House: cardinfo.HEARTS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.THREE, House: cardinfo.HEARTS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.FIVE, House: cardinfo.HEARTS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.KING, House: cardinfo.HEARTS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.QUEEN, House: cardinfo.HEARTS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.SEVEN, House: cardinfo.HEARTS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.FOUR, House: cardinfo.HEARTS})
	return cards
}
