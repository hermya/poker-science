package passes

import (
	"poker_science/internal/app/domain/model/game"
	"poker_science/internal/app/domain/service/converter"
)

func GetWinCategory(ppscore game.PPScore, ffscore game.FFScore, ssscore game.SSScore) game.Category {
	finalScore := ppscore.Score
	finalScore = max(finalScore, combineFFAndSS(ffscore, ssscore))
	return converter.GetCategoryByScore(finalScore)
}

func max(a int8, b int8) int8 {
	if a > b {
		return a
	}
	return b
}

func combineFFAndSS(ffscore game.FFScore, ssscore game.SSScore) int8 {
	if ffscore.Score == 1 {
		return ssscore.Score
	} else if ssscore.Score == 1 {
		return ffscore.Score
	}
	card := game.Card{House: ffscore.FHouse}
	for shard := 0; shard < 3; shard++ {
		if ssscore.Shard[shard][0] == int8(game.ACE) {
			// check ace exists, then 2 to 5 exists
			card.Value = game.ACE
			if ffscore.Cards[card] && isStraightFlush(ffscore, game.TWO, game.FIVE) {
				return ffscore.Score * ssscore.Score
			}
		} else if ssscore.Shard[shard][0] != -1 {
			if isStraightFlush(ffscore, game.Value(ssscore.Shard[shard][0]), game.Value(ssscore.Shard[shard][1])) {
				return isRoyalFlush(ffscore) * ffscore.Score * ssscore.Score
			}
		}
	}
	return int8(8)
}

func isStraightFlush(ffscore game.FFScore, start game.Value, end game.Value) bool {
	key := game.Card{House: ffscore.FHouse}
	for tempCard := start; tempCard <= end; tempCard++ {
		key.Value = tempCard
		if !ffscore.Cards[key] {
			return false
		}
	}
	return true
}

func isRoyalFlush(ffscore game.FFScore) int8 {
	if ffscore.Cards[game.Card{Value: game.ACE, House: ffscore.FHouse}] {
		return 2
	}
	return 1
}
