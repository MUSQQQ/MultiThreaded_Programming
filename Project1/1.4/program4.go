// f(x) = 3x^3 + cos(7x) - ln(2x)
//[1;40]
//metoda prostokatow z dx = 0.00001

package main

import (
	"fmt"
	"math"
	"time"
)

var (
	dx     float64   = 0.00001
	fx     float64   = 0.0
	values []float64 = []float64{1, 10, 20, 30}
)

//variables declaration
func integral(minimal float64, results chan float64) {
	tmp := 0.0
	time.Sleep(10 * time.Second)
	for i := minimal; i < minimal+10; i += dx {
		tmp += (3*math.Pow(i, 3) + math.Cos(7*i) - math.Log(2*i)) * dx
	}
	results <- tmp
}

func main() {
	channel := make(chan float64)
	time0 := time.Now()
	for i := 0; i < 4; i++ {
		go integral(values[i], channel)
	}
	for i := 0; i < 4; i++ {
		fx += <-channel
	}
	time1 := time.Now()
	fmt.Printf("Given integral on the interval <1;40> is equal to:\n%v\n", fx)
	fmt.Printf("Program took %v nanoseconds to run.\n", time1.Sub(time0).Nanoseconds())
}
