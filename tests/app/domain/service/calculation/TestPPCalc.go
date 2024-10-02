package calculation

import (
	"fmt"
	"poker_science/internal/app/domain/model"
	"poker_science/internal/app/domain/service/calculation"
)

func dummyCalculatePPScore(cards []model.Card) int8 {
	ppCountKeeper := calculation.GetNewCountKeeper()
	for _, card := range cards {
		calculation.AddToCountKeeper(card, ppCountKeeper)
	}
	fmt.Println(ppCountKeeper)
	return calculation.EvaluatePPScore(ppCountKeeper)
}

func getSampleCardsForNoPair() []model.Card {
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

func getSampleCardsForOnePAndNoOtherPair() []model.Card {
	var cards []model.Card
	cards = append(cards, model.Card{Value: model.QUEEN, House: model.HEARTS})
	cards = append(cards, model.Card{Value: model.QUEEN, House: model.SPADES})
	cards = append(cards, model.Card{Value: model.THREE, House: model.DIAMONDS})
	cards = append(cards, model.Card{Value: model.FIVE, House: model.CLUBS})
	cards = append(cards, model.Card{Value: model.SEVEN, House: model.HEARTS})
	cards = append(cards, model.Card{Value: model.TEN, House: model.SPADES})
	cards = append(cards, model.Card{Value: model.KING, House: model.HEARTS})
	return cards
}

func getSampleCardsForTwoPAndNoOtherPair() []model.Card {
	var cards []model.Card
	cards = append(cards, model.Card{Value: model.QUEEN, House: model.HEARTS})
	cards = append(cards, model.Card{Value: model.QUEEN, House: model.SPADES})
	cards = append(cards, model.Card{Value: model.FIVE, House: model.DIAMONDS})
	cards = append(cards, model.Card{Value: model.FIVE, House: model.CLUBS})
	cards = append(cards, model.Card{Value: model.SEVEN, House: model.HEARTS})
	cards = append(cards, model.Card{Value: model.TEN, House: model.SPADES})
	cards = append(cards, model.Card{Value: model.KING, House: model.HEARTS})
	return cards
}

func getSampleCardsForThreePAndNoOtherPair() []model.Card {
	var cards []model.Card
	cards = append(cards, model.Card{Value: model.QUEEN, House: model.HEARTS})
	cards = append(cards, model.Card{Value: model.FIVE, House: model.SPADES})
	cards = append(cards, model.Card{Value: model.FIVE, House: model.DIAMONDS})
	cards = append(cards, model.Card{Value: model.FIVE, House: model.CLUBS})
	cards = append(cards, model.Card{Value: model.SEVEN, House: model.HEARTS})
	cards = append(cards, model.Card{Value: model.TEN, House: model.SPADES})
	cards = append(cards, model.Card{Value: model.KING, House: model.HEARTS})
	return cards
}

func getSampleCardsForFullHouseAndNoOtherPair() []model.Card {
	var cards []model.Card
	cards = append(cards, model.Card{Value: model.QUEEN, House: model.HEARTS})
	cards = append(cards, model.Card{Value: model.FIVE, House: model.SPADES})
	cards = append(cards, model.Card{Value: model.FIVE, House: model.DIAMONDS})
	cards = append(cards, model.Card{Value: model.FIVE, House: model.CLUBS})
	cards = append(cards, model.Card{Value: model.QUEEN, House: model.DIAMONDS})
	cards = append(cards, model.Card{Value: model.TEN, House: model.SPADES})
	cards = append(cards, model.Card{Value: model.KING, House: model.HEARTS})
	return cards
}

func getSampleCardsForFourOfAKindAndNoOtherPair() []model.Card {
	var cards []model.Card
	cards = append(cards, model.Card{Value: model.FIVE, House: model.HEARTS})
	cards = append(cards, model.Card{Value: model.FIVE, House: model.SPADES})
	cards = append(cards, model.Card{Value: model.FIVE, House: model.DIAMONDS})
	cards = append(cards, model.Card{Value: model.FIVE, House: model.CLUBS})
	cards = append(cards, model.Card{Value: model.QUEEN, House: model.DIAMONDS})
	cards = append(cards, model.Card{Value: model.TEN, House: model.SPADES})
	cards = append(cards, model.Card{Value: model.KING, House: model.HEARTS})
	return cards
}
