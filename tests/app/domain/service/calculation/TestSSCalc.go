package calculation

import (
	"fmt"
	"poker_science/internal/app/domain/model/game"
	"poker_science/internal/app/domain/service/calculation"
)

func dummyCalculateSSScore(cards []game.Card) game.SSScore {
	ppCountKeeper := calculation.GetNewSequenceKeeper()
	for _, card := range cards {
		calculation.AddToSequenceKeeper(card, ppCountKeeper)
	}
	fmt.Println(ppCountKeeper)
	return calculation.EvaluateSSScore(ppCountKeeper)
}

func getSampleCardsForNoSequence() []game.Card {
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

func getSampleCardsForOneNoneSpecialSequence() []game.Card {
	var cards []game.Card
	cards = append(cards, game.Card{Value: game.SEVEN, House: game.HEARTS})
	cards = append(cards, game.Card{Value: game.NINE, House: game.SPADES})
	cards = append(cards, game.Card{Value: game.QUEEN, House: game.DIAMONDS})
	cards = append(cards, game.Card{Value: game.JACK, House: game.CLUBS})
	cards = append(cards, game.Card{Value: game.SEVEN, House: game.DIAMONDS})
	cards = append(cards, game.Card{Value: game.TEN, House: game.SPADES})
	cards = append(cards, game.Card{Value: game.KING, House: game.HEARTS})
	return cards
}

func getSampleCardsForTwoNonSpecialSequence() []game.Card {
	var cards []game.Card
	cards = append(cards, game.Card{Value: game.TWO, House: game.HEARTS})
	cards = append(cards, game.Card{Value: game.THREE, House: game.SPADES})
	cards = append(cards, game.Card{Value: game.SIX, House: game.DIAMONDS})
	cards = append(cards, game.Card{Value: game.FIVE, House: game.CLUBS})
	cards = append(cards, game.Card{Value: game.SEVEN, House: game.HEARTS})
	cards = append(cards, game.Card{Value: game.TEN, House: game.SPADES})
	cards = append(cards, game.Card{Value: game.FOUR, House: game.HEARTS})
	return cards
}

func getSampleCardsForThreeNonSpecialSequence() []game.Card {
	var cards []game.Card
	cards = append(cards, game.Card{Value: game.EIGHT, House: game.HEARTS})
	cards = append(cards, game.Card{Value: game.TEN, House: game.DIAMONDS})
	cards = append(cards, game.Card{Value: game.QUEEN, House: game.HEARTS})
	cards = append(cards, game.Card{Value: game.ACE, House: game.SPADES})
	cards = append(cards, game.Card{Value: game.KING, House: game.HEARTS})
	cards = append(cards, game.Card{Value: game.NINE, House: game.SPADES})
	cards = append(cards, game.Card{Value: game.JACK, House: game.CLUBS})
	return cards
}

func getSampleCardsForOneSpecialSequence() []game.Card {
	var cards []game.Card
	cards = append(cards, game.Card{Value: game.TWO, House: game.HEARTS})
	cards = append(cards, game.Card{Value: game.THREE, House: game.SPADES})
	cards = append(cards, game.Card{Value: game.FIVE, House: game.DIAMONDS})
	cards = append(cards, game.Card{Value: game.FIVE, House: game.CLUBS})
	cards = append(cards, game.Card{Value: game.QUEEN, House: game.DIAMONDS})
	cards = append(cards, game.Card{Value: game.ACE, House: game.SPADES})
	cards = append(cards, game.Card{Value: game.FOUR, House: game.HEARTS})
	return cards
}

func getSampleCardsForOneSpecialAndOneNonSpecialSequence() []game.Card {
	var cards []game.Card
	cards = append(cards, game.Card{Value: game.ACE, House: game.HEARTS})
	cards = append(cards, game.Card{Value: game.FIVE, House: game.SPADES})
	cards = append(cards, game.Card{Value: game.SIX, House: game.DIAMONDS})
	cards = append(cards, game.Card{Value: game.FOUR, House: game.CLUBS})
	cards = append(cards, game.Card{Value: game.QUEEN, House: game.DIAMONDS})
	cards = append(cards, game.Card{Value: game.TWO, House: game.SPADES})
	cards = append(cards, game.Card{Value: game.THREE, House: game.HEARTS})
	return cards
}
