package gofakeit

import (
	"math"
)

// Number will generate a random number between given min And max
func (f *Faker) Number(min int, max int) int {
	return f.randIntRange(min, max)
}

// Uint8 will generate a random uint8 value
func (f *Faker) Uint8() uint8 {
	return uint8(f.randIntRange(0, math.MaxUint8))
}

// Uint16 will generate a random uint16 value
func (f *Faker) Uint16() uint16 {
	return uint16(f.randIntRange(0, math.MaxUint16))
}

// Uint32 will generate a random uint32 value
func (f *Faker) Uint32() uint32 {
	return uint32(f.randIntRange(0, math.MaxInt32))
}

// Uint64 will generate a random uint64 value
func (f *Faker) Uint64() uint64 {
	if f.safe {
		f.mutex.Lock()
		defer f.mutex.Unlock()
	}
	return uint64(f.src.Int63n(math.MaxInt64))
}

// Int8 will generate a random Int8 value
func (f *Faker) Int8() int8 {
	return int8(f.randIntRange(math.MinInt8, math.MaxInt8))
}

// Int16 will generate a random int16 value
func (f *Faker) Int16() int16 {
	return int16(f.randIntRange(math.MinInt16, math.MaxInt16))
}

// Int32 will generate a random int32 value  [math.MinInt32, math.MaxInt32)
func (f *Faker) Int32() int32 {
	return int32(f.randIntRange(math.MinInt32, math.MaxInt32))
}

// Int64 will generate a random int64 value [math.MinInt64, math.MaxInt64)
func (f *Faker) Int64() int64 {
	if f.safe {
		f.mutex.Lock()
		defer f.mutex.Unlock()
	}

	return f.src.Int63n(math.MaxInt64) + math.MinInt64
}

// Int32Range generates a random int32 between [min, max)
func (f *Faker) Int32Range(min, max int32) int32 {
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
func (f *Faker) Int64Range(min, max int64) int64 {
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
func (f *Faker) Float32Range(min, max float32) float32 {
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
func (f *Faker) Float64Range(min, max float64) float64 {
	if min >= max {
		return min
	}

	if f.safe {
		f.mutex.Lock()
		defer f.mutex.Unlock()
	}

	return f.src.Float64()*(max-min) + min
}

// Numerify will replace # with random numerical values (0-9 digits)
func (f *Faker) Numerify(str string) string {
	return f.replaceWithNumbers(str)
}

// ShuffleInts will randomize a slice of ints
func (f *Faker) ShuffleInts(a []int) {
	for i := range a {
		j := f.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}

// Intn returns a random number between [0, n)
func (f *Faker) Intn(n int) int {
	if f.safe {
		f.mutex.Lock()
		defer f.mutex.Unlock()
	}
	return f.src.Intn(n)
}

// Float32 returns a random positive float number [0.000000000001, math.MaxFloat32)
func (f *Faker) Float32() float32 {
	return f.Float32Range(math.SmallestNonzeroFloat32, math.MaxFloat32)
}

// Float64 returns a random positive float number [0.000000000001, math.MaxFloat64)
func (f *Faker) Float64() float64 {
	return f.Float64Range(math.SmallestNonzeroFloat64, math.MaxFloat64)
}
