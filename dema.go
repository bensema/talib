package talib

// Dema - Double Exponential Moving Average
func Dema(inReal []float64, inTimePeriod int) []float64 {

	outReal := make([]float64, len(inReal))
	firstEMA := Ema(inReal, inTimePeriod)
	secondEMA := Ema(firstEMA[inTimePeriod-1:], inTimePeriod)

	for outIdx, secondEMAIdx := (inTimePeriod*2)-2, inTimePeriod-1; outIdx < len(inReal); outIdx, secondEMAIdx = outIdx+1, secondEMAIdx+1 {
		outReal[outIdx] = (2.0 * firstEMA[outIdx]) - secondEMA[secondEMAIdx]
	}

	return outReal
}
