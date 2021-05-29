package main

import (
	"fmt"
	"math/rand"
	//"math/rand"
	"sync"
	"time"
)

//number of philosophers, forks and seats
const numOfPhilosophers = 5 //constant needed to be used while declaring an array of philosophers
const k = 5                 //number of books
var r = rand.New(rand.NewSource(1))

type Fork struct{ sync.Mutex }
type Book struct{ sync.Mutex }

type Philosopher struct {
	Num int
	Mut sync.Mutex
}

func (p *Philosopher) Think() {
	fmt.Printf("Philosopher %d is thinking\n", p.Num)

	time.Sleep(time.Duration(r.Int31()) * 3)
	fmt.Printf("Philosopher %d got hungry\n", p.Num)
	time.Sleep(time.Duration(r.Int31()) * 1)
}

func (p *Philosopher) Eat(bookNum int) {
	fmt.Printf("Philosopher %d is eating and is reading book number %d \n", p.Num, bookNum)
	time.Sleep(time.Duration(r.Int31()) * 5)
	fmt.Printf("Philosopher %d finished eating and finished reading book number %d\n", p.Num, bookNum)
}
func (p *Philosopher) Action(forks []Fork, seat int, books []Book, bookNum int) {
	//for{} works like "while(true)"
	for {
		p.Think()
		books[bookNum].Lock()
		if seat == 0 {
			forks[numOfPhilosophers-1].Lock()
		} else {
			forks[seat-1].Lock()
		}
		if seat == numOfPhilosophers-1 {
			forks[0].Lock()
		} else {
			forks[seat+1].Lock()
		}
		p.Eat(bookNum)
		if seat == 0 {
			forks[numOfPhilosophers-1].Unlock()
		} else {
			forks[seat-1].Unlock()
		}
		if seat == numOfPhilosophers-1 {
			forks[0].Unlock()
		} else {
			forks[seat+1].Unlock()
		}
		books[bookNum].Unlock()
	}
}

func main() {

	philosophers := [numOfPhilosophers]Philosopher{}
	forks := [numOfPhilosophers]Fork{}
	books := [k]Book{}
	for i := 0; i < numOfPhilosophers; i++ {
		philosophers[i] = Philosopher{Num: i}
	}

	var wg sync.WaitGroup
	wg.Add(1)
	for i := 0; i < numOfPhilosophers; i++ {
		//przyjmuję, że wynik funkcji rand.Intn() to "numer" miejsca wybranego przez filozofa
		go philosophers[i].Action(forks[:], rand.Intn(numOfPhilosophers), books[:], rand.Intn(k))
	}

	wg.Wait()
	//time.Sleep(30 * time.Second)
	//fmt.Println(rand.Intn(6))

	fmt.Println("Program finished")
}
