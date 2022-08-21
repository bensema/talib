package talib

/*
ACOS - Vector Trigonometric ACos

Input:

Output:

名称: 反余弦函数，三角函数

简介: 反余弦函数，三角函数

算法: 反余弦函数，三角函数
*/
func ACOS() {

}

/*
AD - Chaikin A/D Line

Input:

Output:

名称：Chaikin A/D Line 累积/派发线（Accumulation/Distribution Line）

简介：Marc Chaikin提出的一种平衡交易量指标，以当日的收盘价位来估算成交流量，用于估定一段时间内该证券累积的资金流量。

算法: AD_todo3
*/
func AD() {

}

/*
ADOSC - Chaikin A/D Oscillator

Input:

Output:

名称：Chaikin A/D Oscillator Chaikin震荡指标

简介：将资金流动情况与价格行为相对比，检测市场中资金流入和流出的情况

算法：fastperiod A/D - slowperiod A/D

研判：

	1、交易信号是背离：看涨背离做多，看跌背离做空
	2、股价与90天移动平均结合，与其他指标结合
	3、由正变负卖出，由负变正买进
*/
func ADOSC() {

}

/*
ADX - Average Directional Movement Index

Input:

Output:

名称: 平均趋向指数

简介: 使用ADX指标，指标判断盘整、振荡和单边趋势。

算法:

一、先决定股价趋势（Directional Movement，DM）是上涨或下跌：
“所谓DM值，今日股价波动幅度大于昨日股价波动幅部分的最大值，可能是创高价的部分或创低价的部分；如果今日股价波动幅度较前一日小，则DM = 0。”
若股价高点持续走高，为上涨趋势，记作 +DM。
若为下跌趋势，记作 -DM。-DM的负号（–）是表示反向趋势（下跌），并非数值为负数。
其他状况：DM = 0。

二、寻找股价的真实波幅（True Range，TR）：
所谓真实波幅（TR）是以最高价，最低价，及前一日收盘价三个价格做比较，求出当日股价波动的最大幅度。

三、趋势方向需经由一段时间来观察，研判上才有意义。一般以14天为指标的观察周期：
先计算出 +DM、–DM及TR的14日算术平均数，得到 +DM14、–DM14及TR14三组数据作为起始值，再计算各自的移动平均值（EMA）。

	+DI14 = +DM/TR14*100
	-DI14 = +DM/TR14*100

	DX = |(+DI14)-(-DI14)| / |(+DI14)+(-DI14)|

	DX运算结果取其绝对值，再将DX作移动平均，得到ADX。

特点：

	ADX无法告诉你趋势的发展方向。
	如果趋势存在，ADX可以衡量趋势的强度。不论上升趋势或下降趋势，ADX看起来都一样。
	ADX的读数越大，趋势越明显。衡量趋势强度时，需要比较几天的ADX 读数，观察ADX究竟是上升或下降。ADX读数上升，代表趋势转强；如果ADX读数下降，意味着趋势转弱。
	当ADX曲线向上攀升，趋势越来越强，应该会持续发展。如果ADX曲线下滑，代表趋势开始转弱，反转的可能性增加。
	单就ADX本身来说，由于指标落后价格走势，所以算不上是很好的指标，不适合单就ADX进行操作。可是，如果与其他指标配合运用，ADX可以确认市场是否存在趋势，并衡量趋势的强度

指标应用：

	+DI与–DI表示多空相反的二个动向，当据此绘出的两条曲线彼此纠结相缠时，代表上涨力道与下跌力道相当，多空势均力敌。当 +DI与–DI彼此穿越时，由下往上的一方其力道开始压过由上往下的另一方，此时出现买卖讯号。
	ADX可作为趋势行情的判断依据，当行情明显朝多空任一方向进行时，ADX数值都会显著上升，趋势走强。若行情呈现盘整格局时，ADX会低于 +DI与–DI二条线。若ADX数值低于20，则不论DI表现如何，均显示市场没有明显趋势。
	ADX持续偏高时，代表“超买”（Overbought）或“超卖”（Oversold）的现象，行情反转的机会将增加，此时则不适宜顺势操作。当ADX数值从上升趋势转为下跌时，则代表行情即将反转；若ADX数值由下跌趋势转为上升时，行情将止跌回升。
	总言之，DMI指标包含4条线：+DI、-DI、ADX和ADXR。+DI代表买盘的强度、-DI代表卖盘的强度；ADX代表趋势的强度、ADXR则为ADX的移动平均。
*/
func ADX() {

}

/*
ADXR - Average Directional Movement Index Rating

Input:

Output:

名称: 平均趋向指数的趋向指数

简介: 使用ADXR指标，指标判断ADX趋势。

算法: ADXR_todo3
*/
func ADXR() {

}

/*
APO - Absolute Price Oscillator

Input:

Output:

名称: APO_todo1

简介: APO_todo2

算法: APO_todo3
*/
func APO() {

}

/*
AROON - Aroon

Input:

Output:

名称: 阿隆指标

简介: 该指标是通过计算自价格达到近期最高值和最低值以来所经过的期间数，阿隆指标帮助你预测价格趋势到趋势区域（或者反过来，从趋势区域到趋势）的变化。

算法:

	Aroon(上升)=[(计算期天数-最高价后的天数)/计算期天数]*100
	Aroon(下降)=[(计算期天数-最低价后的天数)/计算期天数]*100

指数应用

	1、极值0和100
	当UP线达到100时，市场处于强势；如果维持在70100之间，表示一个上升趋势。同样，如果Down线达到0，表示处于弱势，如果维持在030之间，表示处于下跌趋势。如果两条线同处于极值水平，则表明一个更强的趋势。
	2、平行运动
	如果两条线平行运动时，表明市场趋势被打破。可以预期该状况将持续下去，只到由极值水平或交叉穿行西安市出方向性运动为止。
	3、交叉穿行
	当下行线上穿上行线时，表明潜在弱势，预期价格开始趋于下跌。反之，表明潜在强势，预期价格趋于走高。
*/
func AROON() {

}

/*
AROONOSC - Aroon Oscillator

Input:

Output:

名称: 阿隆振荡

简介: AROONOSC_todo2

算法: AROONOSC_todo3
*/
func AROONOSC() {

}

/*
ASIN - Vector Trigonometric ASin

Input:

Output:

名称: 反正弦函数，三角函数

简介: 反正弦函数，三角函数

算法: 反正弦函数，三角函数
*/
func ASIN() {

}

/*
ATAN - Vector Trigonometric ATan

Input:

Output:

名称: 数字的反正切值，三角函数

简介: 数字的反正切值，三角函数

算法: 数字的反正切值，三角函数
*/
func ATAN() {

}

