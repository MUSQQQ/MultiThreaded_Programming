// f(x) = 3x^3 + cos(7x) - ln(2x)
//[1;40]
//metoda prostokatow z dx = 0.0001

package main

import (
    "fmt"
    "time"
    "math"
    "sync"
)

var (
    dx float64 = 0.00001
    fx float64 = 0.0
    mutex = sync.Mutex{}
    wg sync.WaitGroup
)

//variables declaration
func integral(minimal float64) {
    tmp := 0.0
    wg.Add(1)
    time.Sleep(5 * time.Second)
    for i := minimal; i < minimal + 10; i += dx {
        tmp += (3 * math.Pow(i, 3) + math.Cos(7 * i) - math.Log(2 * i))*dx
    }
    mutex.Lock()
    fx += tmp
    mutex.Unlock()
    wg.Done()

}


func main() {
    time0 := time.Now()
    wg.Add(0)
    go integral(1)
    go integral(10)
    go integral(20)
    go integral(30)
    wg.Wait()
    time1 := time.Now()
    fmt.Printf("Given integral in the section <0;40> is equal to:\n")
    fmt.Printf("%v\n",fx)
    fmt.Printf("Program took %v nanoseconds to run.\n",time1.Sub(time0).Nanoseconds())
}