package converter

import (
	"poker_science/internal/app/domain/model/game"
)

func GetCategoryByScore(score int8) game.Category {
	switch score {
	case int8(112):
		return game.RF
	case int8(56):
		return game.SF
	case int8(24):
		return game.FK
	case int8(12):
		return game.FH
	case int8(8):
		return game.FF
	case int8(7):
		return game.SS
	case int8(6):
		return game.TK
	case int8(4):
		return game.TP
	case int8(2):
		return game.PP
	}
	return game.HC
}