/*
ATR - Average True Range

Input:

Output:

名称：真实波动幅度均值

简介：真实波动幅度均值（ATR)是 以 N 天的指数移动平均数平均後的交易波动幅度。

算法：

	一天的交易幅度只是单纯地 最大值 - 最小值。
	而真实波动幅度则包含昨天的收盘价，若其在今天的幅度之外：
	真实波动幅度 = max(最大值,昨日收盘价) − min(最小值,昨日收盘价) 真实波动幅度均值便是「真实波动幅度」的 N 日 指数移动平均数。

特性:

	波动幅度的概念表示可以显示出交易者的期望和热情。
	大幅的或增加中的波动幅度表示交易者在当天可能准备持续买进或卖出股票。
	波动幅度的减少则表示交易者对股市没有太大的兴趣。
*/
func ATR() {

}

/*
AVGPRICE - Average Price

Input:

Output:

名称: 平均价格函数

简介: AVGPRICE_todo2

算法: AVGPRICE_todo3
*/
func AVGPRICE() {

}

/*
BBANDS - Bollinger Bands

Input:

Output:

名称: 布林线指标

简介: 其利用统计原理，求出股价的标准差及其信赖区间，从而确定股价的波动范围及未来走势，利用波带显示股价的安全高低价位，因而也被称为布林带。

算法: BBANDS_todo3
*/
func BBANDS() {

}

/*
BETA - Beta

Input:

Output:

名称: β系数也称为贝塔系数

简介: 一种风险指数，用来衡量个别股票或 股票基金相对于整个股市的价格波动情况 贝塔系数衡量股票收益相对于业绩评价基准收益的总体波动性，是一个相对指标。 β 越高，意味着股票相对于业绩评价基准的波动性越大。 β 大于 1 ， 则股票的波动性大于业绩评价基准的波动性。反之亦然。

算法: BETA_todo3

用途：

	1）计算资本成本，做出投资决策（只有回报率高于资本成本的项目才应投资）；
	2）计算资本成本，制定业绩考核及激励标准；
	3）计算资本成本，进行资产估值（Beta是现金流贴现模型的基础）；
	4）确定单个资产或组合的系统风险，用于资产组合的投资管理，特别是股指期货或其他金融衍生品的避险（或投机）
*/
func BETA() {

}

/*
BOP - Balance Of Power

Input:

Output:

名称: 均势指标

简介: BOP_todo2

算法: BOP_todo3
*/
func BOP() {

}

/*
CCI - Commodity Channel Index

Input:

Output:

名称: 顺势指标

简介: CCI指标专门测量股价是否已超出常态分布范围

算法: CCI_todo3

指标应用

	1.当CCI指标曲线在+100线～-100线的常态区间里运行时,CCI指标参考意义不大，可以用KDJ等其它技术指标进行研判。
	2.当CCI指标曲线从上向下突破+100线而重新进入常态区间时，表明市场价格的上涨阶段可能结束，将进入一个比较长时间的震荡整理阶段，应及时平多做空。
	3.当CCI指标曲线从上向下突破-100线而进入另一个非常态区间（超卖区）时，表明市场价格的弱势状态已经形成，将进入一个比较长的寻底过程，可以持有空单等待更高利润。如果CCI指标曲线在超卖区运行了相当长的一段时间后开始掉头向上，表明价格的短期底部初步探明，可以少量建仓。CCI指标曲线在超卖区运行的时间越长，确认短期的底部的准确度越高。
	4.CCI指标曲线从下向上突破-100线而重新进入常态区间时，表明市场价格的探底阶段可能结束，有可能进入一个盘整阶段，可以逢低少量做多。
	5.CCI指标曲线从下向上突破+100线而进入非常态区间(超买区)时，表明市场价格已经脱离常态而进入强势状态，如果伴随较大的市场交投，应及时介入成功率将很大。
	6.CCI指标曲线从下向上突破+100线而进入非常态区间(超买区)后，只要CCI指标曲线一直朝上运行，表明价格依然保持强势可以继续持有待涨。但是，如果在远离+100线的地方开始掉头向下时，则表明市场价格的强势状态将可能难以维持，涨势可能转弱，应考虑卖出。如果前期的短期涨幅过高同时价格回落时交投活跃，则应该果断逢高卖出或做空。
	CCI主要是在超买和超卖区域发生作用，对急涨急跌的行情检测性相对准确。非常适用于股票、外汇、贵金属等市场的短期操作。[1]
*/
func CCI() {

}

/*
CDL2CROWS - Two Crows

Input:

Output:

名称：Two Crows 两只乌鸦

简介：三日K线模式，第一天长阳，第二天高开收阴，第三天再次高开继续收阴， 收盘比前一日收盘价低，预示股价下跌。

算法: CDL2CROWS_todo3
*/
func CDL2CROWS() {

}

/*
CDL3BLACKCROWS - Three Black Crows

Input:

Output:

名称：Three Black Crows 三只乌鸦

简介：三日K线模式，连续三根阴线，每日收盘价都下跌且接近最低价， 每日开盘价都在上根K线实体内，预示股价下跌。

算法: CDL3BLACKCROWS_todo3
*/
func CDL3BLACKCROWS() {

}

/*
CDL3INSIDE - Three Inside Up/Down

Input:

Output:

名称： Three-Line Strike 三线打击

简介：四日K线模式，前三根阳线，每日收盘价都比前一日高， 开盘价在前一日实体内，第四日市场高开，收盘价低于第一日开盘价，预示股价下跌。

算法: CDL3INSIDE_todo3
*/
func CDL3INSIDE() {

}

/*
CDL3LINESTRIKE - Three-Line Strike

Input:

Output:

名称: CDL3LINESTRIKE_todo1

简介: CDL3LINESTRIKE_todo2

算法: CDL3LINESTRIKE_todo3
*/
func CDL3LINESTRIKE() {

}

/*
CDL3OUTSIDE - Three Outside Up/Down

Input:

Output:

名称：Three Outside Up/Down 三外部上涨和下跌

简介：三日K线模式，与三内部上涨和下跌类似，K线为阴阳阳，但第一日与第二日的K线形态相反， 以三外部上涨为例，第一日K线在第二日K线内部，预示着股价上涨。

算法: CDL3OUTSIDE_todo3
*/
func CDL3OUTSIDE() {

}

/*
CDL3STARSINSOUTH - Three Stars In The South

Input:

Output:

名称：Three Stars In The South 南方三星

简介：三日K线模式，与大敌当前相反，三日K线皆阴，第一日有长下影线， 第二日与第一日类似，K线整体小于第一日，第三日无下影线实体信号， 成交价格都在第一日振幅之内，预示下跌趋势反转，股价上升。

算法: CDL3STARSINSOUTH_todo3
*/
func CDL3STARSINSOUTH() {

}

/*
CDL3WHITESOLDIERS - Three Advancing White Soldiers

Input:

Output:

名称：Three Advancing White Soldiers 三个白兵

简介：三日K线模式，三日K线皆阳， 每日收盘价变高且接近最高价，开盘价在前一日实体上半部，预示股价上升。

算法: CDL3WHITESOLDIERS_todo3
*/
func CDL3WHITESOLDIERS() {

}

