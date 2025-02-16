package calculation

import (
	"poker_science/internal/domain/cardinfo"
)

/*
Calculating Pair values as per pre-defined chart
No Pair -> 1! = 1
One Pair -> 2! = 2
Two Pair -> 2! * 2! = 4
Three of a Kind = 3! = 6
Full House = Two Pair && Three Pair = 2! * 3! = 12
Four of a Kind = 4! = 24
*/

func fact(n int8) int8 {
	if n == 1 {
		return 1
	}
	return n * fact(n-1)
}

func GetNewCountKeeper() map[cardinfo.Value]int8 {
	countKeeper := map[cardinfo.Value]int8{}
	for _, card := range cardinfo.Cards {
		countKeeper[card] = 0
	}
	return countKeeper
}

func AddToCountKeeper(card cardinfo.Card, countKeeper map[cardinfo.Value]int8) {
	countKeeper[card.Value]++
}

// EvaluatePPScore This function returns score for Pair hands. This includes P, 2P, 3P, 3P2P, and 4P
func EvaluatePPScore(countKeeper map[cardinfo.Value]int8) cardinfo.PPScore {
	pairs := map[int8]int8{}
	ppscore := cardinfo.GetNewPPScore()
	for value, count := range countKeeper {
		if count > 1 {
			pairs[count]++
			ppscore.CardValues[value] = count
		}
	}

	if pairs[4] != 0 {
		ppscore.Score = fact(4)
		return ppscore
	}

	if pairs[3] != 0 {
		ppscore.Score = fact(3)
		if pairs[2] != 0 {
			ppscore.Score *= fact(2)
			return ppscore
		}
		return ppscore
	}

	if pairs[2] != 0 {
		ppscore.Score = fact(2)
		if pairs[2] > 1 {
			ppscore.Score *= fact(2)
			return ppscore
		}
		return ppscore
	}

	return ppscore
}
