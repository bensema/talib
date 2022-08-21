package talib

import "math"

// Sin - Vector Trigonometric SIN
func Sin(inReal []float64) []float64 {
	outReal := make([]float64, len(inReal))
	for i := 0; i < len(inReal); i++ {
		outReal[i] = math.Sin(inReal[i])
	}
	return outReal
}