/*
CDLABANDONEDBABY - Abandoned Baby

Input:

Output:

名称：Abandoned Baby 弃婴

简介：三日K线模式，第二日价格跳空且收十字星（开盘价与收盘价接近， 最高价最低价相差不大），预示趋势反转，发生在顶部下跌，底部上涨。

算法: CDLABANDONEDBABY_todo3
*/
func CDLABANDONEDBABY() {

}

/*
CDLADVANCEBLOCK - Advance Block

Input:

Output:

名称：Advance Block 大敌当前

简介：三日K线模式，三日都收阳，每日收盘价都比前一日高， 开盘价都在前一日实体以内，实体变短，上影线变长。

算法: CDLADVANCEBLOCK_todo3
*/
func CDLADVANCEBLOCK() {

}

/*
CDLBELTHOLD - Belt-hold

Input:

Output:

名称：Belt-hold 捉腰带线

简介：两日K线模式，下跌趋势中，第一日阴线， 第二日开盘价为最低价，阳线，收盘价接近最高价，预示价格上涨。

算法: CDLBELTHOLD_todo3
*/
func CDLBELTHOLD() {

}

/*
CDLBREAKAWAY - Breakaway

Input:

Output:

名称：Breakaway 脱离

简介：五日K线模式，以看涨脱离为例，下跌趋势中，第一日长阴线，第二日跳空阴线，延续趋势开始震荡， 第五日长阳线，收盘价在第一天收盘价与第二天开盘价之间，预示价格上涨。

算法: CDLBREAKAWAY_todo3
*/
func CDLBREAKAWAY() {

}

/*
CDLCLOSINGMARUBOZU - Closing Marubozu

Input:

Output:

名称：Closing Marubozu 收盘缺影线

简介：一日K线模式，以阳线为例，最低价低于开盘价，收盘价等于最高价， 预示着趋势持续。

算法: CDLCLOSINGMARUBOZU_todo3
*/
func CDLCLOSINGMARUBOZU() {

}

/*
CDLCONCEALBABYSWALL - Concealing Baby Swallow

Input:

Output:

名称： Concealing Baby Swallow 藏婴吞没

简介：四日K线模式，下跌趋势中，前两日阴线无影线 ，第二日开盘、收盘价皆低于第二日，第三日倒锤头， 第四日开盘价高于前一日最高价，收盘价低于前一日最低价，预示着底部反转。

算法: CDLCONCEALBABYSWALL_todo3
*/
func CDLCONCEALBABYSWALL() {

}

/*
CDLCOUNTERATTACK - Counterattack

Input:

Output:

名称：Counterattack 反击线

简介：二日K线模式，与分离线类似。

算法: CDLCOUNTERATTACK_todo3
*/
func CDLCOUNTERATTACK() {

}

/*
CDLDARKCLOUDCOVER - Dark Cloud Cover

Input:

Output:

名称：Dark Cloud Cover 乌云压顶

简介：二日K线模式，第一日长阳，第二日开盘价高于前一日最高价， 收盘价处于前一日实体中部以下，预示着股价下跌。

算法: CDLDARKCLOUDCOVER_todo3
*/
func CDLDARKCLOUDCOVER() {

}

/*
CDLDOJI - Doji

Input:

Output:

名称：Doji 十字

简介：一日K线模式，开盘价与收盘价基本相同。

算法: CDLDOJI_todo3
*/
func CDLDOJI() {

}

/*
CDLDOJISTAR - Doji Star

Input:

Output:

名称：Doji Star 十字星

简介：一日K线模式，开盘价与收盘价基本相同，上下影线不会很长，预示着当前趋势反转。

算法: CDLDOJISTAR_todo3
*/
func CDLDOJISTAR() {

}

/*
CDLDRAGONFLYDOJI - Dragonfly Doji

Input:

Output:

名称：Dragonfly Doji 蜻蜓十字/T形十字

简介：一日K线模式，开盘后价格一路走低， 之后收复，收盘价与开盘价相同，预示趋势反转。

算法: CDLDRAGONFLYDOJI_todo3
*/
func CDLDRAGONFLYDOJI() {

}

/*
CDLENGULFING - Engulfing Pattern

Input:

Output:

名称：Engulfing Pattern 吞噬模式

简介：两日K线模式，分多头吞噬和空头吞噬，以多头吞噬为例，第一日为阴线， 第二日阳线，第一日的开盘价和收盘价在第二日开盘价收盘价之内，但不能完全相同。

算法: CDLENGULFING_todo3
*/
func CDLENGULFING() {

}

/*
CDLEVENINGDOJISTAR - Evening Doji Star

Input:

Output:

名称：Evening Doji Star 十字暮星

简介：三日K线模式，基本模式为暮星，第二日收盘价和开盘价相同，预示顶部反转。

算法: CDLEVENINGDOJISTAR_todo3
*/
func CDLEVENINGDOJISTAR() {

}

/*
CDLEVENINGSTAR - Evening Star

Input:

Output:

名称：Evening Star 暮星

简介：三日K线模式，与晨星相反，上升趋势中, 第一日阳线，第二日价格振幅较小，第三日阴线，预示顶部反转。

算法: CDLEVENINGSTAR_todo3
*/
func CDLEVENINGSTAR() {

}

/*
CDLGAPSIDESIDEWHITE - Up/Down-gap side-by-side white lines

Input:

Output:

名称：Up/Down-gap side-by-side white lines 向上/下跳空并列阳线

简介：二日K线模式，上升趋势向上跳空，下跌趋势向下跳空, 第一日与第二日有相同开盘价，实体长度差不多，则趋势持续。

算法: CDLGAPSIDESIDEWHITE_todo3
*/
func CDLGAPSIDESIDEWHITE() {

}

/*
CDLGRAVESTONEDOJI - Gravestone Doji

Input:

Output:

名称：Gravestone Doji 墓碑十字/倒T十字

简介：一日K线模式，开盘价与收盘价相同，上影线长，无下影线，预示底部反转。

算法: CDLGRAVESTONEDOJI_todo3
*/
func CDLGRAVESTONEDOJI() {

}

/*
CDLHAMMER - Hammer

Input:

Output:

名称：Hammer 锤头

简介：一日K线模式，实体较短，无上影线， 下影线大于实体长度两倍，处于下跌趋势底部，预示反转。

算法: CDLHAMMER_todo3
*/
func CDLHAMMER() {

}

/*
CDLHANGINGMAN - Hanging Man

Input:

Output:

名称：Hanging Man 上吊线

简介：一日K线模式，形状与锤子类似，处于上升趋势的顶部，预示着趋势反转。

算法: CDLHANGINGMAN_todo3
*/
func CDLHANGINGMAN() {

}

/*
CDLHARAMI - Harami Pattern

Input:

Output:

名称：Harami Pattern 母子线

简介：二日K线模式，分多头母子与空头母子，两者相反，以多头母子为例，在下跌趋势中，第一日K线长阴， 第二日开盘价收盘价在第一日价格振幅之内，为阳线，预示趋势反转，股价上升。

算法: CDLHARAMI_todo3
*/
func CDLHARAMI() {

}

