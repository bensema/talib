// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
纯Go的Talib库

	Overlap Studies（重叠研究类）
	Momentum Indicators（动量指标类）
	Volume Indicators（成交量指标类）
	Volatility Indicators（波动性指标类）
	Price Transform（价格指标类）
	Cycle Indicators（周期指标类）
	Pattern Recognition（形态识别类）
	Statistic Functions（统计函数类）
	Math Transform（数学变换类）
	Math Operators（数学运算符类）

Overlap Studies（重叠研究类）

	BBANDS               Bollinger Bands
	DEMA                 Double Exponential Moving Average
	EMA                  Exponential Moving Average
	HT_TRENDLINE         Hilbert Transform - Instantaneous Trendline
	KAMA                 Kaufman Adaptive Moving Average
	MA                   Moving average
	MAMA                 MESA Adaptive Moving Average
	MAVP                 Moving average with variable period
	MIDPOINT             MidPoint over period
	MIDPRICE             Midpoint Price over period
	SAR                  Parabolic SAR
	SAREXT               Parabolic SAR - Extended
	SMA                  Simple Moving Average
	T3                   Triple Exponential Moving Average (T3)
	TEMA                 Triple Exponential Moving Average
	TRIMA                Triangular Moving Average
	WMA                  Weighted Moving Average

Momentum Indicators（动量指标类）

	ADX                  Average Directional Movement Index
	ADXR                 Average Directional Movement Index Rating
	APO                  Absolute Price Oscillator
	AROON                Aroon
	AROONOSC             Aroon Oscillator
	BOP                  Balance Of Power
	CCI                  Commodity Channel Index
	CMO                  Chande Momentum Oscillator
	DX                   Directional Movement Index
	MACD                 Moving Average Convergence/Divergence
	MACDEXT              MACD with controllable MA type
	MACDFIX              Moving Average Convergence/Divergence Fix 12/26
	MFI                  Money Flow Index
	MINUS_DI             Minus Directional Indicator
	MINUS_DM             Minus Directional Movement
	MOM                  Momentum
	PLUS_DI              Plus Directional Indicator
	PLUS_DM              Plus Directional Movement
	PPO                  Percentage Price Oscillator
	ROC                  Rate of change : ((price/prevPrice)-1)*100
	ROCP                 Rate of change Percentage: (price-prevPrice)/prevPrice
	ROCR                 Rate of change ratio: (price/prevPrice)
	ROCR100              Rate of change ratio 100 scale: (price/prevPrice)*100
	RSI                  Relative Strength Index
	STOCH                Stochastic
	STOCHF               Stochastic Fast
	STOCHRSI             Stochastic Relative Strength Index
	TRIX                 1-day Rate-Of-Change (ROC) of a Triple Smooth EMA
	ULTOSC               Ultimate Oscillator
	WILLR                Williams' %R

Volume Indicators（成交量指标类）

	AD                   Chaikin A/D Line
	ADOSC                Chaikin A/D Oscillator
	OBV                  On Balance Volume

Volatility Indicators（波动性指标类）

	ATR                  Average True Range
	NATR                 Normalized Average True Range
	TRANGE               True Range

Price Transform（价格指标类）

	AVGPRICE             Average Price
	MEDPRICE             Median Price
	TYPPRICE             Typical Price
	WCLPRICE             Weighted Close Price

Cycle Indicators（周期指标类）

	HT_DCPERIOD          Hilbert Transform - Dominant Cycle Period
	HT_DCPHASE           Hilbert Transform - Dominant Cycle Phase
	HT_PHASOR            Hilbert Transform - Phasor Components
	HT_SINE              Hilbert Transform - SineWave
	HT_TRENDMODE         Hilbert Transform - Trend vs Cycle Mode

