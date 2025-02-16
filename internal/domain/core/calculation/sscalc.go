package calculation

import (
	"poker_science/internal/domain/cardinfo"
)

/*
Calculating Sequence Value as per predefined chart
No Sequence -> SSScore.Score = 1
One Sequence -> SSScore.Score = 7, SSScore.Shard = {SequenceStartCardValue, SequenceEndCardValue}{-1, -1}{-1, -1}
Two Sequences -> SSScore.Score = 7, SSScore.Shard = {SequenceStartCardValue, SequenceEndCardValue}{SequenceStartCardValue, SequenceEndCardValue}{-1, -1}
...
Why score contains Shards? Shards will help in Straight Flush and Royal Flush Check.
*/

var (
	// specialCaseCardValues ACE's index in this implementation is 12 (2 = 0, 3 = 1 and so on), Hence A, 2, 3, 4, 5 is a Shard of value {12, 3}
	specialCaseCardValues = [5]cardinfo.Value{cardinfo.ACE, cardinfo.TWO, cardinfo.THREE, cardinfo.FOUR, cardinfo.FIVE}
)

func GetNewSequenceKeeper() [13]map[cardinfo.House]bool {
	var sequenceKeeper [13]map[cardinfo.House]bool
	for _, card := range cardinfo.Cards {
		sequenceKeeper[cardinfo.GetOperationalValue(card)-2] = map[cardinfo.House]bool{}
	}
	return sequenceKeeper
}

func AddToSequenceKeeper(card cardinfo.Card, sequenceKeeper [13]map[cardinfo.House]bool) {
	sequenceKeeper[card.Value][card.House] = true
}

// EvaluateSSScore this function returns binary score (1, 7) based on existence of straight sequence along with its shard
// Shard is a lower and upper bound of straight
func EvaluateSSScore(sequenceKeeper [13]map[cardinfo.House]bool) cardinfo.SSScore {
	ssscore := cardinfo.GetNewSSScore()
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
			ssscore.Shard[shard][0] = int8(index - 4)
			ssscore.Shard[shard][1] = int8(index)
			shard++
		}
	}
	return handleSpecialCase(ssscore, shard, sequenceKeeper)
}

func handleSpecialCase(ssscore cardinfo.SSScore, shard int, sequenceKeeper [13]map[cardinfo.House]bool) cardinfo.SSScore {
	for index := range specialCaseCardValues {
		if !anyHouseExists(sequenceKeeper[specialCaseCardValues[index]]) {
			return ssscore
		}
	}
	ssscore.Score = 7
	ssscore.Shard[shard][0] = int8(cardinfo.ACE)
	ssscore.Shard[shard][1] = int8(cardinfo.FIVE)
	return ssscore
}

func anyHouseExists(houseKeeper map[cardinfo.House]bool) bool {
	for key := range houseKeeper {
		if houseKeeper[key] {
			return true
		}
	}
	return false
}
