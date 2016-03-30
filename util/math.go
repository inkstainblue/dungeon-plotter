package util

import (
	"math"
)

func Round(x float64) int {
	_, div := math.Modf(x)

	var round float64

	switch {
	case div < 0.5:
		round = math.Floor(x)
	default:
		round = math.Ceil(x)
	}

	return int(round)
}