/*
CDLHARAMICROSS - Harami Cross Pattern

Input:

Output:

名称：Harami Cross Pattern 十字孕线

简介：二日K线模式，与母子县类似，若第二日K线是十字线， 便称为十字孕线，预示着趋势反转。

算法: CDLHARAMICROSS_todo3
*/
func CDLHARAMICROSS() {

}

/*
CDLHIGHWAVE - High-Wave Candle

Input:

Output:

名称：High-Wave Candle 风高浪大线

简介：三日K线模式，具有极长的上/下影线与短的实体，预示着趋势反转。

算法: CDLHIGHWAVE_todo3
*/
func CDLHIGHWAVE() {

}

/*
CDLHIKKAKE - Hikkake Pattern

Input:

Output:

名称：Hikkake Pattern 陷阱

简介：三日K线模式，与母子类似，第二日价格在前一日实体范围内, 第三日收盘价高于前两日，反转失败，趋势继续。

算法: CDLHIKKAKE_todo3
*/
func CDLHIKKAKE() {

}

/*
CDLHIKKAKEMOD - Modified Hikkake Pattern

Input:

Output:

名称：Modified Hikkake Pattern 修正陷阱

简介：三日K线模式，与陷阱类似，上升趋势中，第三日跳空高开； 下跌趋势中，第三日跳空低开，反转失败，趋势继续。

算法: CDLHIKKAKEMOD_todo3
*/
func CDLHIKKAKEMOD() {

}

/*
CDLHOMINGPIGEON - Homing Pigeon

Input:

Output:

名称：Homing Pigeon 家鸽

简介：二日K线模式，与母子线类似，不同的的是二日K线颜色相同， 第二日最高价、最低价都在第一日实体之内，预示着趋势反转。

算法: CDLHOMINGPIGEON_todo3
*/
func CDLHOMINGPIGEON() {

}

/*
CDLIDENTICAL3CROWS - Identical Three Crows

Input:

Output:

名称：Identical Three Crows 三胞胎乌鸦

简介：三日K线模式，上涨趋势中，三日都为阴线，长度大致相等， 每日开盘价等于前一日收盘价，收盘价接近当日最低价，预示价格下跌。

算法: CDLIDENTICAL3CROWS_todo3
*/
func CDLIDENTICAL3CROWS() {

}

/*
CDLINNECK - In-Neck Pattern

Input:

Output:

名称：In-Neck Pattern 颈内线

简介：二日K线模式，下跌趋势中，第一日长阴线， 第二日开盘价较低，收盘价略高于第一日收盘价，阳线，实体较短，预示着下跌继续。

算法: CDLINNECK_todo3
*/
func CDLINNECK() {

}

/*
CDLINVERTEDHAMMER - Inverted Hammer

Input:

Output:

名称：Inverted Hammer 倒锤头

简介：一日K线模式，上影线较长，长度为实体2倍以上， 无下影线，在下跌趋势底部，预示着趋势反转。

算法: CDLINVERTEDHAMMER_todo3
*/
func CDLINVERTEDHAMMER() {

}

/*
CDLKICKING - Kicking

Input:

Output:

名称：Kicking 反冲形态

简介：二日K线模式，与分离线类似，两日K线为秃线，颜色相反，存在跳空缺口。

算法: CDLKICKING_todo3
*/
func CDLKICKING() {

}

/*
CDLKICKINGBYLENGTH - Kicking - bull/bear determined by the longer marubozu

Input:

Output:

名称：Kicking - bull/bear determined by the longer marubozu 由较长缺影线决定的反冲形态

简介：二日K线模式，与反冲形态类似，较长缺影线决定价格的涨跌。

算法: CDLKICKINGBYLENGTH_todo3
*/
func CDLKICKINGBYLENGTH() {

}

/*
CDLLADDERBOTTOM - Ladder Bottom

Input:

Output:

名称：Ladder Bottom 梯底

简介：五日K线模式，下跌趋势中，前三日阴线， 开盘价与收盘价皆低于前一日开盘、收盘价，第四日倒锤头，第五日开盘价高于前一日开盘价， 阳线，收盘价高于前几日价格振幅，预示着底部反转。

算法: CDLLADDERBOTTOM_todo3
*/
func CDLLADDERBOTTOM() {

}

/*
CDLLONGLEGGEDDOJI - Long Legged Doji

Input:

Output:

名称：Long Legged Doji 长脚十字

简介：一日K线模式，开盘价与收盘价相同居当日价格中部，上下影线长， 表达市场不确定性。

算法: CDLLONGLEGGEDDOJI_todo3
*/
func CDLLONGLEGGEDDOJI() {

}

/*
CDLLONGLINE - Long Line Candle

Input:

Output:

名称：Long Line Candle 长蜡烛

简介：一日K线模式，K线实体长，无上下影线。

算法: CDLLONGLINE_todo3
*/
func CDLLONGLINE() {

}

/*
CDLMARUBOZU - Marubozu

Input:

Output:

名称：Marubozu 光头光脚/缺影线

简介：一日K线模式，上下两头都没有影线的实体， 阴线预示着熊市持续或者牛市反转，阳线相反。

算法: CDLMARUBOZU_todo3
*/
func CDLMARUBOZU() {

}

/*
CDLMATCHINGLOW - Matching Low

Input:

Output:

名称：Matching Low 相同低价

简介：二日K线模式，下跌趋势中，第一日长阴线， 第二日阴线，收盘价与前一日相同，预示底部确认，该价格为支撑位。

算法: CDLMATCHINGLOW_todo3
*/
func CDLMATCHINGLOW() {

}

/*
CDLMATHOLD - Mat Hold

Input:

Output:

名称：Mat Hold 铺垫

简介：五日K线模式，上涨趋势中，第一日阳线，第二日跳空高开影线， 第三、四日短实体影线，第五日阳线，收盘价高于前四日，预示趋势持续。

算法: CDLMATHOLD_todo3
*/
func CDLMATHOLD() {

}

/*
CDLMORNINGDOJISTAR - Morning Doji Star

Input:

Output:

名称：Morning Doji Star 十字晨星

简介：三日K线模式， 基本模式为晨星，第二日K线为十字星，预示底部反转。

算法: CDLMORNINGDOJISTAR_todo3
*/
func CDLMORNINGDOJISTAR() {

}

/*
CDLMORNINGSTAR - Morning Star

Input:

Output:

名称：Morning Star 晨星

简介：三日K线模式，下跌趋势，第一日阴线， 第二日价格振幅较小，第三天阳线，预示底部反转。

算法: CDLMORNINGSTAR_todo3
*/
func CDLMORNINGSTAR() {

}

/*
CDLONNECK - On-Neck Pattern

Input:

Output:

名称：On-Neck Pattern 颈上线

简介：二日K线模式，下跌趋势中，第一日长阴线，第二日开盘价较低， 收盘价与前一日最低价相同，阳线，实体较短，预示着延续下跌趋势。

算法: CDLONNECK_todo3
*/
func CDLONNECK() {

}

