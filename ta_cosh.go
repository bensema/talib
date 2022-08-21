package talib

import "math"

// Cosh - Vector Trigonometric COSH
func Cosh(inReal []float64) []float64 {
	outReal := make([]float64, len(inReal))
	for i := 0; i < len(inReal); i++ {
		outReal[i] = math.Cosh(inReal[i])
	}
	return outReal
}
