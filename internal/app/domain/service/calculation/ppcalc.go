package calculation

import (
	"poker_science/internal/app/domain/model"
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

func GetNewCountKeeper() map[model.Value]int8 {
	countKeeper := map[model.Value]int8{}
	for _, card := range model.Cards {
		countKeeper[card] = 0
	}
	return countKeeper
}

func AddToCountKeeper(card model.Card, countKeeper map[model.Value]int8) {
	countKeeper[card.Value]++
}

// EvaluatePPScore This function returns score for Pair hands. This includes P, 2P, 3P, 3P2P, and 4P
func EvaluatePPScore(countKeeper map[model.Value]int8) int8 {
	pairs := map[int8]int8{}

	for _, count := range countKeeper {
		if count > 1 {
			pairs[count]++
		}
	}

	if pairs[4] != 0 {
		return fact(4)
	}

	if pairs[3] != 0 {
		temp := fact(3)
		if pairs[2] != 0 {
			return temp * fact(2)
		}
		return temp
	}

	if pairs[2] != 0 {
		temp := fact(2)
		if pairs[2] > 1 {
			return temp * 2
		}
		return temp
	}

	return 1
}
