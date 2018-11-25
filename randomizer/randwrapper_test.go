package randomizer

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"testing"
)

func Example() {
	gen := NewRandWrapper(false, nil)
	gen.Seed(42)

	fmt.Printf("Int8: %d\n", gen.Int8())
	fmt.Printf("Int16: %d\n", gen.Int16())
	fmt.Printf("Int32: %d\n", gen.Int32())
	fmt.Printf("Int32: %d\n", gen.Int64())
	fmt.Printf("Int64Positive: %d\n", gen.Int64Positive())

	fmt.Printf("Uint8: %d\n", gen.Uint8())
	fmt.Printf("Uint16: %d\n", gen.Uint16())
	fmt.Printf("Uint32: %d\n", gen.Uint32())

	fmt.Printf("Float32: %f\n", gen.Float32())
	fmt.Printf("Float64: %f\n", gen.Float64())
	// Output: Int8: 49
	//Int16: 13387
	//Int32: -1805934616
	//Int32: -7297359450328151799
	//Int64Positive: 404153945743547657
	//Uint8: 97
	//Uint16: 23461
	//Uint32: 1276434112
	//Float32: 130343333431123176975039154477607157760.000000
	//Float64: 116198632802140183205174997036310832568880167563288775268797842968811055189411147761992698981182854609204521400932481006915558598725364731145245182600684556023951651432775097450784441328418806702862825300469152646818691585777068968631119723028318015329488476327072796506412228424374316456135062145352853356544.000000
}

func TestRandWrapperImplementsRandomizer(t *testing.T) {
	if Randomizer(&RandWrapper{}) == nil {
		t.Error("cannot cast to interface")
	}
}

func TestRandWrapper_Intn(t *testing.T) {
	randomizer := NewRandWrapper(false, rand.NewSource(42))
	for i := 0; i < 5000; i++ {
		v := randomizer.Intn(101)
		if v >= 0 && v < 101 {
			continue
		}
		t.Errorf("out of range number [0,%d): %d",
			101, v)
	}
}

func TestFaker_Int32Range(t *testing.T) {
	randomizer := NewRandWrapper(false, nil)
	randomizer.Seed(42)

	min := int32(-100)
	max := int32(25)
	for i := 0; i < 5000; i++ {
		v := randomizer.Int32Range(min, max)
		if v >= min && v < max {
			continue
		}
		t.Errorf("out of range number [%d,%d): %d",
			min, max, v)
	}

	v := randomizer.Int32Range(43, 31)
	if v != 43 {
		t.Errorf("result is not min when bigger than max, exp %d got %d",
			min, v)
	}
}

func TestFaker_Int64Range(t *testing.T) {
	randomizer := NewRandWrapper(false, nil)
	randomizer.Seed(42)

	//we choose a range outside of int32
	//and small enough to check the out of bounds
	//for a series of 5000
	min := int64(math.MinInt64 + 100)
	max := int64(min + 150)
	for i := 0; i < 5000; i++ {
		v := randomizer.Int64Range(min, max)
		if v >= min && v < max {
			continue
		}
		t.Errorf("out of range number [%d,%d): %d",
			min, max, v)
	}

	v := randomizer.Int64Range(43, 31)
	if v != 43 {
		t.Errorf("result is not min when bigger than max, exp %d got %d",
			min, v)
	}
}

func TestFaker_Float32Range(t *testing.T) {
	randomizer := NewRandWrapper(false, nil)
	randomizer.Seed(42)

	//we choose a range outside of float32
	//and small enough to check the out of bounds
	//for a series of 5000
	min := float32(math.SmallestNonzeroFloat32 + float32(23.5))
	max := float32(min + float32(12.5))
	for i := 0; i < 5000; i++ {
		v := randomizer.Float32Range(min, max)
		if v >= min && v < max {
			continue
		}
		t.Errorf("out of range number [%f,%f): %f",
			min, max, v)
	}

	if randomizer.Float32Range(42, 1) != 42 {
		t.Errorf("must return min when input is wrong")
	}
}

func TestFaker_Float64Range(t *testing.T) {
	randomizer := NewRandWrapper(false, nil)
	randomizer.Seed(42)

	//we choose a range outside of float32
	//and small enough to check the out of bounds
	//for a series of 5000
	min := float64(math.SmallestNonzeroFloat64 + 100.2)
	max := float64(min + 15.3)
	for i := 0; i < 5000; i++ {
		v := randomizer.Float64Range(min, max)
		if v >= min && v < max {
			continue
		}
		t.Errorf("out of range number [%f,%f): %f",
			min, max, v)
	}

	if randomizer.Float64Range(42, 1) != 42 {
		t.Errorf("must return min when input is wrong")
	}
}

func TestNewSafeFaker_Deadlocks(t *testing.T) {
	wg := sync.WaitGroup{}
	steps := 8000
	threads := 4
	safe := NewRandWrapper(true, nil)
	wg.Add(threads)

	for test := 0; test < threads; test++ {
		go func() {
			for i := 0; i < steps; i++ {
				as64 := int64(i + 10000)
				safe.Seed(as64)
				safe.Number(-i, i)
				safe.Int64()
				safe.Int64Positive()
				safe.Uint64()
				safe.Int32Range(42, 100)
				safe.Int64Range(int64(-i-2000), as64)
				safe.Float32Range(32.3, 45.6)
				safe.Float64Range(-23.2, 452.2)
				safe.ShuffleInts([]int{3, 2, 1})
				safe.Intn(i + 102)
				_, err := safe.Read([]byte{0, 0, 0})
				if err != nil {
					t.Error(err)
				}
				safe.Float64Unary()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