/*
CDLPIERCING - Piercing Pattern

Input:

Output:

名称：Piercing Pattern 刺透形态

简介：两日K线模式，下跌趋势中，第一日阴线，第二日收盘价低于前一日最低价， 收盘价处在第一日实体上部，预示着底部反转。

算法: CDLPIERCING_todo3
*/
func CDLPIERCING() {

}

/*
CDLRICKSHAWMAN - Rickshaw Man

Input:

Output:

名称：Rickshaw Man 黄包车夫

简介：一日K线模式，与长腿十字线类似， 若实体正好处于价格振幅中点，称为黄包车夫。

算法: CDLRICKSHAWMAN_todo3
*/
func CDLRICKSHAWMAN() {

}

/*
CDLRISEFALL3METHODS - Rising/Falling Three Methods

Input:

Output:

名称：Rising/Falling Three Methods 上升/下降三法

简介： 五日K线模式，以上升三法为例，上涨趋势中， 第一日长阳线，中间三日价格在第一日范围内小幅震荡， 第五日长阳线，收盘价高于第一日收盘价，预示股价上升。

算法: CDLRISEFALL3METHODS_todo3
*/
func CDLRISEFALL3METHODS() {

}

/*
CDLSEPARATINGLINES - Separating Lines

Input:

Output:

名称：Separating Lines 分离线

简介：二日K线模式，上涨趋势中，第一日阴线，第二日阳线， 第二日开盘价与第一日相同且为最低价，预示着趋势继续。

算法: CDLSEPARATINGLINES_todo3
*/
func CDLSEPARATINGLINES() {

}

/*
CDLSHOOTINGSTAR - Shooting Star

Input:

Output:

名称：Shooting Star 射击之星

简介：一日K线模式，上影线至少为实体长度两倍， 没有下影线，预示着股价下跌

算法: CDLSHOOTINGSTAR_todo3
*/
func CDLSHOOTINGSTAR() {

}

/*
CDLSHORTLINE - Short Line Candle

Input:

Output:

名称：Short Line Candle 短蜡烛

简介：一日K线模式，实体短，无上下影线

算法: CDLSHORTLINE_todo3
*/
func CDLSHORTLINE() {

}

/*
CDLSPINNINGTOP - Spinning Top

Input:

Output:

名称：Spinning Top 纺锤

简介：一日K线，实体小。

算法: CDLSPINNINGTOP_todo3
*/
func CDLSPINNINGTOP() {

}

/*
CDLSTALLEDPATTERN - Stalled Pattern

Input:

Output:

名称：Stalled Pattern 停顿形态

简介：三日K线模式，上涨趋势中，第二日长阳线， 第三日开盘于前一日收盘价附近，短阳线，预示着上涨结束

算法: CDLSTALLEDPATTERN_todo3
*/
func CDLSTALLEDPATTERN() {

}

/*
CDLSTICKSANDWICH - Stick Sandwich

Input:

Output:

名称：Stick Sandwich 条形三明治

简介：三日K线模式，第一日长阴线，第二日阳线，开盘价高于前一日收盘价， 第三日开盘价高于前两日最高价，收盘价于第一日收盘价相同。

算法: CDLSTICKSANDWICH_todo3
*/
func CDLSTICKSANDWICH() {

}

/*
CDLTAKURI - Takuri (Dragonfly Doji with very long lower shadow)

Input:

Output:

名称：Takuri (Dragonfly Doji with very long lower shadow) 探水竿

简介：一日K线模式，大致与蜻蜓十字相同，下影线长度长。

算法: CDLTAKURI_todo3
*/
func CDLTAKURI() {

}

/*
CDLTASUKIGAP - Tasuki Gap

Input:

Output:

名称：Tasuki Gap 跳空并列阴阳线

简介：三日K线模式，分上涨和下跌，以上升为例， 前两日阳线，第二日跳空，第三日阴线，收盘价于缺口中，上升趋势持续。

算法: CDLTASUKIGAP_todo3
*/
func CDLTASUKIGAP() {

}

/*
CDLTHRUSTING - Thrusting Pattern

Input:

Output:

名称：Thrusting Pattern 插入

简介：二日K线模式，与颈上线类似，下跌趋势中，第一日长阴线，第二日开盘价跳空， 收盘价略低于前一日实体中部，与颈上线相比实体较长，预示着趋势持续。

算法: CDLTHRUSTING_todo3
*/
func CDLTHRUSTING() {

}

/*
CDLTRISTAR - Tristar Pattern

Input:

Output:

名称：Tristar Pattern 三星

简介：三日K线模式，由三个十字组成， 第二日十字必须高于或者低于第一日和第三日，预示着反转。

算法: CDLTRISTAR_todo3
*/
func CDLTRISTAR() {

}

/*
CDLUNIQUE3RIVER - Unique 3 River

Input:

Output:

名称：Unique 3 River 奇特三河床

简介：三日K线模式，下跌趋势中，第一日长阴线，第二日为锤头，最低价创新低，第三日开盘价低于第二日收盘价，收阳线， 收盘价不高于第二日收盘价，预示着反转，第二日下影线越长可能性越大。

算法: CDLUNIQUE3RIVER_todo3
*/
func CDLUNIQUE3RIVER() {

}

/*
CDLUPSIDEGAP2CROWS - Upside Gap Two Crows

Input:

Output:

名称：Upside Gap Two Crows 向上跳空的两只乌鸦

简介：三日K线模式，第一日阳线，第二日跳空以高于第一日最高价开盘， 收阴线，第三日开盘价高于第二日，收阴线，与第一日比仍有缺口。

算法: CDLUPSIDEGAP2CROWS_todo3
*/
func CDLUPSIDEGAP2CROWS() {

}

/*
CDLXSIDEGAP3METHODS - Upside/Downside Gap Three Methods

Input:

Output:

名称：Upside/Downside Gap Three Methods 上升/下降跳空三法

简介：五日K线模式，以上升跳空三法为例，上涨趋势中，第一日长阳线，第二日短阳线，第三日跳空阳线，第四日阴线，开盘价与收盘价于前两日实体内， 第五日长阳线，收盘价高于第一日收盘价，预示股价上升。

算法: CDLXSIDEGAP3METHODS_todo3
*/
func CDLXSIDEGAP3METHODS() {

}

/*
CEIL - Vector Ceil

Input:

Output:

名称: 向上取整数

简介: 向上取整数

算法: 向上取整数
*/
func CEIL() {

}

