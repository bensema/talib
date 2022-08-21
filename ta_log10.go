package talib

import "math"

// Log10 - Vector LOG10
func Log10(inReal []float64) []float64 {
	outReal := make([]float64, len(inReal))
	for i := 0; i < len(inReal); i++ {
		outReal[i] = math.Log10(inReal[i])
	}
	return outReal
}
