package talib

// Ma - Moving average
func Ma(inReal []float64, inTimePeriod int, inMAType MaType) []float64 {

	outReal := make([]float64, len(inReal))

	if inTimePeriod == 1 {
		copy(outReal, inReal)
		return outReal
	}

	switch inMAType {
	case SMA:
		outReal = Sma(inReal, inTimePeriod)
	case EMA:
		outReal = Ema(inReal, inTimePeriod)
	case WMA:
		outReal = Wma(inReal, inTimePeriod)
	case DEMA:
		outReal = Dema(inReal, inTimePeriod)
	case TEMA:
		outReal = Tema(inReal, inTimePeriod)
	case TRIMA:
		outReal = Trima(inReal, inTimePeriod)
	case KAMA:
		outReal = Kama(inReal, inTimePeriod)
	case MAMA:
		outReal, _ = Mama(inReal, 0.5, 0.05)
	case T3MA:
		outReal = T3(inReal, inTimePeriod, 0.7)
	}
	return outReal
}
