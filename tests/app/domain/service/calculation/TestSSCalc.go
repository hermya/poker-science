package calculation

import (
	"fmt"
	"poker_science/internal/domain/cardinfo"
	"poker_science/internal/domain/core/calculation"
)

func dummyCalculateSSScore(cards []cardinfo.Card) cardinfo.SSScore {
	ppCountKeeper := calculation.GetNewSequenceKeeper()
	for _, card := range cards {
		calculation.AddToSequenceKeeper(card, ppCountKeeper)
	}
	fmt.Println(ppCountKeeper)
	return calculation.EvaluateSSScore(ppCountKeeper)
}

func getSampleCardsForNoSequence() []cardinfo.Card {
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

func getSampleCardsForOneNoneSpecialSequence() []cardinfo.Card {
	var cards []cardinfo.Card
	cards = append(cards, cardinfo.Card{Value: cardinfo.SEVEN, House: cardinfo.HEARTS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.NINE, House: cardinfo.SPADES})
	cards = append(cards, cardinfo.Card{Value: cardinfo.QUEEN, House: cardinfo.DIAMONDS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.JACK, House: cardinfo.CLUBS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.SEVEN, House: cardinfo.DIAMONDS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.TEN, House: cardinfo.SPADES})
	cards = append(cards, cardinfo.Card{Value: cardinfo.KING, House: cardinfo.HEARTS})
	return cards
}

func getSampleCardsForTwoNonSpecialSequence() []cardinfo.Card {
	var cards []cardinfo.Card
	cards = append(cards, cardinfo.Card{Value: cardinfo.TWO, House: cardinfo.HEARTS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.THREE, House: cardinfo.SPADES})
	cards = append(cards, cardinfo.Card{Value: cardinfo.SIX, House: cardinfo.DIAMONDS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.FIVE, House: cardinfo.CLUBS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.SEVEN, House: cardinfo.HEARTS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.TEN, House: cardinfo.SPADES})
	cards = append(cards, cardinfo.Card{Value: cardinfo.FOUR, House: cardinfo.HEARTS})
	return cards
}

func getSampleCardsForThreeNonSpecialSequence() []cardinfo.Card {
	var cards []cardinfo.Card
	cards = append(cards, cardinfo.Card{Value: cardinfo.EIGHT, House: cardinfo.HEARTS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.TEN, House: cardinfo.DIAMONDS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.QUEEN, House: cardinfo.HEARTS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.ACE, House: cardinfo.SPADES})
	cards = append(cards, cardinfo.Card{Value: cardinfo.KING, House: cardinfo.HEARTS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.NINE, House: cardinfo.SPADES})
	cards = append(cards, cardinfo.Card{Value: cardinfo.JACK, House: cardinfo.CLUBS})
	return cards
}

func getSampleCardsForOneSpecialSequence() []cardinfo.Card {
	var cards []cardinfo.Card
	cards = append(cards, cardinfo.Card{Value: cardinfo.TWO, House: cardinfo.HEARTS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.THREE, House: cardinfo.SPADES})
	cards = append(cards, cardinfo.Card{Value: cardinfo.FIVE, House: cardinfo.DIAMONDS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.FIVE, House: cardinfo.CLUBS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.QUEEN, House: cardinfo.DIAMONDS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.ACE, House: cardinfo.SPADES})
	cards = append(cards, cardinfo.Card{Value: cardinfo.FOUR, House: cardinfo.HEARTS})
	return cards
}

func getSampleCardsForOneSpecialAndOneNonSpecialSequence() []cardinfo.Card {
	var cards []cardinfo.Card
	cards = append(cards, cardinfo.Card{Value: cardinfo.ACE, House: cardinfo.HEARTS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.FIVE, House: cardinfo.SPADES})
	cards = append(cards, cardinfo.Card{Value: cardinfo.SIX, House: cardinfo.DIAMONDS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.FOUR, House: cardinfo.CLUBS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.QUEEN, House: cardinfo.DIAMONDS})
	cards = append(cards, cardinfo.Card{Value: cardinfo.TWO, House: cardinfo.SPADES})
	cards = append(cards, cardinfo.Card{Value: cardinfo.THREE, House: cardinfo.HEARTS})
	return cards
}
