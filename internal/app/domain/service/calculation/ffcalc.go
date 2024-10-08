package calculation

import (
	"poker_science/internal/app/domain/model/game"
)

func GetNewHouseKeeper() map[game.House]map[game.Card]bool {
	return map[game.House]map[game.Card]bool{}
}

func AddToHouseKeeper(card game.Card, houseKeeper map[game.House]map[game.Card]bool) {
	houseKeeper[card.House][card] = true
}

func getNewFFScore() game.FFScore {
	return game.FFScore{
		Score: 1,
		Cards: map[game.Card]bool{},
	}
}

func EvaluateFFScore(houseKeeper map[game.House]map[game.Card]bool) game.FFScore {
	ffscore := getNewFFScore()

	for index := range game.Houses {
		if len(houseKeeper[game.Houses[index]]) >= 5 {
			ffscore.Score = 8
			ffscore.Cards = houseKeeper[game.Houses[index]]
			ffscore.FHouse = game.Houses[index]
			break
		}
	}

	return ffscore
}
