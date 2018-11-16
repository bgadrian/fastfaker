package randomizer

import (
	"math"
	"math/rand"
	"sync"
	"time"
)

// RandWrapper encapsulates a rand.Rand and adds more features. Implements Randomizer.
// The only way to make it concurrent safe is to use NewRandWrapper(true, ...)
// Do not use this struct directly to avoid braking changes, use the Randomizer interface!
type RandWrapper struct {
	//each faker has its own entropy source
	src *rand.Rand
	//use the mutex only if this is true
	safe bool
	//the mutex has to protect all calls to the src
	mutex sync.Mutex
}

// NewRandWrapper returns a wrapper of a rand.Rand with extra features.
//All functions are concurrent safe only if the param concurrentSafe==true
//If source is nil a new source will be generated with a random seed.
func NewRandWrapper(concurrentSafe bool, source rand.Source) Randomizer {
	result := &RandWrapper{}
	result.safe = concurrentSafe
	if source == nil {
		result.src = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	} else {
		result.src = rand.New(source)
	}
	return result
}

// Seed uses the provided seed value to initialize the generator to a deterministic state.
func (f *RandWrapper) Seed(seed int64) {
	if f.safe {
		f.mutex.Lock()
		defer f.mutex.Unlock()
	}
	f.src.Seed(seed)
}

// Number will generate a random number between given min And max
func (f *RandWrapper) Number(min int, max int) int {
	if min == max {
		return min
	}
	return f.Intn((max+1)-min) + min
}

// Uint8 will generate a random uint8 value
func (f *RandWrapper) Uint8() uint8 {
	return uint8(f.Number(0, math.MaxUint8))
}

// Uint16 will generate a random uint16 value
func (f *RandWrapper) Uint16() uint16 {
	return uint16(f.Number(0, math.MaxUint16))
}

// Uint32 will generate a random uint32 value
func (f *RandWrapper) Uint32() uint32 {
	return uint32(f.Number(0, math.MaxInt32))
}

// Uint64 will generate a random uint64 value
func (f *RandWrapper) Uint64() uint64 {
	if f.safe {
		f.mutex.Lock()
		defer f.mutex.Unlock()
	}
	return uint64(f.src.Int63n(math.MaxInt64))
}

// Int8 will generate a random Int8 value
func (f *RandWrapper) Int8() int8 {
	return int8(f.Number(math.MinInt8, math.MaxInt8))
}

// Int16 will generate a random int16 value
func (f *RandWrapper) Int16() int16 {
	return int16(f.Number(math.MinInt16, math.MaxInt16))
}

// Int32 will generate a random int32 value  [math.MinInt32, math.MaxInt32)
func (f *RandWrapper) Int32() int32 {
	return int32(f.Number(math.MinInt32, math.MaxInt32))
}

// Int64 will generate a random int64 value [math.MinInt64, math.MaxInt64)
func (f *RandWrapper) Int64() int64 {
	if f.safe {
		f.mutex.Lock()
		defer f.mutex.Unlock()
	}

	return f.src.Int63n(math.MaxInt64) + math.MinInt64
}

// Int64Positive returns a non-negative pseudo-random 63-bit integer as an int64 [0, math.Maxint64)
func (f *RandWrapper) Int64Positive() int64 {
	if f.safe {
		f.mutex.Lock()
		defer f.mutex.Unlock()
	}

	return f.src.Int63()
}

// Int32Range generates a random int32 between [min, max)
func (f *RandWrapper) Int32Range(min, max int32) int32 {
	if min >= max {
		return min
	}

	if f.safe {
		f.mutex.Lock()
		defer f.mutex.Unlock()
	}

	return f.src.Int31()%(max-min) + min
}

// Int64Range generates a random int64 between [min, max)
func (f *RandWrapper) Int64Range(min, max int64) int64 {
	if min >= max {
		return min
	}

	if f.safe {
		f.mutex.Lock()
		defer f.mutex.Unlock()
	}

	return f.src.Int63()%(max-min) + min
}

// Float32Range will generate a random float32 value between min and max
func (f *RandWrapper) Float32Range(min, max float32) float32 {
	if min >= max {
		return min
	}

	if f.safe {
		f.mutex.Lock()
		defer f.mutex.Unlock()
	}

	return f.src.Float32()*(max-min) + min
}

// Float64Range will generate a random float64 value between min and max
func (f *RandWrapper) Float64Range(min, max float64) float64 {
	if min >= max {
		return min
	}

	if f.safe {
		f.mutex.Lock()
		defer f.mutex.Unlock()
	}

	return f.src.Float64()*(max-min) + min
}

// ShuffleInts will randomize a slice of ints
func (f *RandWrapper) ShuffleInts(a []int) {
	for i := range a {
		j := f.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}

// Intn returns a random number between [0, n)
func (f *RandWrapper) Intn(n int) int {
	if f.safe {
		f.mutex.Lock()
		defer f.mutex.Unlock()
	}
	return f.src.Intn(n)
}

// Float32 returns a random positive float number [0.000000000001, math.MaxFloat32)
func (f *RandWrapper) Float32() float32 {
	return f.Float32Range(math.SmallestNonzeroFloat32, math.MaxFloat32)
}

// Float64 returns a random positive float number [0.000000000001, math.MaxFloat64)
func (f *RandWrapper) Float64() float64 {
	return f.Float64Range(math.SmallestNonzeroFloat64, math.MaxFloat64)
}

// Float64Unary returns, as a float64, a pseudo-random number in [0.0,1.0) from the default Source.
func (f *RandWrapper) Float64Unary() float64 {
	if f.safe {
		f.mutex.Lock()
		defer f.mutex.Unlock()
	}
	return f.src.Float64()
}

// Read generates len(p) random bytes from the default Source and writes them into p. It always returns len(p) and a nil error.
func (f *RandWrapper) Read(p []byte) (n int, err error) {
	if f.safe {
		f.mutex.Lock()
		defer f.mutex.Unlock()
	}
	return f.src.Read(p)
}
