package calculation

import "poker_science/internal/app/domain/model"

func GetNewHouseKeeper() map[model.House][]model.Card {
	houseKeeper := map[model.House][]model.Card{}
	for index := range model.Houses {
		houseKeeper[model.Houses[index]] = []model.Card{}
	}
	return houseKeeper
}

func AddToHouseKeeper(card model.Card, houseKeeper map[model.House][]model.Card) {
	houseMap := houseKeeper[card.House]
	houseMap = append(houseMap, card)
	houseKeeper[card.House] = houseMap
}

func getNewFFScore() model.FFScore {
	return model.FFScore{
		Score: 1,
		Cards: []model.Card{},
	}
}

func EvaluateFFScore(houseKeeper map[model.House][]model.Card) model.FFScore {
	ffscore := getNewFFScore()

	for index := range model.Houses {
		if len(houseKeeper[model.Houses[index]]) >= 5 {
			ffscore.Score = 8
			ffscore.Cards = houseKeeper[model.Houses[index]]
			break
		}
	}

	return ffscore
}
