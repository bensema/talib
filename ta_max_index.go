package talib

// MaxIndex - Index of highest value over a specified period
func MaxIndex(inReal []float64, inTimePeriod int) []float64 {

	outReal := make([]float64, len(inReal))

	if inTimePeriod < 2 {
		return outReal
	}

	nbInitialElementNeeded := inTimePeriod - 1
	startIdx := nbInitialElementNeeded
	outIdx := startIdx
	today := startIdx
	trailingIdx := startIdx - nbInitialElementNeeded
	highestIdx := -1
	highest := 0.0
	for today < len(inReal) {
		tmp := inReal[today]
		if highestIdx < trailingIdx {
			highestIdx = trailingIdx
			highest = inReal[highestIdx]
			i := highestIdx + 1
			for i <= today {
				tmp := inReal[i]
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
		outReal[outIdx] = float64(highestIdx)
		outIdx++
		trailingIdx++
		today++
	}

	return outReal
}
