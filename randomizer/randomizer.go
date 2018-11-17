/*
Package randomizer contains pseudo-random data generators, used by the Faker main package.
It is concurrent safe, has 100% test coverage, benchmarks and examples.
Can be used as a rand.* replacement.

RandWrapper encapsulates a rand.Rand but it is more user friendly and has more functions.
It has 2 modes:
* fast - for single-thread/goroutine performance requirements
* safe - slower but safe for concurrent use (like rand.* it uses a mutex)

For more info and examples see https://github.com/bgadrian/fastfaker
*/
package randomizer

// Randomizer represents a base generator for the faker
type Randomizer interface {
	Seed(seed int64)
	Number(min int, max int) int
	Uint8() uint8
	Uint16() uint16
	Uint32() uint32
	Uint64() uint64
	Int8() int8
	Int16() int16
	Int32() int32
	Int64() int64
	Int64Positive() int64
	Int32Range(min, max int32) int32
	Int64Range(min, max int64) int64
	Float32Range(min, max float32) float32
	Float64Range(min, max float64) float64
	ShuffleInts(a []int)
	Intn(n int) int
	Float32() float32
	Float64() float64
	Read(p []byte) (n int, err error)
	Float64Unary() float64
}
