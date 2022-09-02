package talib_test

import (
	"fmt"
	"github.com/bensema/talib"
)

func ExampleAdd() {

	a := []float64{1.1, 2.2, 3.3}
	b := []float64{1.1, 2.2, 3.3}

	out := talib.Add(a, b)

	fmt.Println(out)
	// Output: [2.2 4.4 6.6]
}

func ExampleDiv() {

	a := []float64{1.1, 2.22, 3.33}
	b := []float64{1.1, 2.2, 3.3}

	out := talib.Div(a, b)

	fmt.Println(out)
	// Output: [1 1.009090909090909 1.0090909090909093]
}

func ExampleMult() {

	a := []float64{3, 4, 1.2}
	b := []float64{3, 5, 5}

	out := talib.Mult(a, b)

	fmt.Println(out)
	// Output: [9 20 6]
}

func ExampleSub() {

	a := []float64{3, 4, 1.2}
	b := []float64{3, 5, 5}

	out := talib.Sub(a, b)

	fmt.Println(out)
	// Output: [0 -1 -3.8]
}

func ExampleSum() {

	in := []float64{3, 4, 1.2}

	out := talib.Sum(in, 2)

	fmt.Println(out)
	// Output: [0 7 5.2]
}

func ExampleMax() {

	a := []float64{1, 2, 3, 4, 5, 6, 7, 11, 9, 10}

	out := talib.Max(a, 3)

	fmt.Println(out)
	// Output: [0 0 3 4 5 6 7 11 11 11]
}

func ExampleMaxIndex() {

	a := []float64{1, 2, 3, 4, 5, 6, 7, 11, 9, 10}

	out := talib.MaxIndex(a, 3)

	fmt.Println(out)
	// Output: [0 0 2 3 4 5 6 7 7 7]
}

func ExampleMin() {

	a := []float64{1, 2, 3, 4, 5, 6, 7, 11, 9, 10}

	out := talib.Min(a, 3)

	fmt.Println(out)
	// Output: [0 0 1 2 3 4 5 6 7 9]
}

func ExampleMinIndex() {

	a := []float64{1, 2, 3, 4, 5, 6, 7, 11, 9, 10}

	out := talib.MinIndex(a, 3)

	fmt.Println(out)
	// Output: [0 0 0 1 2 3 4 5 6 8]
}

func ExampleMinMax() {

	a := []float64{1, 2, 3, 4, 5, 6, 7, 11, 9, 10}
	minOut, maxOut := talib.MinMax(a, 3)
	fmt.Println(minOut)
	fmt.Println(maxOut)
	// Output: [0 0 1 2 3 4 5 6 7 9]
	//[0 0 3 4 5 6 7 11 11 11]
}

func ExampleMinMaxIndex() {

	a := []float64{1, 2, 3, 4, 5, 6, 7, 11, 9, 10}

	minOut, maxOut := talib.MinMaxIndex(a, 3)

	fmt.Println(minOut)
	fmt.Println(maxOut)
	// Output: [0 0 0 1 2 3 4 5 6 8]
	//[0 0 2 3 4 5 6 7 7 7]
}

func ExampleEMA() {

	a := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	out := talib.Ema(a, 3)
	fmt.Println(out)
	// Output: [0 0 2 3 4 5 6 7 8 9]
}
