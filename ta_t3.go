package talib

// T3 - Triple Exponential Moving Average (T3) (lookback=6*inTimePeriod)
func T3(inReal []float64, inTimePeriod int, inVFactor float64) []float64 {

	outReal := make([]float64, len(inReal))

	lookbackTotal := 6 * (inTimePeriod - 1)
	startIdx := lookbackTotal
	today := startIdx - lookbackTotal
	k := 2.0 / (float64(inTimePeriod) + 1.0)
	oneMinusK := 1.0 - k
	tempReal := inReal[today]
	today++
	for i := inTimePeriod - 1; i > 0; i-- {
		tempReal += inReal[today]
		today++
	}
	e1 := tempReal / float64(inTimePeriod)
	tempReal = e1
	for i := inTimePeriod - 1; i > 0; i-- {
		e1 = (k * inReal[today]) + (oneMinusK * e1)
		tempReal += e1
		today++
	}
	e2 := tempReal / float64(inTimePeriod)
	tempReal = e2
	for i := inTimePeriod - 1; i > 0; i-- {
		e1 = (k * inReal[today]) + (oneMinusK * e1)
		e2 = (k * e1) + (oneMinusK * e2)
		tempReal += e2
		today++
	}
	e3 := tempReal / float64(inTimePeriod)
	tempReal = e3
	for i := inTimePeriod - 1; i > 0; i-- {
		e1 = (k * inReal[today]) + (oneMinusK * e1)
		e2 = (k * e1) + (oneMinusK * e2)
		e3 = (k * e2) + (oneMinusK * e3)
		tempReal += e3
		today++
	}
	e4 := tempReal / float64(inTimePeriod)
	tempReal = e4
	for i := inTimePeriod - 1; i > 0; i-- {
		e1 = (k * inReal[today]) + (oneMinusK * e1)
		e2 = (k * e1) + (oneMinusK * e2)
		e3 = (k * e2) + (oneMinusK * e3)
		e4 = (k * e3) + (oneMinusK * e4)
		tempReal += e4
		today++
	}
	e5 := tempReal / float64(inTimePeriod)
	tempReal = e5
	for i := inTimePeriod - 1; i > 0; i-- {
		e1 = (k * inReal[today]) + (oneMinusK * e1)
		e2 = (k * e1) + (oneMinusK * e2)
		e3 = (k * e2) + (oneMinusK * e3)
		e4 = (k * e3) + (oneMinusK * e4)
		e5 = (k * e4) + (oneMinusK * e5)
		tempReal += e5
		today++
	}
	e6 := tempReal / float64(inTimePeriod)
	for today <= startIdx {
		e1 = (k * inReal[today]) + (oneMinusK * e1)
		e2 = (k * e1) + (oneMinusK * e2)
		e3 = (k * e2) + (oneMinusK * e3)
		e4 = (k * e3) + (oneMinusK * e4)
		e5 = (k * e4) + (oneMinusK * e5)
		e6 = (k * e5) + (oneMinusK * e6)
		today++
	}
	tempReal = inVFactor * inVFactor
	c1 := -(tempReal * inVFactor)
	c2 := 3.0 * (tempReal - c1)
	c3 := -6.0*tempReal - 3.0*(inVFactor-c1)
	c4 := 1.0 + 3.0*inVFactor - c1 + 3.0*tempReal
	outIdx := lookbackTotal
	outReal[outIdx] = c1*e6 + c2*e5 + c3*e4 + c4*e3
	outIdx++
	for today < len(inReal) {
		e1 = (k * inReal[today]) + (oneMinusK * e1)
		e2 = (k * e1) + (oneMinusK * e2)
		e3 = (k * e2) + (oneMinusK * e3)
		e4 = (k * e3) + (oneMinusK * e4)
		e5 = (k * e4) + (oneMinusK * e5)
		e6 = (k * e5) + (oneMinusK * e6)
		outReal[outIdx] = c1*e6 + c2*e5 + c3*e4 + c4*e3
		outIdx++
		today++
	}

	return outReal
}
