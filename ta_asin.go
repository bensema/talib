package talib

import "math"

// Asin - Vector Trigonometric ASIN
func Asin(inReal []float64) []float64 {
	outReal := make([]float64, len(inReal))
	for i := 0; i < len(inReal); i++ {
		outReal[i] = math.Asin(inReal[i])
	}
	return outReal
}
