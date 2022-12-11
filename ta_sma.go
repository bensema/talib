package talib

// Sma - Simple Moving Average

func Sma(inReal []float64, inTimePeriod int) []float64 {
	outReal := make([]float64, len(inReal))

	lookbackTotal := inTimePeriod - 1
	startIdx := lookbackTotal
	trailingIdx := startIdx - lookbackTotal
	periodTotal := sum(inReal[trailingIdx:startIdx])

	for i := startIdx; i < len(inReal); i++ {
		periodTotal += inReal[i]
		tempReal := periodTotal
		periodTotal -= inReal[trailingIdx]
		outReal[i] = tempReal / float64(inTimePeriod)
		trailingIdx++
	}

	return outReal
}