/*
CMO - Chande Momentum Oscillator

Input:

Output:

名称: 钱德动量摆动指标

简介: 与其他动量指标摆动指标如相对强弱指标（RSI）和随机指标（KDJ）不同，钱德动量指标在计算公式的分子中采用上涨日和下跌日的数据。

算法:

	CMO=（Su－Sd）*100/（Su+Sd）
	其中：Su是今日收盘价与昨日收盘价（上涨日）差值加总。若当日下跌，则增加值为0；Sd是今日收盘价与做日收盘价（下跌日）差值的绝对值加总。若当日上涨，则增加值为0；

指标应用

	本指标类似RSI指标。
	当本指标下穿-50水平时是买入信号，上穿+50水平是卖出信号。
	钱德动量摆动指标的取值介于-100和100之间。
	本指标也能给出良好的背离信号。
	当股票价格创出新低而本指标未能创出新低时，出现牛市背离；
	当股票价格创出新高而本指标未能创出新高时，当出现熊市背离时。
	我们可以用移动均值对该指标进行平滑。
*/
func CMO() {

}

/*
CORREL - Pearson's Correlation Coefficient (r)

Input:

Output:

名称: 皮尔逊相关系数

简介: 用于度量两个变量X和Y之间的相关（线性相关），其值介于-1与1之间皮尔逊相关系数是一种度量两个变量间相关程度的方法。它是一个介于 1 和 -1 之间的值， 其中，1 表示变量完全正相关， 0 表示无关，-1 表示完全负相关。

算法: CORREL_todo3
*/
func CORREL() {

}

/*
COS - Vector Trigonometric Cos

Input:

Output:

名称: 余弦函数，三角函数

简介: 余弦函数，三角函数

算法: 余弦函数，三角函数
*/
func COS() {

}

/*
COSH - Vector Trigonometric Cosh

Input:

Output:

名称: 双曲正弦函数，三角函数

简介: 双曲正弦函数，三角函数

算法: 双曲正弦函数，三角函数
*/
func COSH() {

}

/*
DEMA - Double Exponential Moving Average

Input:

Output:

名称: 双移动平均线

简介: 两条移动平均线来产生趋势信号，较长期者用来识别趋势，较短期者用来选择时机。正是两条平均线及价格三者的相互作用，才共同产生了趋势信号。

算法: DEMA_todo3
*/
func DEMA() {

}

/*
DX - Directional Movement Index

Input:

Output:

名称: 动向指标或趋向指标

简介: 通过分析股票价格在涨跌过程中买卖双方力量均衡点的变化情况，即多空双方的力量的变化受价格波动的影响而发生由均衡到失衡的循环过程，从而提供对趋势判断依据的一种技术指标。

算法: DX_todo3
*/
func DX() {

}

/*
EMA - Exponential Moving Average

Input:

Output:

名称: 指数平均数

简介: 是一种趋向类指标，其构造原理是仍然对价格收盘价进行算术平均，并根据计算结果来进行分析，用于判断价格未来走势的变动趋势。

算法: EMA_todo3
*/
func EMA() {

}

/*
EXP - Vector Arithmetic Exp

Input:

Output:

名称: 指数曲线，三角函数

简介: 指数曲线，三角函数

算法: 指数曲线，三角函数
*/
func EXP() {

}

/*
FLOOR - Vector Floor

Input:

Output:

名称: 向下取整数

简介: 向下取整数

算法: 向下取整数
*/
func FLOOR() {

}

/*
HT_DCPERIOD - Hilbert Transform - Dominant Cycle Period

Input:

Output:

名称： 希尔伯特变换-主导周期

简介：将价格作为信息信号，计算价格处在的周期的位置，作为择时的依

算法: HT_DCPERIOD_todo3
*/
func HT_DCPERIOD() {

}

/*
HT_DCPHASE - Hilbert Transform - Dominant Cycle Phase

Input:

Output:

名称: 希尔伯特变换-主导循环阶段

简介: HT_DCPHASE_todo2

算法: HT_DCPHASE_todo3
*/
func HT_DCPHASE() {

}

/*
HT_PHASOR - Hilbert Transform - Phasor Components

Input:

Output:

名称: 希尔伯特变换-希尔伯特变换相量分量

简介: HT_PHASOR_todo2

算法: HT_PHASOR_todo3
*/
func HT_PHASOR() {

}

/*
HT_SINE - Hilbert Transform - SineWave

Input:

Output:

名称: 希尔伯特变换-正弦波

简介: HT_SINE_todo2

算法: HT_SINE_todo3
*/
func HT_SINE() {

}

/*
HT_TRENDLINE - Hilbert Transform - Instantaneous Trendline

Input:

Output:

名称: 希尔伯特瞬时变换

简介: 是一种趋向类指标，其构造原理是仍然对价格收盘价进行算术平均，并根据计算结果来进行分析，用于判断价格未来走势的变动趋势。

算法: HT_TRENDLINE_todo3
*/
func HT_TRENDLINE() {

}

/*
HT_TRENDMODE - Hilbert Transform - Trend vs Cycle Mode

Input:

Output:

名称: 希尔伯特变换-趋势与周期模式

简介: HT_TRENDMODE_todo2

算法: HT_TRENDMODE_todo3
*/
func HT_TRENDMODE() {

}

/*
KAMA - Kaufman Adaptive Moving Average

Input:

Output:

名称: 考夫曼的自适应移动平均线

简介: 短期均线贴近价格走势，灵敏度高，但会有很多噪声，产生虚假信号；长期均线在判断趋势上一般比较准确 ，但是长期均线有着严重滞后的问题。我们想得到这样的均线，当价格沿一个方向快速移动时，短期的移动 平均线是最合适的；当价格在横盘的过程中，长期移动平均线是合适的。

算法: KAMA_todo3
*/
func KAMA() {

}

/*
LINEARREG - Linear Regression

Input:

Output:

名称: 线性回归

简介: 来确定两种或两种以上变量间相互依赖的定量关系的一种统计分析方法 其表达形式为y = w'x+e，e为误差服从均值为0的正态分布。

算法: LINEARREG_todo3
*/
func LINEARREG() {

}

/*
LINEARREG_ANGLE - Linear Regression Angle

Input:

Output:

名称: 线性回归的角度

简介: LINEARREG_ANGLE_todo2

算法: LINEARREG_ANGLE_todo3
*/
func LINEARREG_ANGLE() {

}

/*
LINEARREG_INTERCEPT - Linear Regression Intercept

Input:

Output:

名称: 线性回归截距

简介: LINEARREG_INTERCEPT_todo2

算法: LINEARREG_INTERCEPT_todo3
*/
func LINEARREG_INTERCEPT() {

}

/*
LINEARREG_SLOPE - Linear Regression Slope

Input:

Output:

名称: 线性回归斜率指标

简介: LINEARREG_SLOPE_todo2

算法: LINEARREG_SLOPE_todo3
*/
func LINEARREG_SLOPE() {

}

/*
LN - Vector Log Natural

Input:

Output:

名称: 自然对数

简介: 自然对数

算法: 自然对数
*/
func LN() {

}

/*
LOG10 - Vector Log10

Input:

Output:

名称: 对数函数log

简介: 对数函数log

算法: 对数函数log
*/
func LOG10() {

}

