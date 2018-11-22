package producer

import (
	"errors"
	"math/rand"
	"time"
)

// KeyUserID contains the User property key for its ID
const KeyUserID = "userId"

// UserGenerator generates one new random users at each call.
type UserGenerator func() User

// User represents a fake user that generates analytics
type User interface {
	// Properties of an user like userID, OS ...
	Property(key string) (string, error)
	SetProperty(key, value string)
}

// NewSimpleUser must use the evg to generate eventsPerMin events
// and send them to the out channel, at random intervals.
func NewSimpleUser(eventsPerMin int, evg EventGenerator, out chan<- Event,
	minMs, maxMs int) User {
	u := &simpleuser{}
	u.evg = evg
	u.out = out
	u.props = GetRandomUserProperties()

	u.running = true

	go func() {
		var ms int
		if maxMs <= minMs {
			ms = minMs
		} else {
			ms = rand.Intn(maxMs-minMs) + minMs
		}
		amount := int(float64(eventsPerMin) * float64(ms/60000))
		ticker := time.NewTicker(time.Millisecond * time.Duration(ms))

		for {
			//TODO randomize the value of each interval,
			//simulating lag/network issues
			if u.running == false {
				break //do not make an extra tick if it was stopped
			}
			//blocking, no 2 ticks overlap
			u.tick(amount)
			<-ticker.C
		}
	}()
	return u
}

type simpleuser struct {
	evg     EventGenerator
	out     chan<- Event
	props   map[string]string
	running bool
	tickID  int
}

func (u *simpleuser) tick(amount int) {
	for _, ev := range u.evg(amount) {
		//add to each even unique user properties like
		//userID,OS ...
		for k, v := range u.props {
			ev.AddProperty(k, v)
		}
		//ev.AddProperty("userTick", strconv.Itoa(u.tickID))
		u.out <- ev
	}
	u.tickID++
}

func (u *simpleuser) SetProperty(key, value string) {
	u.props[key] = value
}
func (u *simpleuser) Property(key string) (string, error) {
	val, ok := u.props[key]
	if ok {
		return val, nil
	}
	return "", errors.New("not found")
}