Pattern Recognition（形态识别类）

	CDL2CROWS            Two Crows
	CDL3BLACKCROWS       Three Black Crows
	CDL3INSIDE           Three Inside Up/Down
	CDL3LINESTRIKE       Three-Line Strike
	CDL3OUTSIDE          Three Outside Up/Down
	CDL3STARSINSOUTH     Three Stars In The South
	CDL3WHITESOLDIERS    Three Advancing White Soldiers
	CDLABANDONEDBABY     Abandoned Baby
	CDLADVANCEBLOCK      Advance Block
	CDLBELTHOLD          Belt-hold
	CDLBREAKAWAY         Breakaway
	CDLCLOSINGMARUBOZU   Closing Marubozu
	CDLCONCEALBABYSWALL  Concealing Baby Swallow
	CDLCOUNTERATTACK     Counterattack
	CDLDARKCLOUDCOVER    Dark Cloud Cover
	CDLDOJI              Doji
	CDLDOJISTAR          Doji Star
	CDLDRAGONFLYDOJI     Dragonfly Doji
	CDLENGULFING         Engulfing Pattern
	CDLEVENINGDOJISTAR   Evening Doji Star
	CDLEVENINGSTAR       Evening Star
	CDLGAPSIDESIDEWHITE  Up/Down-gap side-by-side white lines
	CDLGRAVESTONEDOJI    Gravestone Doji
	CDLHAMMER            Hammer
	CDLHANGINGMAN        Hanging Man
	CDLHARAMI            Harami Pattern
	CDLHARAMICROSS       Harami Cross Pattern
	CDLHIGHWAVE          High-Wave Candle
	CDLHIKKAKE           Hikkake Pattern
	CDLHIKKAKEMOD        Modified Hikkake Pattern
	CDLHOMINGPIGEON      Homing Pigeon
	CDLIDENTICAL3CROWS   Identical Three Crows
	CDLINNECK            In-Neck Pattern
	CDLINVERTEDHAMMER    Inverted Hammer
	CDLKICKING           Kicking
	CDLKICKINGBYLENGTH   Kicking - bull/bear determined by the longer marubozu
	CDLLADDERBOTTOM      Ladder Bottom
	CDLLONGLEGGEDDOJI    Long Legged Doji
	CDLLONGLINE          Long Line Candle
	CDLMARUBOZU          Marubozu
	CDLMATCHINGLOW       Matching Low
	CDLMATHOLD           Mat Hold
	CDLMORNINGDOJISTAR   Morning Doji Star
	CDLMORNINGSTAR       Morning Star
	CDLONNECK            On-Neck Pattern
	CDLPIERCING          Piercing Pattern
	CDLRICKSHAWMAN       Rickshaw Man
	CDLRISEFALL3METHODS  Rising/Falling Three Methods
	CDLSEPARATINGLINES   Separating Lines
	CDLSHOOTINGSTAR      Shooting Star
	CDLSHORTLINE         Short Line Candle
	CDLSPINNINGTOP       Spinning Top
	CDLSTALLEDPATTERN    Stalled Pattern
	CDLSTICKSANDWICH     Stick Sandwich
	CDLTAKURI            Takuri (Dragonfly Doji with very long lower shadow)
	CDLTASUKIGAP         Tasuki Gap
	CDLTHRUSTING         Thrusting Pattern
	CDLTRISTAR           Tristar Pattern
	CDLUNIQUE3RIVER      Unique 3 River
	CDLUPSIDEGAP2CROWS   Upside Gap Two Crows
	CDLXSIDEGAP3METHODS  Upside/Downside Gap Three Methods

Statistic Functions（统计函数类）

	BETA                 Beta
	CORREL               Pearson's Correlation Coefficient (r)
	LINEARREG            Linear Regression
	LINEARREG_ANGLE      Linear Regression Angle
	LINEARREG_INTERCEPT  Linear Regression Intercept
	LINEARREG_SLOPE      Linear Regression Slope
	STDDEV               Standard Deviation
	TSF                  Time Series Forecast
	VAR                  Variance

Math Transform（数学变换类）

	ACOS                 Vector Trigonometric ACos
	ASIN                 Vector Trigonometric ASin
	ATAN                 Vector Trigonometric ATan
	CEIL                 Vector Ceil
	COS                  Vector Trigonometric Cos
	COSH                 Vector Trigonometric Cosh
	EXP                  Vector Arithmetic Exp
	FLOOR                Vector Floor
	LN                   Vector Log Natural
	LOG10                Vector Log10
	SIN                  Vector Trigonometric Sin
	SINH                 Vector Trigonometric Sinh
	SQRT                 Vector Square Root
	TAN                  Vector Trigonometric Tan
	TANH                 Vector Trigonometric Tanh
	PVI                  Positive Volume Index
	NVI                  Negative Volume Index

Math Operators（数学运算符类）

	ADD                  Vector Arithmetic Add
	DIV                  Vector Arithmetic Div
	MAX                  Highest value over a specified period
	MAXINDEX             Index of highest value over a specified period
	MIN                  Lowest value over a specified period
	MININDEX             Index of lowest value over a specified period
	MINMAX               Lowest and highest values over a specified period
	MINMAXINDEX          Indexes of lowest and highest values over a specified period
	MULT                 Vector Arithmetic Mult
	SUB                  Vector Arithmetic Substraction
	SUM                  Summation
*/
package talib
