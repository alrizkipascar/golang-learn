package main

// Compulsory package,
// the only one generating an executable

import (
	"fmt"  // Format library, including I/O methods
	"sync" // Library providing basic synchronization primitives
	"time" // Package providing functionality for handling time.
)

// Host structure
// The host is the one giving permission to eat.
// In this implementation is the arbitrator choosing the picking order of
// chopsticks.
// It has two attributes: a channel to communicate with philosophers and the
// channel buffer.
// It has just one method: Authorise.
type Host struct {
	channel  chan *Philosopher
	chanBuff int
}

// Authorise method
// Host method to determine the chopsticks picking order.
// The function transmit on the channel waiting for some philosopher to receive.
func (host *Host) Authorise(wg *sync.WaitGroup) {
	for {
		if len(host.channel) == host.chanBuff {
			<-host.channel
			<-host.channel

			// Add delay
			time.Sleep(500 * time.Millisecond)
		}
	}
}

// Chopstick structure
// It is a wrapper for a mutex, since picking a chopstick
// would cause some semaphore locking.
type Chopstick struct {
	sync.Mutex
}

// Philosopher structure
// It contains five attributes, the name (only for aesthetic purposes), the id
// (position at table), left and right chopsticks and a counter, keepnig track
// of the number of times the philosopher has already eaten.
type Philosopher struct {
	name            string
	id              int
	leftCS, rightCS *Chopstick
	eatCount        int
}

// Eat method for philosophers
// It is written such that it has to be called in a goroutine.
// First, it checks whether the host gave permission to eat.
// It makes the philosopher grab the two chopsticks and increases the counter of
// eaten times.
func (philosopher *Philosopher) Eat(host *Host, wg *sync.WaitGroup) {
	for i := 0; i < 3; i++ {
		host.channel <- philosopher

		if philosopher.eatCount < 3 {
			philosopher.leftCS.Lock()
			philosopher.rightCS.Lock()

			fmt.Printf("%v starting to eat \n", philosopher.name)
			philosopher.eatCount++ // Increase counter
			fmt.Printf("%v finishing eating \n", philosopher.name)

			philosopher.leftCS.Unlock()
			philosopher.rightCS.Unlock()

			wg.Done()
		}
	}
}

// A collection of philosophers names
var names []string = []string{"Kant", "Heidegger", "Wittgenstein",
	"Locke", "Descartes", "Newton", "Hume", "Leibniz"}

// The host channel buffer
var chanBuff int = 2

// The max number of times philosophers are allowed to eat
var maxMeals int = 3

// Script to model the philosophers' dining problem.
// Here, we implemented the arbitrator solution, with a central authority,
// the host and five hungry philosophers.
func main() {
	var wg sync.WaitGroup // WaitGroup object, to wait for routines to complete
	// tasks
	var host Host
	host.chanBuff = chanBuff
	host.channel = make(chan *Philosopher, chanBuff)

	chopsticks := make([]*Chopstick, len(names))
	for i := 0; i < len(names); i++ {
		chopsticks[i] = new(Chopstick)
	}

	philosophers := make([]*Philosopher, len(names))
	for i := 0; i < len(names); i++ {
		philosophers[i] = &Philosopher{names[i], i + 1, chopsticks[i],
			chopsticks[(i+1)%len(names)], 0}
	}
	fmt.Printf("There are %v philosophers sitting at a table.\n",
		len(philosophers))

	// Add semaphore counter values.
	wg.Add(maxMeals * len(philosophers))

	// Authorisation of the host
	go host.Authorise(&wg)

	// Each Philosopher checks the authorisation and eats if he is allowed to.
	for i := 0; i < len(philosophers); i++ {
		go philosophers[i].Eat(&host, &wg)
	}

	// Be sure the main goroutine waits the others to complete.
	wg.Wait()

}
