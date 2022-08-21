package talib

import "math"

// Exp - Vector atrithmetic EXP
func Exp(inReal []float64) []float64 {
	outReal := make([]float64, len(inReal))
	for i := 0; i < len(inReal); i++ {
		outReal[i] = math.Exp(inReal[i])
	}
	return outReal
}
