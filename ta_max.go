package talib

/*
 * TA_MAX - Highest value over a specified period
 * 区间范围最大
 * noptInTimePeriod 范围
 *
 * Input  = double
 * Output = double
 *
 * Optional Parameters
 * -------------------
 * optInTimePeriod:(From 2 to 100000)
 *    Number of period
 *
 *
 */

func Max(inReal []float64, optInTimePeriod int) []float64 {

	outReal := make([]float64, len(inReal))

	if optInTimePeriod < 2 {
		return outReal
	}

	nbInitialElementNeeded := optInTimePeriod - 1
	startIdx := nbInitialElementNeeded
	outIdx := startIdx
	today := startIdx
	trailingIdx := startIdx - nbInitialElementNeeded
	highestIdx := -1
	highest := 0.0

	for today < len(outReal) {

		tmp := inReal[today]

		if highestIdx < trailingIdx {
			highestIdx = trailingIdx
			highest = inReal[highestIdx]
			i := highestIdx + 1
			for i <= today {
				tmp = inReal[i]
				if tmp > highest {
					highestIdx = i
					highest = tmp
				}
				i++
			}
		} else if tmp >= highest {
			highestIdx = today
			highest = tmp
		}
		outReal[outIdx] = highest
		outIdx++
		trailingIdx++
		today++
	}

	return outReal
}
