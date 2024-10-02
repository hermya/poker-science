package passes

import "poker_science/internal/app/domain/model"

func GetWinCategory() {

}

func finalScoreAggregator() {

}

func combineFFAndSS(ffscore model.FFScore, ssscore model.SSScore) int8 {
	if ffscore.Score == 1 {
		return ssscore.Score
	} else if ssscore.Score == 1 {
		return ffscore.Score
	}
	score := int8(8)
	key := model.Card{House: ffscore.FHouse}
	for shard := 0; shard < 3; shard++ {
		if ssscore.Shard[shard][0] == 14 {
			// check ace exists, then 2 to 5 exists
			key.Value = model.ACE
			if ffscore.Cards[key] {
				if isStraightFlush(ffscore, model.TWO, model.FIVE) {
					return ffscore.Score * ssscore.Score
				}
			}
		} else if ssscore.Shard[shard][0] != -1 {
			if isStraightFlush(ffscore, model.Value(ssscore.Shard[shard][0]), model.Value(ssscore.Shard[shard][1])) {
				return ffscore.Score * ssscore.Score
			}
		}
	}
	return score
}

func isStraightFlush(ffscore model.FFScore, start model.Value, end model.Value) bool {
	key := model.Card{House: ffscore.FHouse}
	for tempCard := start; tempCard <= end; tempCard++ {
		key.Value = tempCard
		if !ffscore.Cards[key] {
			return false
		}
	}
	return true
}
