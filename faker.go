package fastfaker

import (
	"math/rand"
	"sync"
	"time"
)

// Seed uses the provided seed value to initialize the generator to a deterministic state.
func (f *Faker) Seed(seed int64) {
	if f.safe {
		f.mutex.Lock()
		defer f.mutex.Unlock()
	}
	f.src.Seed(seed)
}

// Global is a singleton, safe for concurrency instance of a Faker.
var Global *Faker

// Faker is the main strut that encapsulate all the package functionality.
type Faker struct {
	//each faker has its own entropy source
	src *rand.Rand
	//use the mutex only if this is true
	safe bool
	//the mutex has to protect all calls to the src
	mutex sync.Mutex
}

// NewSafeFaker creates a new Faker instance that can be used in a concurrent/parallel env.
//Its state is non-deterministic (seeded with a random value) but you can use the .Seed() method.
//It uses a mutex, so the performance may be affected.
func NewSafeFaker() *Faker {
	return &Faker{
		src:  rand.New(rand.NewSource(time.Now().UTC().UnixNano())),
		safe: true,
	}
}

// NewFastFaker generates a new Faker instance that CANNOT be used from multiple goroutines in the same time.
//Its state is non-deterministic (seeded with a random value) but you can use the .Seed() method.
//Use this faker when you want better performance in a multi-thread scenario.
func NewFastFaker() *Faker {
	return &Faker{
		src:  rand.New(rand.NewSource(time.Now().UTC().UnixNano())),
		safe: false,
	}
}

func init() {
	Global = NewSafeFaker()
}
