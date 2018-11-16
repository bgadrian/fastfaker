package fastfaker

import (
	"github.com/bgadrian/fastfaker/randomizer"
	"github.com/bgadrian/fastfaker/subtext"
)

// Global is a singleton, safe for concurrency instance of a Faker.
var Global *Faker

// Faker is the main strut that encapsulate all the package functionality.
type Faker struct {
	*randomizer.RandWrapper
	*subtext.FakerText
}

// NewSafeFaker creates a new Faker instance that can be used in a concurrent/parallel env.
//Its state is non-deterministic (seeded with a random value) but you can use the .Seed() method.
//It uses a mutex, so the performance may be affected.
func NewSafeFaker() *Faker {
	result := &Faker{}
	result.RandWrapper = randomizer.NewRandWrapper(true, nil).(*randomizer.RandWrapper)
	result.FakerText = subtext.New(result.RandWrapper)
	return result
}

// NewFastFaker generates a new Faker instance that CANNOT be used from multiple goroutines in the same time.
//Its state is non-deterministic (seeded with a random value) but you can use the .Seed() method.
//Use this faker when you want better performance in a multi-thread scenario.
func NewFastFaker() *Faker {
	result := &Faker{}
	result.RandWrapper = randomizer.NewRandWrapper(false, nil).(*randomizer.RandWrapper)
	return result
}

func init() {
	Global = NewSafeFaker()
}
