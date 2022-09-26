package main

import (
	"fmt"
	"sync"
	"time"
)

type fork struct{ sync.Mutex }

var host struct{ sync.Mutex }

type philosopher struct {
	id                  int
	leftFork, rightFork *fork
	eaten				int
}

// Goes from thinking to hungry to eating and done eating then starts over.
// Adapt the pause values to increased or decrease contentions
// around the forks.
func (p philosopher) eat() {
	
	for j := 0; j < 2; j++ {
		host.Lock()
		p.leftFork.Lock()
		p.rightFork.Lock()
		time.Sleep(time.Second)
		fmt.Println("Filosofo",p.id+1,"comiendo")
		time.Sleep(time.Second)

		p.rightFork.Unlock()
		p.leftFork.Unlock()
		host.Unlock()

		fmt.Println("Filosofo",p.id+1,"termino de comer")
		p.eaten+=1
		time.Sleep(time.Second)
	}
	eatWgroup.Done()
}





var eatWgroup sync.WaitGroup

func main() {
	// How many philosophers and forks

	count := 5

	

	forks := make([]*fork, count)
	for i := 0; i < count; i++ {
		forks[i] = new(fork)
	}

	// Create philospoher, assign them 2 forks and send them to the dining table
	philosophers := make([]*philosopher, count)
	for i := 0; i < count; i++ {
		philosophers[i] = &philosopher{
			id: i, leftFork: forks[i], rightFork: forks[(i+1)%count], eaten: 0}
		eatWgroup.Add(1)
		go philosophers[i].eat()

	}
	eatWgroup.Wait()

}

