package talib

// Ema - Exponential Moving Average
func ema(inReal []float64, inTimePeriod int, k1 float64) []float64 {

	outReal := make([]float64, len(inReal))

	lookbackTotal := inTimePeriod - 1
	startIdx := lookbackTotal
	today := startIdx - lookbackTotal
	i := inTimePeriod
	tempReal := 0.0
	for i > 0 {
		tempReal += inReal[today]
		today++
		i--
	}

	prevMA := tempReal / float64(inTimePeriod)

	for today <= startIdx {
		prevMA = ((inReal[today] - prevMA) * k1) + prevMA
		today++
	}
	outReal[startIdx] = prevMA
	outIdx := startIdx + 1
	for today < len(inReal) {
		prevMA = ((inReal[today] - prevMA) * k1) + prevMA
		outReal[outIdx] = prevMA
		today++
		outIdx++
	}

	return outReal
}

// Ema - Exponential Moving Average
func Ema(inReal []float64, inTimePeriod int) []float64 {

	k := 2.0 / float64(inTimePeriod+1)
	outReal := ema(inReal, inTimePeriod, k)
	return outReal
}
