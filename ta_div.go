package talib

// Div - Vector arithmetic division
// 列表1 与 列表2 每个元素互除
func Div(inReal0 []float64, inReal1 []float64) []float64 {
	outReal := make([]float64, len(inReal0))
	for i := 0; i < len(inReal0); i++ {
		outReal[i] = inReal0[i] / inReal1[i]
	}
	return outReal
}
