package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic" //użycie funkcji z pakietu atomic zapewnia że będą one atomowe
	"time"
)

var r = rand.New(rand.NewSource(1))

type SeqLock struct {
	Counter    int32 //zwiększa się o 1 po zajęciu zamka przez pisarza i tuż przed zwolnieniem tego zamka (sequence number)
	sync.Mutex       //zamek tylko dla pisarzy
}

//funkcja odczytująca atomowo licznik seqlocka
func (seq *SeqLock) rdRead() int32 {
	return atomic.LoadInt32(&seq.Counter)
}

//funkcja sprawdzająca czy licznik nie zmienił się w trakcie odczytu danych lub czy nie jest nieparzysty(czy aktualnie pisarz nie działa na danych)
func (seq *SeqLock) rdAgain(val int32) bool {
	return (atomic.LoadInt32(&seq.Counter)&1) != 0 || val != seq.Counter
}

func (seq *SeqLock) wrLock() {
	seq.Lock()
	atomic.AddInt32(&seq.Counter, 1) //counter staje sie nieparzysty gdy pisarz zaczyna działać
}

func (seq *SeqLock) wrUnlock() {
	atomic.AddInt32(&seq.Counter, 1) //counter staje się z powrotem parzysty po skończeniu pracy na danych
	seq.Unlock()
}

//funkcja przedstawiająca cały proces odczytywania danych
func (seq *SeqLock) ReadingData(wg *sync.WaitGroup) {

	tmp := int32(0)
	for {
		time.Sleep(time.Duration(r.Int31()) * 1)
		tmp = seq.rdRead()
		/*

			odczytywanie danych

		*/
		if !seq.rdAgain(tmp) {

			fmt.Printf("odczyt licznika po wykryciu braku zmiany/nieparzystości: %d\n", tmp)
			break
		}
	}
	defer wg.Done()
}

//funkcja przedstawiająca proces zapisu danych
func (seq *SeqLock) WritingData(wg *sync.WaitGroup) {

	time.Sleep(time.Duration(r.Int31()) * 1)
	seq.wrLock()
	/*

		zapisywanie danych

	*/
	fmt.Printf("zapis (w trakcie), licznik nieparzysty: %d\n", seq.Counter)
	seq.wrUnlock()
	defer wg.Done()
}

//Przykładowy program przedstawiający uproszczone działanie zaimplementowanego seqlocka
//bez modyfikowania danych, tylko ze zmienianiem wartości sequence number
func main() {
	SequenceLock := SeqLock{Counter: 0}

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		if rand.Intn(100) <= 80 {
			wg.Add(1)
			go SequenceLock.ReadingData(&wg)
		} else {
			wg.Add(1)
			go SequenceLock.WritingData(&wg)
		}
	}

	wg.Wait()
}
