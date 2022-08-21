package talib

import "math"

// Ceil - Vector CEIL
func Ceil(inReal []float64) []float64 {
	outReal := make([]float64, len(inReal))
	for i := 0; i < len(inReal); i++ {
		outReal[i] = math.Ceil(inReal[i])
	}
	return outReal
}
