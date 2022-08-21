package talib

// Mult - Vector arithmetic multiply
func Mult(inReal0 []float64, inReal1 []float64) []float64 {
	outReal := make([]float64, len(inReal0))
	for i := 0; i < len(inReal0); i++ {
		outReal[i] = inReal0[i] * inReal1[i]
	}
	return outReal
}
