package calculation

import "poker_science/internal/app/domain/model"

func GetNewHouseKeeper() map[model.House]map[model.Card]bool {
	return map[model.House]map[model.Card]bool{}
}

func AddToHouseKeeper(card model.Card, houseKeeper map[model.House]map[model.Card]bool) {
	houseKeeper[card.House][card] = true
}

func getNewFFScore() model.FFScore {
	return model.FFScore{
		Score: 1,
		Cards: map[model.Card]bool{},
	}
}

func EvaluateFFScore(houseKeeper map[model.House]map[model.Card]bool) model.FFScore {
	ffscore := getNewFFScore()

	for index := range model.Houses {
		if len(houseKeeper[model.Houses[index]]) >= 5 {
			ffscore.Score = 8
			ffscore.Cards = houseKeeper[model.Houses[index]]
			ffscore.FHouse = model.Houses[index]
			break
		}
	}

	return ffscore
}
