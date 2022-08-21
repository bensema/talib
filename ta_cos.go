package talib

import "math"

// Cos - Vector Trigonometric COS
func Cos(inReal []float64) []float64 {
	outReal := make([]float64, len(inReal))
	for i := 0; i < len(inReal); i++ {
		outReal[i] = math.Cos(inReal[i])
	}
	return outReal
}
