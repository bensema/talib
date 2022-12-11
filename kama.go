package talib

import "math"

// Kama - Kaufman Adaptive Moving Average
func Kama(inReal []float64, inTimePeriod int) []float64 {

	outReal := make([]float64, len(inReal))

	constMax := 2.0 / (30.0 + 1.0)
	constDiff := 2.0/(2.0+1.0) - constMax
	lookbackTotal := inTimePeriod
	startIdx := lookbackTotal
	sumROC1 := 0.0
	today := startIdx - lookbackTotal
	trailingIdx := today
	i := inTimePeriod
	for i > 0 {
		tempReal := inReal[today]
		today++
		tempReal -= inReal[today]
		sumROC1 += math.Abs(tempReal)
		i--
	}
	prevKAMA := inReal[today-1]
	tempReal := inReal[today]
	tempReal2 := inReal[trailingIdx]
	trailingIdx++
	periodROC := tempReal - tempReal2
	trailingValue := tempReal2
	if (sumROC1 <= periodROC) || (((-(0.00000000000001)) < sumROC1) && (sumROC1 < (0.00000000000001))) {
		tempReal = 1.0
	} else {
		tempReal = math.Abs(periodROC / sumROC1)
	}
	tempReal = (tempReal * constDiff) + constMax
	tempReal *= tempReal
	prevKAMA = ((inReal[today] - prevKAMA) * tempReal) + prevKAMA
	today++
	for today <= startIdx {
		tempReal = inReal[today]
		tempReal2 = inReal[trailingIdx]
		trailingIdx++
		periodROC = tempReal - tempReal2
		sumROC1 -= math.Abs(trailingValue - tempReal2)
		sumROC1 += math.Abs(tempReal - inReal[today-1])
		trailingValue = tempReal2
		if (sumROC1 <= periodROC) || (((-(0.00000000000001)) < sumROC1) && (sumROC1 < (0.00000000000001))) {
			tempReal = 1.0
		} else {
			tempReal = math.Abs(periodROC / sumROC1)
		}
		tempReal = (tempReal * constDiff) + constMax
		tempReal *= tempReal
		prevKAMA = ((inReal[today] - prevKAMA) * tempReal) + prevKAMA
		today++
	}
	outReal[inTimePeriod] = prevKAMA
	outIdx := inTimePeriod + 1
	for today < len(inReal) {
		tempReal = inReal[today]
		tempReal2 = inReal[trailingIdx]
		trailingIdx++
		periodROC = tempReal - tempReal2
		sumROC1 -= math.Abs(trailingValue - tempReal2)
		sumROC1 += math.Abs(tempReal - inReal[today-1])
		trailingValue = tempReal2
		if (sumROC1 <= periodROC) || (((-(0.00000000000001)) < sumROC1) && (sumROC1 < (0.00000000000001))) {
			tempReal = 1.0
		} else {
			tempReal = math.Abs(periodROC / sumROC1)
		}
		tempReal = (tempReal * constDiff) + constMax
		tempReal *= tempReal
		prevKAMA = ((inReal[today] - prevKAMA) * tempReal) + prevKAMA
		today++
		outReal[outIdx] = prevKAMA
		outIdx++
	}

	return outReal
}
