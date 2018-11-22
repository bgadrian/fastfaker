package producer

import (
	"fmt"
	"math/rand"
	"time"
)

// Dispatcher makes sure the users are spawned and killed accordnly to the TrafficPattern pattern
type Dispatcher struct {
	running    bool
	online     []User
	usersQueen UserGenerator
	//lock         sync.Mutex //for Stop and Ticker in the same time
}

// NewDispatcher is making sure that [traffic.Next()] users are running and in
// total generating [eventsPerUserMin] events.
// The Traffic.next()] is called each [intervalSize].
func NewDispatcher(queen UserGenerator) *Dispatcher {
	d := &Dispatcher{}
	d.usersQueen = queen
	return d
}

// Start ticking and adjusting the online users based on the traffic pattern
func (d *Dispatcher) Start(intervalSize time.Duration) {
	if d.running {
		return
	}

	ticker := time.NewTicker(intervalSize)
	d.running = true

	//create first users
	d.tick()

	go func() {
	exit:
		for range ticker.C {
			if d.running == false {
				break exit
			}
			//must block to avoid concurrent d.online updates
			d.tick()
		}
		ticker.Stop()
		d.online = nil
	}()
}

// Stop pauses the dispatcher, the online users will still SEND events, but their numbers will not fluctuate any more
func (d *Dispatcher) Stop() {
	d.running = false
}

func (d *Dispatcher) tick() {
	//between 5 and 25 users
	usersShould := rand.Intn(20) + 5
	usersNow := len(d.online)

	if usersNow == usersShould {
		return
	}

	diff := usersShould - usersNow
	if diff > 0 {
		//add more users
		for i := 0; i < diff; i++ {
			user := d.usersQueen()
			d.online = append(d.online, user)
			userId, _ := user.Property(KeyUserID)
			fmt.Printf("new user: %s\n", userId)
		}
		return
	}

	//remove random users
	for len(d.online) > usersShould {
		killIndex := rand.Intn(len(d.online))
		userId, _ := d.online[killIndex].Property(KeyUserID)
		d.online = append(d.online[:killIndex], d.online[killIndex+1:]...)
		fmt.Printf("kill user: %s\n", userId)
	}
}
