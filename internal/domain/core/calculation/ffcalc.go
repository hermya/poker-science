package calculation

import (
	"poker_science/internal/domain/cardinfo"
)

func GetNewHouseKeeper() map[cardinfo.House]map[cardinfo.Card]bool {
	return map[cardinfo.House]map[cardinfo.Card]bool{}
}

func AddToHouseKeeper(card cardinfo.Card, houseKeeper map[cardinfo.House]map[cardinfo.Card]bool) {
	houseKeeper[card.House][card] = true
}

func EvaluateFFScore(houseKeeper map[cardinfo.House]map[cardinfo.Card]bool) cardinfo.FFScore {
	ffscore := cardinfo.GetNewFFScore()

	for index := range cardinfo.Houses {
		if len(houseKeeper[cardinfo.Houses[index]]) >= 5 {
			ffscore.Score = 8
			ffscore.Cards = houseKeeper[cardinfo.Houses[index]]
			ffscore.FHouse = cardinfo.Houses[index]
			break
		}
	}

	return ffscore
}
