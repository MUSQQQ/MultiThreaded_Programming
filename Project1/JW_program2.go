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
    tmp float64 = 0.0
    mutex = sync.Mutex{}
    wg sync.WaitGroup
)

//variables declaration
func integral(minimal float64) {
    wg.Add(1)
    time.Sleep(5 * time.Second)
    tmp = 0
    for i := minimal; i < minimal + 10; i += dx {
        mutex.Lock()
        fx += (3 * math.Pow(i, 3) + math.Cos(7 * i) - math.Log(2 * i))*dx
        mutex.Unlock()
        //fmt.Printf("%v\n",fx)
    }
    //fmt.Printf("a")
    wg.Done()

}


func main() {
    time0 := time.Now()
    wg.Add(0)
    go integral(1)
    //time.Sleep(1000)
    go integral(10)
    //time.Sleep(1000)
    go integral(20)
    //time.Sleep(1000)
    go integral(30)
    wg.Wait()
    //time.Sleep(1000)
    time1 := time.Now()
    fmt.Printf("%v\n",fx)
    fmt.Printf("It took %v nanoseconds.\n",time1.Sub(time0).Nanoseconds())
}