/*
MA - All Moving Average

Input:

Output:

名称: 移动平均线

简介: 移动平均线，Moving Average，简称MA，原本的意思是移动平均，由于我们将其制作成线形，所以一般称之为移动平均线，简称均线。它是将某一段时间的收盘价之和除以该周期。 比如日线MA5指5天内的收盘价除以5 。

算法: MA_todo3
*/
func MA() {

}

/*
MACD - Moving Average Convergence/Divergence

Input:

Output:

名称: 平滑异同移动平均线

简介: 利用收盘价的短期（常用为12日）指数移动平均线与长期（常用为26日）指数移动平均线之间的聚合与分离状况，对买进、卖出时机作出研判的技术指标。

算法: MACD_todo3
*/
func MACD() {

}

/*
MACDEXT - MACD with controllable MA type

Input:

Output:

名称: 平滑异同移动平均线(可控制移动平均算法)

简介: 同MACD函数(固定使用EMA作为matype),并提供参数控制计算DIF, DEM时使用的移动平均算法。计算DIF时使用fastmatype与slowmatype，计算DEM时使用signalmatype，Histogram = DIF - DEM。matype参数详见talib.MA_Type与Overlap Studies Functions 重叠研究指标文档。

算法: MACDEXT_todo3
*/
func MACDEXT() {

}

/*
MACDFIX - Moving Average Convergence/Divergence Fix 12/26

Input:

Output:

名称：平滑异同移动平均线(固定快慢均线周期为12/26)

简介：同MACD函数, 固定快均线周期fastperiod=12, 慢均线周期slowperiod=26.

算法: MACDFIX_todo3
*/
func MACDFIX() {

}

/*
MAMA - MESA Adaptive Moving Average

Input:

Output:

名称: MAMA_todo1

简介: MAMA_todo2

算法: MAMA_todo3
*/
func MAMA() {

}

/*
MAVP - todo doc MAVP

Input:

Output:

名称: MAVP_todo1

简介: MAVP_todo2

算法: MAVP_todo3
*/
func MAVP() {

}

/*
MEDPRICE - Median Price

Input:

Output:

名称: 中位数价格

简介: MEDPRICE_todo2

算法: MEDPRICE_todo3
*/
func MEDPRICE() {

}

/*
MFI - Money Flow Index

Input:

Output:

名称：资金流量指标

简介：属于量价类指标，反映市场的运行趋势

算法: MFI_todo3
*/
func MFI() {

}

/*
MIDPOINT - MidPoint over period

Input:

Output:

名称: MIDPOINT_todo1

简介: MIDPOINT_todo2

算法: MIDPOINT_todo3
*/
func MIDPOINT() {

}

/*
MIDPRICE - Midpoint Price over period

Input:

Output:

名称: MIDPRICE_todo1

简介: MIDPRICE_todo2

算法: MIDPRICE_todo3
*/
func MIDPRICE() {

}

/*
MINUS_DI - Minus Directional Indicator

Input:

Output:

名称：下升动向值

简介：通过分析股票价格在涨跌过程中买卖双方力量均衡点的变化情况，即多空双方的力量的变化受价格波动的影响而发生由均衡到失衡的循环过程，从而提供对趋势判断依据的一种技术指标。

算法: MINUS_DI_todo3
*/
func MINUS_DI() {

}

/*
MINUS_DM - Minus Directional Movement

Input:

Output:

名称： 上升动向值 DMI中的DM代表正趋向变动值即上升动向值

简介：通过分析股票价格在涨跌过程中买卖双方力量均衡点的变化情况，即多空双方的力量的变化受价格波动的影响而发生由均衡到失衡的循环过程，从而提供对趋势判断依据的一种技术指标。

算法: MINUS_DM_todo3
*/
func MINUS_DM() {

}

/*
MOM - Momentum

Input:

Output:

名称： 上升动向值

简介：投资学中意思为续航，指股票(或经济指数)持续增长的能力。研究发现，赢家组合在牛市中存在着正的动量效应，输家组合在熊市中存在着负的动量效应。

算法: MOM_todo3
*/
func MOM() {

}

/*
NATR - Normalized Average True Range

Input:

Output:

名称：归一化波动幅度均值

简介：归一化波动幅度均值（NATR)是

算法: NATR_todo3
*/
func NATR() {

}

/*
NVI - todo doc NVI

Input:

Output:

名称: NVI_todo1

简介: NVI_todo2

算法: NVI_todo3
*/
func NVI() {

}

/*
OBV - On Balance Volume

Input:

Output:

名称：On Balance Volume 能量潮

简介：Joe Granville提出，通过统计成交量变动的趋势推测股价趋势

算法：多空比率净额= [（收盘价－最低价）－（最高价-收盘价）] ÷（ 最高价－最低价）×成交量 以某日为基期，逐日累计每日上市股票总成交量，若隔日指数或股票上涨 ，则基期OBV加上本日成交量为本日OBV。隔日指数或股票下跌， 则基期OBV减去本日成交量为本日OBV

研判：

	1、以“N”字型为波动单位，一浪高于一浪称“上升潮”，下跌称“跌潮”；上升潮买进，跌潮卖出
	2、须配合K线图走势
	3、用多空比率净额法进行修正，但不知TA-Lib采用哪种方法
*/
func OBV() {

}

/*
PLUS_DI - Plus Directional Indicator

Input:

Output:

名称: PLUS_DI_todo1

简介: PLUS_DI_todo2

算法: PLUS_DI_todo3
*/
func PLUS_DI() {

}

/*
PLUS_DM - Plus Directional Movement

Input:

Output:

名称: PLUS_DM_todo1

简介: PLUS_DM_todo2

算法: PLUS_DM_todo3
*/
func PLUS_DM() {

}

/*
PPO - Percentage Price Oscillator

Input:

Output:

函数名：PPO 名称： 价格震荡百分比指数

简介：价格震荡百分比指标（PPO）是一个和MACD指标非常接近的指标。PPO标准设定和MACD设定非常相似：12,26,9和PPO，和MACD一样说明了两条移动平均线的差距，但是它们有一个差别是PPO是用百分比说明。

算法: PPO_todo3
*/
func PPO() {

}

/*
PVI - todo doc PVI

Input:

Output:

名称: PVI_todo1

简介: PVI_todo2

算法: PVI_todo3
*/
func PVI() {

}

/*
ROC - Rate of change : ((price/prevPrice)-1)*100

Input:

Output:

名称： 变动率指标

简介：ROC是由当天的股价与一定的天数之前的某一天股价比较，其变动速度的大小,来反映股票市变动的快慢程度

算法: ROC_todo3
*/
func ROC() {

}

/*
ROCP - Rate of change Percentage: (price-prevPrice)/prevPrice

Input:

Output:

名称: ROCP_todo1

简介: ROCP_todo2

算法: ROCP_todo3
*/
func ROCP() {

}

/*
ROCR - Rate of change ratio: (price/prevPrice)

Input:

Output:

名称: ROCR_todo1

简介: ROCR_todo2

算法: ROCR_todo3
*/
func ROCR() {

}

