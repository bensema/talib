package talib

// Ema - Exponential Moving Average
func Ema(inReal []float64, inTimePeriod int) []float64 {

	k := 2.0 / float64(inTimePeriod+1)
	outReal := ema(inReal, inTimePeriod, k)
	return outReal
}
