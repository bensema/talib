package talib

func sum(arr []float64) float64 {
	var result float64
	for _, v := range arr {
		result += v
	}
	return result
}

func ema(inReal []float64, inTimePeriod int, k1 float64) []float64 {
	outReal := make([]float64, len(inReal))

	prevMA := sum(inReal[:inTimePeriod]) / float64(inTimePeriod)
	outReal[inTimePeriod-1] = prevMA

	for i := inTimePeriod; i < len(inReal); i++ {
		prevMA = ((inReal[i] - prevMA) * k1) + prevMA
		outReal[i] = prevMA
	}

	return outReal
}
