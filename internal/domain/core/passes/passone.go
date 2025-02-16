package passes

import (
	"poker_science/internal/domain/cardinfo"
	"poker_science/internal/domain/core/converter"
)

func GetWinCategory(ppscore cardinfo.PPScore, ffscore cardinfo.FFScore, ssscore cardinfo.SSScore) cardinfo.Category {
	finalScore := ppscore.Score
	finalScore = max(finalScore, combineFFAndSS(ffscore, ssscore))
	return converter.GetCategoryByScore(finalScore)
}

func combineFFAndSS(ffscore cardinfo.FFScore, ssscore cardinfo.SSScore) int8 {
	if ffscore.Score == 1 {
		return ssscore.Score
	} else if ssscore.Score == 1 {
		return ffscore.Score
	}
	card := cardinfo.Card{House: ffscore.FHouse}
	for shard := 0; shard < 3; shard++ {
		if ssscore.Shard[shard][0] == int8(cardinfo.ACE) {
			// check ace exists, then 2 to 5 exists
			card.Value = cardinfo.ACE
			if ffscore.Cards[card] && isStraightFlush(ffscore, cardinfo.TWO, cardinfo.FIVE) {
				return ffscore.Score * ssscore.Score
			}
		} else if ssscore.Shard[shard][0] != -1 {
			if isStraightFlush(ffscore, cardinfo.Value(ssscore.Shard[shard][0]), cardinfo.Value(ssscore.Shard[shard][1])) {
				return isRoyalFlush(ffscore) * ffscore.Score * ssscore.Score
			}
		}
	}
	return int8(8)
}

func isStraightFlush(ffscore cardinfo.FFScore, start cardinfo.Value, end cardinfo.Value) bool {
	key := cardinfo.Card{House: ffscore.FHouse}
	for tempCard := start; tempCard <= end; tempCard++ {
		key.Value = tempCard
		if !ffscore.Cards[key] {
			return false
		}
	}
	return true
}

func isRoyalFlush(ffscore cardinfo.FFScore) int8 {
	if ffscore.Cards[cardinfo.Card{Value: cardinfo.ACE, House: ffscore.FHouse}] {
		return 2
	}
	return 1
}
