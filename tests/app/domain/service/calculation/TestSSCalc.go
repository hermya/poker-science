package calculation

import (
	"fmt"
	"poker_science/internal/app/domain/model"
	"poker_science/internal/app/domain/service/calculation"
)

func dummyCalculateSSScore(cards []model.Card) model.SSScore {
	ppCountKeeper := calculation.GetNewSequenceKeeper()
	for _, card := range cards {
		calculation.AddToSequenceKeeper(card, ppCountKeeper)
	}
	fmt.Println(ppCountKeeper)
	return calculation.EvaluateSSScore(ppCountKeeper)
}

func getSampleCardsForNoSequence() []model.Card {
	var cards []model.Card
	cards = append(cards, model.Card{Value: model.QUEEN, House: model.HEARTS})
	cards = append(cards, model.Card{Value: model.ACE, House: model.SPADES})
	cards = append(cards, model.Card{Value: model.THREE, House: model.DIAMONDS})
	cards = append(cards, model.Card{Value: model.FIVE, House: model.CLUBS})
	cards = append(cards, model.Card{Value: model.SEVEN, House: model.HEARTS})
	cards = append(cards, model.Card{Value: model.TEN, House: model.SPADES})
	cards = append(cards, model.Card{Value: model.KING, House: model.HEARTS})
	return cards
}

func getSampleCardsForOneNoneSpecialSequence() []model.Card {
	var cards []model.Card
	cards = append(cards, model.Card{Value: model.SEVEN, House: model.HEARTS})
	cards = append(cards, model.Card{Value: model.NINE, House: model.SPADES})
	cards = append(cards, model.Card{Value: model.QUEEN, House: model.DIAMONDS})
	cards = append(cards, model.Card{Value: model.JACK, House: model.CLUBS})
	cards = append(cards, model.Card{Value: model.SEVEN, House: model.DIAMONDS})
	cards = append(cards, model.Card{Value: model.TEN, House: model.SPADES})
	cards = append(cards, model.Card{Value: model.KING, House: model.HEARTS})
	return cards
}

func getSampleCardsForTwoNonSpecialSequence() []model.Card {
	var cards []model.Card
	cards = append(cards, model.Card{Value: model.TWO, House: model.HEARTS})
	cards = append(cards, model.Card{Value: model.THREE, House: model.SPADES})
	cards = append(cards, model.Card{Value: model.SIX, House: model.DIAMONDS})
	cards = append(cards, model.Card{Value: model.FIVE, House: model.CLUBS})
	cards = append(cards, model.Card{Value: model.SEVEN, House: model.HEARTS})
	cards = append(cards, model.Card{Value: model.TEN, House: model.SPADES})
	cards = append(cards, model.Card{Value: model.FOUR, House: model.HEARTS})
	return cards
}

func getSampleCardsForThreeNonSpecialSequence() []model.Card {
	var cards []model.Card
	cards = append(cards, model.Card{Value: model.EIGHT, House: model.HEARTS})
	cards = append(cards, model.Card{Value: model.TEN, House: model.DIAMONDS})
	cards = append(cards, model.Card{Value: model.QUEEN, House: model.HEARTS})
	cards = append(cards, model.Card{Value: model.ACE, House: model.SPADES})
	cards = append(cards, model.Card{Value: model.KING, House: model.HEARTS})
	cards = append(cards, model.Card{Value: model.NINE, House: model.SPADES})
	cards = append(cards, model.Card{Value: model.JACK, House: model.CLUBS})
	return cards
}

func getSampleCardsForOneSpecialSequence() []model.Card {
	var cards []model.Card
	cards = append(cards, model.Card{Value: model.TWO, House: model.HEARTS})
	cards = append(cards, model.Card{Value: model.THREE, House: model.SPADES})
	cards = append(cards, model.Card{Value: model.FIVE, House: model.DIAMONDS})
	cards = append(cards, model.Card{Value: model.FIVE, House: model.CLUBS})
	cards = append(cards, model.Card{Value: model.QUEEN, House: model.DIAMONDS})
	cards = append(cards, model.Card{Value: model.ACE, House: model.SPADES})
	cards = append(cards, model.Card{Value: model.FOUR, House: model.HEARTS})
	return cards
}

func getSampleCardsForOneSpecialAndOneNonSpecialSequence() []model.Card {
	var cards []model.Card
	cards = append(cards, model.Card{Value: model.ACE, House: model.HEARTS})
	cards = append(cards, model.Card{Value: model.FIVE, House: model.SPADES})
	cards = append(cards, model.Card{Value: model.SIX, House: model.DIAMONDS})
	cards = append(cards, model.Card{Value: model.FOUR, House: model.CLUBS})
	cards = append(cards, model.Card{Value: model.QUEEN, House: model.DIAMONDS})
	cards = append(cards, model.Card{Value: model.TWO, House: model.SPADES})
	cards = append(cards, model.Card{Value: model.THREE, House: model.HEARTS})
	return cards
}
