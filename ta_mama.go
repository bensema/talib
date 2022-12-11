package talib

import "math"

// Mama - MESA Adaptive Moving Average (lookback=32)
func Mama(inReal []float64, inFastLimit float64, inSlowLimit float64) ([]float64, []float64) {

	outMAMA := make([]float64, len(inReal))
	outFAMA := make([]float64, len(inReal))

	a := 0.0962
	b := 0.5769
	detrenderOdd := make([]float64, 3)
	detrenderEven := make([]float64, 3)
	q1Odd := make([]float64, 3)
	q1Even := make([]float64, 3)
	jIOdd := make([]float64, 3)
	jIEven := make([]float64, 3)
	jQOdd := make([]float64, 3)
	jQEven := make([]float64, 3)
	rad2Deg := 180.0 / (4.0 * math.Atan(1))
	lookbackTotal := 32
	startIdx := lookbackTotal
	trailingWMAIdx := startIdx - lookbackTotal
	today := trailingWMAIdx
	tempReal := inReal[today]
	today++
	periodWMASub := tempReal
	periodWMASum := tempReal
	tempReal = inReal[today]
	today++
	periodWMASub += tempReal
	periodWMASum += tempReal * 2.0
	tempReal = inReal[today]
	today++
	periodWMASub += tempReal
	periodWMASum += tempReal * 3.0
	trailingWMAValue := 0.0
	i := 9
	smoothedValue := 0.0
	for ok := true; ok; {
		tempReal = inReal[today]
		today++
		periodWMASub += tempReal
		periodWMASub -= trailingWMAValue
		periodWMASum += tempReal * 4.0
		trailingWMAValue = inReal[trailingWMAIdx]
		trailingWMAIdx++
		smoothedValue = periodWMASum * 0.1
		periodWMASum -= periodWMASub
		i--
		ok = i != 0
	}
	hilbertIdx := 0
	detrenderOdd[0] = 0.0
	detrenderOdd[1] = 0.0
	detrenderOdd[2] = 0.0
	detrenderEven[0] = 0.0
	detrenderEven[1] = 0.0
	detrenderEven[2] = 0.0
	detrender := 0.0
	prevDetrenderOdd := 0.0
	prevDetrenderEven := 0.0
	prevDetrenderInputOdd := 0.0
	prevDetrenderInputEven := 0.0

	q1Odd[0] = 0.0
	q1Odd[1] = 0.0
	q1Odd[2] = 0.0
	q1Even[0] = 0.0
	q1Even[1] = 0.0
	q1Even[2] = 0.0
	q1 := 0.0
	prevq1Odd := 0.0
	prevq1Even := 0.0
	prevq1InputOdd := 0.0
	prevq1InputEven := 0.0

	jIOdd[0] = 0.0
	jIOdd[1] = 0.0
	jIOdd[2] = 0.0
	jIEven[0] = 0.0
	jIEven[1] = 0.0
	jIEven[2] = 0.0
	jI := 0.0
	prevjIOdd := 0.0
	prevjIEven := 0.0
	prevjIInputOdd := 0.0
	prevjIInputEven := 0.0

	jQOdd[0] = 0.0
	jQOdd[1] = 0.0
	jQOdd[2] = 0.0
	jQEven[0] = 0.0
	jQEven[1] = 0.0
	jQEven[2] = 0.0
	jQ := 0.0
	prevjQOdd := 0.0
	prevjQEven := 0.0
	prevjQInputOdd := 0.0
	prevjQInputEven := 0.0

	period := 0.0
	outIdx := startIdx
	previ2, prevq2 := 0.0, 0.0
	Re, Im := 0.0, 0.0
	mama, fama := 0.0, 0.0
	i1ForOddPrev3, i1ForEvenPrev3 := 0.0, 0.0
	i1ForOddPrev2, i1ForEvenPrev2 := 0.0, 0.0
	prevPhase := 0.0
	adjustedPrevPeriod := 0.0
	todayValue := 0.0
	hilbertTempReal := 0.0
	for today < len(inReal) {
		adjustedPrevPeriod = (0.075 * period) + 0.54
		todayValue = inReal[today]

		periodWMASub += todayValue
		periodWMASub -= trailingWMAValue
		periodWMASum += todayValue * 4.0
		trailingWMAValue = inReal[trailingWMAIdx]
		trailingWMAIdx++
		smoothedValue = periodWMASum * 0.1
		periodWMASum -= periodWMASub
		q2, i2 := 0.0, 0.0
		tempReal2 := 0.0
		if (today % 2) == 0 {

			hilbertTempReal = a * smoothedValue
			detrender = -detrenderEven[hilbertIdx]
			detrenderEven[hilbertIdx] = hilbertTempReal
			detrender += hilbertTempReal
			detrender -= prevDetrenderEven
			prevDetrenderEven = b * prevDetrenderInputEven
			detrender += prevDetrenderEven
			prevDetrenderInputEven = smoothedValue
			detrender *= adjustedPrevPeriod

			hilbertTempReal = a * detrender
			q1 = -q1Even[hilbertIdx]
			q1Even[hilbertIdx] = hilbertTempReal
			q1 += hilbertTempReal
			q1 -= prevq1Even
			prevq1Even = b * prevq1InputEven
			q1 += prevq1Even
			prevq1InputEven = detrender
			q1 *= adjustedPrevPeriod

			hilbertTempReal = a * i1ForEvenPrev3
			jI = -jIEven[hilbertIdx]
			jIEven[hilbertIdx] = hilbertTempReal
			jI += hilbertTempReal
			jI -= prevjIEven
			prevjIEven = b * prevjIInputEven
			jI += prevjIEven
			prevjIInputEven = i1ForEvenPrev3
			jI *= adjustedPrevPeriod

			hilbertTempReal = a * q1
			jQ = -jQEven[hilbertIdx]
			jQEven[hilbertIdx] = hilbertTempReal
			jQ += hilbertTempReal
			jQ -= prevjQEven
			prevjQEven = b * prevjQInputEven
			jQ += prevjQEven
			prevjQInputEven = q1
			jQ *= adjustedPrevPeriod
			hilbertIdx++
			if hilbertIdx == 3 {
				hilbertIdx = 0
			}
			q2 = (0.2 * (q1 + jI)) + (0.8 * prevq2)
			i2 = (0.2 * (i1ForEvenPrev3 - jQ)) + (0.8 * previ2)
			i1ForOddPrev3 = i1ForOddPrev2
			i1ForOddPrev2 = detrender
			if i1ForEvenPrev3 != 0.0 {
				tempReal2 = (math.Atan(q1/i1ForEvenPrev3) * rad2Deg)
			} else {
				tempReal2 = 0.0
			}
		} else {

			hilbertTempReal = a * smoothedValue
			detrender = -detrenderOdd[hilbertIdx]
			detrenderOdd[hilbertIdx] = hilbertTempReal
			detrender += hilbertTempReal
			detrender -= prevDetrenderOdd
			prevDetrenderOdd = b * prevDetrenderInputOdd
			detrender += prevDetrenderOdd
			prevDetrenderInputOdd = smoothedValue
			detrender *= adjustedPrevPeriod

			hilbertTempReal = a * detrender
			q1 = -q1Odd[hilbertIdx]
			q1Odd[hilbertIdx] = hilbertTempReal
			q1 += hilbertTempReal
			q1 -= prevq1Odd
			prevq1Odd = b * prevq1InputOdd
			q1 += prevq1Odd
			prevq1InputOdd = detrender
			q1 *= adjustedPrevPeriod

			hilbertTempReal = a * i1ForOddPrev3
			jI = -jIOdd[hilbertIdx]
			jIOdd[hilbertIdx] = hilbertTempReal
			jI += hilbertTempReal
			jI -= prevjIOdd
			prevjIOdd = b * prevjIInputOdd
			jI += prevjIOdd
			prevjIInputOdd = i1ForOddPrev3
			jI *= adjustedPrevPeriod

			hilbertTempReal = a * q1
			jQ = -jQOdd[hilbertIdx]
			jQOdd[hilbertIdx] = hilbertTempReal
			jQ += hilbertTempReal
			jQ -= prevjQOdd
			prevjQOdd = b * prevjQInputOdd
			jQ += prevjQOdd
			prevjQInputOdd = q1
			jQ *= adjustedPrevPeriod

			q2 = (0.2 * (q1 + jI)) + (0.8 * prevq2)
			i2 = (0.2 * (i1ForOddPrev3 - jQ)) + (0.8 * previ2)
			i1ForEvenPrev3 = i1ForEvenPrev2
			i1ForEvenPrev2 = detrender
			if i1ForOddPrev3 != 0.0 {
				tempReal2 = (math.Atan(q1/i1ForOddPrev3) * rad2Deg)
			} else {
				tempReal2 = 0.0
			}
		}
		tempReal = prevPhase - tempReal2
		prevPhase = tempReal2
		if tempReal < 1.0 {
			tempReal = 1.0
		}
		if tempReal > 1.0 {
			tempReal = inFastLimit / tempReal
			if tempReal < inSlowLimit {
				tempReal = inSlowLimit
			}
		} else {
			tempReal = inFastLimit
		}
		mama = (tempReal * todayValue) + ((1 - tempReal) * mama)
		tempReal *= 0.5
		fama = (tempReal * mama) + ((1 - tempReal) * fama)
		if today >= startIdx {
			outMAMA[outIdx] = mama
			outFAMA[outIdx] = fama
			outIdx++
		}
		Re = (0.2 * ((i2 * previ2) + (q2 * prevq2))) + (0.8 * Re)
		Im = (0.2 * ((i2 * prevq2) - (q2 * previ2))) + (0.8 * Im)
		prevq2 = q2
		previ2 = i2
		tempReal = period
		if (Im != 0.0) && (Re != 0.0) {
			period = 360.0 / (math.Atan(Im/Re) * rad2Deg)
		}
		tempReal2 = 1.5 * tempReal
		if period > tempReal2 {
			period = tempReal2
		}
		tempReal2 = 0.67 * tempReal
		if period < tempReal2 {
			period = tempReal2
		}
		if period < 6 {
			period = 6
		} else if period > 50 {
			period = 50
		}
		period = (0.2 * period) + (0.8 * tempReal)
		today++
	}
	return outMAMA, outFAMA
}
