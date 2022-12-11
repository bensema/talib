package talib

// BBands - Bollinger Bands
// upperband, middleband, lowerband = BBands(close, timeperiod=5, nbdevup=2, nbdevdn=2, matype=0)
func BBands(inReal []float64, inTimePeriod int, inNbDevUp float64, inNbDevDn float64, inMAType MaType) ([]float64, []float64, []float64) {
	outRealUpperBand := make([]float64, len(inReal))
	outRealMiddleBand := Ma(inReal, inTimePeriod, inMAType)
	outRealLowerBand := make([]float64, len(inReal))
	tempBuffer2 := StdDev(inReal, inTimePeriod, 1.0)

	for i := 0; i < len(inReal); i++ {
		tempReal := tempBuffer2[i]
		tempReal2 := outRealMiddleBand[i]
		if inNbDevUp == inNbDevDn {
			if inNbDevUp == 1.0 {
				outRealUpperBand[i] = tempReal2 + tempReal
				outRealLowerBand[i] = tempReal2 - tempReal
			} else {
				outRealUpperBand[i] = tempReal2 + tempReal*inNbDevUp
				outRealLowerBand[i] = tempReal2 - tempReal*inNbDevUp
			}
		} else if inNbDevUp == 1.0 {
			outRealUpperBand[i] = tempReal2 + tempReal
			outRealLowerBand[i] = tempReal2 - tempReal*inNbDevDn
		} else if inNbDevDn == 1.0 {
			outRealLowerBand[i] = tempReal2 - tempReal
			outRealUpperBand[i] = tempReal2 + tempReal*inNbDevUp
		} else {
			outRealUpperBand[i] = tempReal2 + tempReal*inNbDevUp
			outRealLowerBand[i] = tempReal2 - tempReal*inNbDevDn
		}
	}
	return outRealUpperBand, outRealMiddleBand, outRealLowerBand
}
