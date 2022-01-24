package util

import "math"

func RoundToTwoDigits(i float64) float64 {
	return math.Round(i*100) / 100
}
