package calculation

import (
	"fmt"
	"poker_science/internal/app/domain/model"
	"poker_science/internal/app/domain/service/calculation"
)

func dummyCalculateFFScore(cards []model.Card) model.FFScore {
	ffHouseKeeper := calculation.GetNewHouseKeeper()
	for _, card := range cards {
		calculation.AddToHouseKeeper(card, ffHouseKeeper)
	}
	fmt.Println(ffHouseKeeper)
	return calculation.EvaluateFFScore(ffHouseKeeper)
}

func getSampleCardsForNoFlush() []model.Card {
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

func getSampleCardsTieBreakingNoFlush() []model.Card {
	var cards []model.Card
	cards = append(cards, model.Card{Value: model.SEVEN, House: model.HEARTS})
	cards = append(cards, model.Card{Value: model.NINE, House: model.CLUBS})
	cards = append(cards, model.Card{Value: model.QUEEN, House: model.HEARTS})
	cards = append(cards, model.Card{Value: model.JACK, House: model.CLUBS})
	cards = append(cards, model.Card{Value: model.SEVEN, House: model.CLUBS})
	cards = append(cards, model.Card{Value: model.TEN, House: model.CLUBS})
	cards = append(cards, model.Card{Value: model.KING, House: model.HEARTS})
	return cards
}

func getSampleCardsForHardFlush() []model.Card {
	var cards []model.Card
	cards = append(cards, model.Card{Value: model.TWO, House: model.HEARTS})
	cards = append(cards, model.Card{Value: model.THREE, House: model.HEARTS})
	cards = append(cards, model.Card{Value: model.SIX, House: model.DIAMONDS})
	cards = append(cards, model.Card{Value: model.FIVE, House: model.CLUBS})
	cards = append(cards, model.Card{Value: model.SEVEN, House: model.HEARTS})
	cards = append(cards, model.Card{Value: model.TEN, House: model.HEARTS})
	cards = append(cards, model.Card{Value: model.FOUR, House: model.HEARTS})
	return cards
}

func getSampleCardsForMidFlush() []model.Card {
	var cards []model.Card
	cards = append(cards, model.Card{Value: model.EIGHT, House: model.HEARTS})
	cards = append(cards, model.Card{Value: model.TEN, House: model.DIAMONDS})
	cards = append(cards, model.Card{Value: model.QUEEN, House: model.HEARTS})
	cards = append(cards, model.Card{Value: model.ACE, House: model.HEARTS})
	cards = append(cards, model.Card{Value: model.KING, House: model.HEARTS})
	cards = append(cards, model.Card{Value: model.NINE, House: model.HEARTS})
	cards = append(cards, model.Card{Value: model.JACK, House: model.HEARTS})
	return cards
}

func getSampleCardsForPlatterFlush() []model.Card {
	var cards []model.Card
	cards = append(cards, model.Card{Value: model.TWO, House: model.HEARTS})
	cards = append(cards, model.Card{Value: model.THREE, House: model.HEARTS})
	cards = append(cards, model.Card{Value: model.FIVE, House: model.HEARTS})
	cards = append(cards, model.Card{Value: model.KING, House: model.HEARTS})
	cards = append(cards, model.Card{Value: model.QUEEN, House: model.HEARTS})
	cards = append(cards, model.Card{Value: model.SEVEN, House: model.HEARTS})
	cards = append(cards, model.Card{Value: model.FOUR, House: model.HEARTS})
	return cards
}
