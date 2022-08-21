package talib

import "math"

// Tan - Vector Trigonometric TAN
func Tan(inReal []float64) []float64 {
	outReal := make([]float64, len(inReal))
	for i := 0; i < len(inReal); i++ {
		outReal[i] = math.Tan(inReal[i])
	}
	return outReal
}