/*
ROCR100 - Rate of change ratio 100 scale: (price/prevPrice)*100

Input:

Output:

名称: ROCR100_todo1

简介: ROCR100_todo2

算法: ROCR100_todo3
*/
func ROCR100() {

}

/*
RSI - Relative Strength Index

Input:

Output:

名称：相对强弱指数

简介：是通过比较一段时期内的平均收盘涨数和平均收盘跌数来分析市场买沽盘的意向和实力，从而作出未来市场的走势。

算法: RSI_todo3
*/
func RSI() {

}

/*
SAR - Parabolic SAR

Input:

Output:

名称: 抛物线指标

简介: 抛物线转向也称停损点转向，是利用抛物线方式，随时调整停损点位置以观察买卖点。由于停损点（又称转向点SAR）以弧形的方式移动，故称之为抛物线转向指标 。

算法: SAR_todo3
*/
func SAR() {

}

/*
SAREXT - Parabolic SAR - Extended

Input:

Output:

名称: SAREXT_todo1

简介: SAREXT_todo2

算法: SAREXT_todo3
*/
func SAREXT() {

}

/*
SIN - Vector Trigonometric Sin

Input:

Output:

名称: 正弦函数，三角函数

简介: 正弦函数，三角函数

算法: 正弦函数，三角函数
*/
func SIN() {

}

/*
SINH - Vector Trigonometric Sinh

Input:

Output:

名称: 双曲正弦函数，三角函数

简介: 双曲正弦函数，三角函数

算法: 双曲正弦函数，三角函数
*/
func SINH() {

}

/*
SMA - Simple Moving Average

Input:

Output:

名称: 简单移动平均线

简介: 移动平均线，Moving Average，简称MA，原本的意思是移动平均，由于我们将其制作成线形，所以一般称之为移动平均线，简称均线。它是将某一段时间的收盘价之和除以该周期。 比如日线MA5指5天内的收盘价除以5 。

算法: SMA_todo3
*/
func SMA() {

}

/*
SQRT - Vector Square Root

Input:

Output:

名称: 非负实数的平方根

简介: 非负实数的平方根

算法: 非负实数的平方根
*/
func SQRT() {

}

/*
STDDEV - Standard Deviation

Input:

Output:

名称: 标准偏差

简介: 种量度数据分布的分散程度之标准，用以衡量数据值偏离算术平均值的程度。标准偏差越小，这些值偏离平均值就越少，反之亦然。标准偏差的大小可通过标准偏差与平均值的倍率关系来衡量。

算法: STDDEV_todo3
*/
func STDDEV() {

}

/*
STOCH - Stochastic

Input:

Output:

名称: 随机指标,俗称KD

简介: STOCH_todo2

算法: STOCH_todo3
*/
func STOCH() {

}

/*
STOCHF - Stochastic Fast

Input:

Output:

名称: STOCHF_todo1

简介: STOCHF_todo2

算法: STOCHF_todo3
*/
func STOCHF() {

}

/*
STOCHRSI - Stochastic Relative Strength Index

Input:

Output:

名称: STOCHRSI_todo1

简介: STOCHRSI_todo2

算法: STOCHRSI_todo3
*/
func STOCHRSI() {

}

/*
T3 - Triple Exponential Moving Average (T3)

Input:

Output:

名称: 三重指数移动平均线

简介: TRIX长线操作时采用本指标的讯号，长时间按照本指标讯号交易，获利百分比大于损失百分比，利润相当可观。 比如日线MA5指5天内的收盘价除以5 。

算法: T3_todo3
*/
func T3() {

}

/*
TAN - Vector Trigonometric Tan

Input:

Output:

名称: 正切函数，三角函数

简介: 正切函数，三角函数

算法: 正切函数，三角函数
*/
func TAN() {

}

/*
TANH - Vector Trigonometric Tanh

Input:

Output:

名称: 双曲正切函数，三角函数

简介: 双曲正切函数，三角函数

算法: 双曲正切函数，三角函数
*/
func TANH() {

}

/*
TEMA - Triple Exponential Moving Average

Input:

Output:

名称: TEMA_todo1

简介: TEMA_todo2

算法: TEMA_todo3
*/
func TEMA() {

}

/*
TRANGE - True Range

Input:

Output:

名称: 真正的范围

简介: TRANGE_todo2

算法: TRANGE_todo3
*/
func TRANGE() {

}

/*
TRIMA - Triangular Moving Average

Input:

Output:

名称: TRIMA_todo1

简介: TRIMA_todo2

算法: TRIMA_todo3
*/
func TRIMA() {

}

/*
TRIX - 1-day Rate-Of-Change (ROC) of a Triple Smooth EMA

Input:

Output:

名称: TRIX_todo1

简介: TRIX_todo2

算法: TRIX_todo3
*/
func TRIX() {

}

/*
TSF - Time Series Forecast

Input:

Output:

名称: 时间序列预测

简介: 一种历史资料延伸预测，也称历史引伸预测法。是以时间数列所能反映的社会经济现象的发展过程和规律性，进行引伸外推，预测其发展趋势的方法

算法: TSF_todo3
*/
func TSF() {

}

/*
TYPPRICE - Typical Price

Input:

Output:

名称: 代表性价格

简介: TYPPRICE_todo2

算法: TYPPRICE_todo3
*/
func TYPPRICE() {

}

/*
ULTOSC - Ultimate Oscillator

Input:

Output:

名称：终极波动指标

简介：UOS是一种多方位功能的指标，除了趋势确认及超买超卖方面的作用之外，它的“突破”讯号不仅可以提供最适当的交易时机之外，更可以进一步加强指标的可靠度。

算法: ULTOSC_todo3
*/
func ULTOSC() {

}

/*
VAR - Variance

Input:

Output:

名称: 方差

简介: 方差用来计算每一个变量（观察值）与总体均数之间的差异。为避免出现离均差总和为零，离均差平方和受样本含量的影响，统计学采用平均离均差平方和来描述变量的变异程度

算法: VAR_todo3
*/
func VAR() {

}

/*
WCLPRICE - Weighted Close Price

Input:

Output:

名称: 加权收盘价

简介: WCLPRICE_todo2

算法: WCLPRICE_todo3
*/
func WCLPRICE() {

}

/*
WILLR - Williams' %R

Input:

Output:

名称：威廉指标

简介：WMS表示的是市场处于超买还是超卖状态。股票投资分析方法主要有如下三种：基本分析、技术分析、演化分析。在实际应用中，它们既相互联系，又有重要区别。

算法: WILLR_todo3
*/
func WILLR() {

}

/*
WMA - Weighted Moving Avera

Input:

Output:

名称: 加权移动平均线

简介: 移动加权平均法是指以每次进货的成本加上原有库存存货的成本，除以每次进货数量与原有库存存货的数量之和，据以计算加权平均单位成本，以此为基础计算当月发出存货的成本和期末存货的成本的一种方法。

算法: WMA_todo3
*/
func WMA() {

}
