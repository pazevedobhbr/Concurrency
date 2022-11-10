package main

import (
	"fmt"
	"sync"
	"time"
)

// The Dining Philosophers problem is well known in computer science circles.
// Five philosophers, numbered from 0 through 4, live in a house where the
// table is laid for them; each philosopher has their own place at the table.
// Their only difficulty – besides those of philosophy – is that the dish
// served is a very difficult kind of spaghetti which has to be eaten with
// two forks. There are two forks next to each plate, so that presents no
// difficulty. As a consequence, however, this means that no two neighbours
// may be eating simultaneously, since there are five philosophers and five forks.
//
// This is a simple implementation of Dijkstra's solution to the "Dining
// Philosophers" dilemma.

// Philosopher is a struct that represents a philosopher
type Philosopher struct {
	name      string
	rightFork int
	leftFork  int
}

// philosophers - List of all philosophers
var philosophers = []Philosopher{
	{name: "Plato", leftFork: 4, rightFork: 0},
	{name: "Socrates", leftFork: 0, rightFork: 1},
	{name: "Aristotle", leftFork: 1, rightFork: 2},
	{name: "Pascal", leftFork: 2, rightFork: 3},
	{name: "Locke", leftFork: 3, rightFork: 4},
}

var hunger = 3                  // How many times each philosopher will eat
var eatTime = 1 * time.Second   // How long it takes to eat
var thinkTime = 3 * time.Second // How long it takes to think
var sleepTime = 1 * time.Second // How long it takes to when printing thinks out

var orderMutex sync.Mutex  // a mutex for the slice orderFinished; part of the challenge!
var orderFinished []string //  the order in which the philosophers finish dining and leave; part of the challenge!

func main() {
	//print out a welcome message
	fmt.Println("Dining Philosophers problem")
	fmt.Println("---------------------------")
	fmt.Println("THe Table is empty.")

	time.Sleep(sleepTime)

	//start the meal
	dine()

	//print out a finished message
	fmt.Println("The table is empty.")

}

func dine() {
	//	eatTime = 0 * time.Second
	//	sleepTime = 0 * time.Second
	//	thinkTime = 0 * time.Second

	// wg is the WaitGroup that keeps track of how many philosophers are still at the table. When
	// it reaches zero, everyone is finished eating and has left. We add 5 (the number of philosophers) to this
	// wait group.
	wg := &sync.WaitGroup{}
	wg.Add(len(philosophers))

	// We want everyone to be seated before they start eating, so create a WaitGroup for that, and set it to 5.
	seated := &sync.WaitGroup{}
	seated.Add(len(philosophers))

	// forks is a map of all 5 forks. Forks are assigned using the fields leftFork and rightFork in the Philosopher
	// type. Each fork, then, can be found using the index (an integer), and each fork has a unique mutex.	var forks = make(map[int]*sync.Mutex)
	var forks = make(map[int]*sync.Mutex)
	for i := 0; i < len(philosophers); i++ {
		forks[i] = &sync.Mutex{}
	}

	//start the meal
	for i := 0; i < len(philosophers); i++ {
		//fire off a goroutine for the current philosopher
		go diningProblem(philosophers[i], wg, forks, seated)
	}
	// Wait for the philosophers to finish. This blocks until the wait group is 0.
	wg.Wait()
}

// diningProblem is the function fired off as a goroutine for each of our philosophers. It takes one
// philosopher, our WaitGroup to determine when everyone is done, a map containing the mutexes for every
// fork on the table, and a WaitGroup used to pause execution of every instance of this goroutine
// until everyone is seated at the table.
func diningProblem(philosopher Philosopher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup) {
	defer wg.Done()

	//seat the philosopher at the table
	fmt.Printf("%s is seated at the table.\n", philosopher.name)

	//Decrement the seated WaiGroup by one.
	seated.Done()

	// Wait for everyone to be seated before starting to eat.
	seated.Wait()

	//eat three times
	for i := hunger; i > 0; i-- {
		// Get a lock on the left and right forks. We have to choose the lower numbered fork first in order
		// to avoid a logical race condition, which is not detected by the -race flag in tests; if we don't do this,
		// we have the potential for a deadlock, since two philosophers will wait endlessly for the same fork.
		// Note that the goroutine will block (pause) until it gets a lock on both the right and left forks.
		//get the lock on both forks
		if philosopher.leftFork > philosopher.rightFork {
			forks[philosopher.rightFork].Lock()
			fmt.Printf("\t%s has picked up the right fork.\n", philosopher.name)
			forks[philosopher.leftFork].Lock()
			fmt.Printf("\t%s has picked up the left fork.\n", philosopher.name)
		} else {
			forks[philosopher.leftFork].Lock()
			fmt.Printf("\t%s has picked up the left fork.\n", philosopher.name)
			forks[philosopher.rightFork].Lock()
			fmt.Printf("\t%s has picked up the right fork.\n", philosopher.name)

		}
		// By the time we get to this line, the philosopher has a lock (mutex) on both forks.
		fmt.Printf("\t%s has both forks and is eating.\n", philosopher.name)
		time.Sleep(eatTime)

		// The philosopher starts to think, but does not drop the forks yet.
		fmt.Printf("\t%s is thinking.\n", philosopher.name)
		time.Sleep(thinkTime)

		// Unlock the mutexes for both forks.
		forks[philosopher.leftFork].Unlock()
		forks[philosopher.rightFork].Unlock()

		fmt.Printf("\t%s put the down the forks.\n", philosopher.name)
	}
	fmt.Println(philosopher.name, "is satisfied.")
	fmt.Println(philosopher.name, "left the table.")
	orderMutex.Lock()
	orderFinished = append(orderFinished, philosopher.name)
	orderMutex.Unlock()

}
