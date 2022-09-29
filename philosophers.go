package main

import (
	"fmt"
	"sync"
	"time"
)


// Los tenedores eran un Lock que simula si esta siendo usado o no
type fork struct{ sync.Mutex }


// El host va a ser un Lock que simula si hay alguien comiendo o no
var host struct{ sync.Mutex }

// Un filosofo tendra 4 atributos, el id, los dos tenedores, el derecho y el izquierdo, y la variable eaten
// que representa cuantas veces ha comido
type philosopher struct {
	id                  int
	leftFork, rightFork *fork
	eaten				int
}


func (p philosopher) eat() {
	
	for j := 0; j < 2; j++ {
		// Primero el host bloquea para que nadie mas pueda estar comiendo a la vez
		host.Lock()
		// Se bloquean los dos cubiertos del filosofo
		p.leftFork.Lock()
		p.rightFork.Lock()
		time.Sleep(time.Second)
		fmt.Println("Filosofo",p.id+1,"comiendo")
		time.Sleep(time.Second)
		// Cuando el filosofo termina de comer, se desocupan los tenedores respectivos, y el host acepta que otra persona pueda comer
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
	// Inicializamos los tenedores y los filosofos

	count := 5

	

	forks := make([]*fork, count)
	for i := 0; i < count; i++ {
		forks[i] = new(fork)
	}

	// Para cada filosofo le asignamos los tenedores segun corresponda
	philosophers := make([]*philosopher, count)
	for i := 0; i < count; i++ {
		philosophers[i] = &philosopher{
			id: i, leftFork: forks[i], rightFork: forks[(i+1)%count], eaten: 0}
		eatWgroup.Add(1)
		go philosophers[i].eat()

	}
	eatWgroup.Wait()

}

