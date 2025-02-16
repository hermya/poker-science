package converter

import (
	"poker_science/internal/domain/cardinfo"
)

func GetCategoryByScore(score int8) cardinfo.Category {
	switch score {
	case int8(112):
		return cardinfo.RF
	case int8(56):
		return cardinfo.SF
	case int8(24):
		return cardinfo.FK
	case int8(12):
		return cardinfo.FH
	case int8(8):
		return cardinfo.FF
	case int8(7):
		return cardinfo.SS
	case int8(6):
		return cardinfo.TK
	case int8(4):
		return cardinfo.TP
	case int8(2):
		return cardinfo.PP
	}
	return cardinfo.HC
}
