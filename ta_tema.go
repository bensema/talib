package talib

// Tema - Triple Exponential Moving Average
func Tema(inReal []float64, inTimePeriod int) []float64 {

	outReal := make([]float64, len(inReal))
	firstEMA := Ema(inReal, inTimePeriod)
	secondEMA := Ema(firstEMA[inTimePeriod-1:], inTimePeriod)
	thirdEMA := Ema(secondEMA[inTimePeriod-1:], inTimePeriod)

	outIdx := (inTimePeriod * 3) - 3
	secondEMAIdx := (inTimePeriod * 2) - 2
	thirdEMAIdx := inTimePeriod - 1

	for outIdx < len(inReal) {
		outReal[outIdx] = thirdEMA[thirdEMAIdx] + ((3.0 * firstEMA[outIdx]) - (3.0 * secondEMA[secondEMAIdx]))
		outIdx++
		secondEMAIdx++
		thirdEMAIdx++
	}

	return outReal
}
