package calculation

import (
	"poker_science/internal/app/domain/model"
)

/*
Calculating Sequence Value as per predefined chart
No Sequence -> SSScore.Score = 1
One Sequence -> SSScore.Score = 7, SSScore.Shard = {SequenceStartCardValue, SequenceEndCardValue}{-1, -1}{-1, -1}
*/

var (
	specialCaseCardValues = [5]model.Value{model.ACE, model.TWO, model.THREE, model.FOUR, model.FIVE}
)

func GetNewSequenceKeeper() [13]map[model.House]bool {
	var sequenceKeeper [13]map[model.House]bool
	for _, card := range model.Cards {
		sequenceKeeper[model.GetOperationalValue(card)-2] = map[model.House]bool{}
	}
	return sequenceKeeper
}

func AddToSequenceKeeper(card model.Card, sequenceKeeper [13]map[model.House]bool) {
	sequenceKeeper[model.GetOperationalValue(card.Value)-2][card.House] = true
}

func EvaluateSSScore(sequenceKeeper [13]map[model.House]bool) model.SSScore {
	ssscore := model.GetNewSSScore()
	count := 0
	shard := 0
	for index := range sequenceKeeper {
		if anyHouseExists(sequenceKeeper[index]) {
			count++
		} else {
			count = 0
		}
		if count >= 5 {
			ssscore.Score = 7
			ssscore.Shard[shard][0] = int8(index - 2)
			ssscore.Shard[shard][1] = int8(index + 2)
			shard++
		}
	}
	return handleSpecialCase(ssscore, shard, sequenceKeeper)
}

func handleSpecialCase(ssscore model.SSScore, shard int, sequenceKeeper [13]map[model.House]bool) model.SSScore {
	for index := range specialCaseCardValues {
		if !anyHouseExists(sequenceKeeper[specialCaseCardValues[index]]) {
			return ssscore
		}
	}
	ssscore.Score = 7
	ssscore.Shard[shard][0] = model.GetOperationalValue(model.ACE)
	ssscore.Shard[shard][1] = model.GetOperationalValue(model.FIVE)
	return ssscore
}

func anyHouseExists(houseKeeper map[model.House]bool) bool {
	for key := range houseKeeper {
		if houseKeeper[key] {
			return true
		}
	}
	return false
}
