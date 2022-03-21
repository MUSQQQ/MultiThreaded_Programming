// f(x) = 3x^3 + cos(7x) - ln(2x)
//[1;40]
//metoda prostokatow z dx = 0.00001

package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

var (
	dx     float64 = 0.00001
	fx     float64 = 0.0
	tmp    float64 = 0.0
	wg     sync.WaitGroup
	values []float64 = []float64{1, 10, 20, 30}
)

//variables declaration
func integral(minimal float64) {
	wg.Add(1)
	time.Sleep(10 * time.Second)
	tmp = 0
	for i := minimal; i < minimal+10; i += dx {
		fx += (3*math.Pow(i, 3) + math.Cos(7*i) - math.Log(2*i)) * dx
	}
	wg.Done()
}

func main() {
	time0 := time.Now()
	wg.Add(0)
	for i := 0; i < 4; i++ {
		go integral(values[i])
	}
	wg.Wait()
	time1 := time.Now()
	fmt.Println("Disclaimer: this program does not calculate the integral correctly due")
	fmt.Printf("to lack of control over access to the shared variable\n")
	fmt.Printf("Given integral on the interval <1;40> is equal to:\n%v\n", fx)
	fmt.Printf("Program took %v nanoseconds to run.\n", time1.Sub(time0).Nanoseconds())
}
