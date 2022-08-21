package talib

// MinMaxIndex - Indexes of lowest and highest values over a specified period
func MinMaxIndex(inReal []float64, inTimePeriod int) ([]float64, []float64) {

	outMinIdx := make([]float64, len(inReal))
	outMaxIdx := make([]float64, len(inReal))

	nbInitialElementNeeded := (inTimePeriod - 1)
	startIdx := nbInitialElementNeeded
	outIdx := startIdx
	today := startIdx
	trailingIdx := startIdx - nbInitialElementNeeded
	highestIdx := -1
	highest := 0.0
	lowestIdx := -1
	lowest := 0.0
	for today < len(inReal) {
		tmpLow, tmpHigh := inReal[today], inReal[today]
		if highestIdx < trailingIdx {
			highestIdx = trailingIdx
			highest = inReal[highestIdx]
			i := highestIdx
			i++
			for i <= today {
				tmpHigh = inReal[i]
				if tmpHigh > highest {
					highestIdx = i
					highest = tmpHigh
				}
				i++
			}
		} else if tmpHigh >= highest {
			highestIdx = today
			highest = tmpHigh
		}
		if lowestIdx < trailingIdx {
			lowestIdx = trailingIdx
			lowest = inReal[lowestIdx]
			i := lowestIdx
			i++
			for i <= today {
				tmpLow = inReal[i]
				if tmpLow < lowest {
					lowestIdx = i
					lowest = tmpLow
				}
				i++
			}
		} else if tmpLow <= lowest {
			lowestIdx = today
			lowest = tmpLow
		}
		outMaxIdx[outIdx] = float64(highestIdx)
		outMinIdx[outIdx] = float64(lowestIdx)
		outIdx++
		trailingIdx++
		today++
	}
	return outMinIdx, outMaxIdx
}
