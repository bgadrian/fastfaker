package faker

import (
	"fmt"
	"testing"

	"github.com/engoengine/math"
)

func ExampleFaker_Number() {
	Global.Seed(11)
	fmt.Println(Global.Number(50, 23456))
	// Output: 14866
}

func BenchmarkNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.Number(10, 999999)
	}
}

func ExampleFaker_Uint8() {
	Global.Seed(11)
	fmt.Println(Global.Uint8())
	// Output: 152
}

func BenchmarkUint8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.Uint8()
	}
}

func ExampleFaker_Uint16() {
	Global.Seed(11)
	fmt.Println(Global.Uint16())
	// Output: 34968
}

func BenchmarkUint16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.Uint16()
	}
}

func ExampleFaker_Uint32() {
	Global.Seed(11)
	fmt.Println(Global.Uint32())
	// Output: 1075055705
}

func BenchmarkUint32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.Uint32()
	}
}

func ExampleFaker_Uint64() {
	Global.Seed(11)
	fmt.Println(Global.Uint64())
	// Output: 843730692693298265
}

func BenchmarkUint64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.Uint64()
	}
}

func ExampleFaker_Int8() {
	Global.Seed(11)
	fmt.Println(Global.Int8())
	// Output: 24
}

func BenchmarkInt8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.Int8()
	}
}

func ExampleFaker_Int16() {
	Global.Seed(11)
	fmt.Println(Global.Int16())
	// Output: 2200
}

func BenchmarkInt16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.Int16()
	}
}

func ExampleFaker_Int32() {
	Global.Seed(11)
	fmt.Println(Global.Int32())
	// Output: -1072427943
}

func BenchmarkInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.Int32()
	}
}

func ExampleFaker_Int64() {
	Global.Seed(11)
	fmt.Println(Global.Int64())
	// Output: -8379641344161477543
}

func BenchmarkInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.Int64()
	}
}

func ExampleFaker_Float32() {
	Global.Seed(11)
	fmt.Println(Global.Float32())
	// Output: 3.1128167e+37
}

func BenchmarkFloat32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.Float32()
	}
}

func ExampleFaker_Float32Range() {
	Global.Seed(11)
	fmt.Println(Global.Float32Range(0, 9999999))
	// Output: 914774.6
}

func BenchmarkFloat32Range(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.Float32Range(0, 9999999)
	}
}

func ExampleFaker_Float64() {
	Global.Seed(11)
	fmt.Println(Global.Float64())
	// Output: 1.644484108270445e+307
}

func BenchmarkFloat64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.Float64()
	}
}

func ExampleFaker_Float64Range() {
	Global.Seed(11)
	fmt.Println(Global.Float64Range(0, 9999999))
	// Output: 914774.5585333086
}

func BenchmarkFloat64Range(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.Float64Range(0, 9999999)
	}
}

func ExampleFaker_ShuffleInts() {
	Global.Seed(11)

	ints := []int{52, 854, 941, 74125, 8413, 777, 89416, 841657}
	Global.ShuffleInts(ints)
	fmt.Println(ints)
	// Output: [74125 777 941 89416 8413 854 52 841657]
}

func BenchmarkShuffleInts(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.ShuffleInts([]int{52, 854, 941, 74125, 8413, 777, 89416, 841657})
	}
}

func TestFaker_Intn(t *testing.T) {
	Global.Seed(42)
	for i := 0; i < 5000; i++ {
		v := Global.Intn(101)
		if v >= 0 && v < 101 {
			continue
		}
		t.Errorf("out of range number [0,%d): %d",
			101, v)
	}
}

func TestFaker_Int32Range(t *testing.T) {
	Global.Seed(42)
	min := int32(-100)
	max := int32(25)
	for i := 0; i < 5000; i++ {
		v := Global.Int32Range(min, max)
		if v >= min && v < max {
			continue
		}
		t.Errorf("out of range number [%d,%d): %d",
			min, max, v)
	}

	v := Global.Int32Range(43, 31)
	if v != 43 {
		t.Errorf("result is not min when bigger than max, exp %d got %d",
			min, v)
	}
}

func TestFaker_Int64Range(t *testing.T) {
	Global.Seed(42)
	//we choose a range outside of int32
	//and small enough to check the out of bounds
	//for a series of 5000
	min := int64(math.MinInt64 + 100)
	max := int64(min + 150)
	for i := 0; i < 5000; i++ {
		v := Global.Int64Range(min, max)
		if v >= min && v < max {
			continue
		}
		t.Errorf("out of range number [%d,%d): %d",
			min, max, v)
	}

	v := Global.Int64Range(43, 31)
	if v != 43 {
		t.Errorf("result is not min when bigger than max, exp %d got %d",
			min, v)
	}
}

func TestFaker_Float32Range(t *testing.T) {
	Global.Seed(42)
	//we choose a range outside of float32
	//and small enough to check the out of bounds
	//for a series of 5000
	min := float32(math.SmallestNonzeroFloat32 + float32(23.5))
	max := float32(min + float32(12.5))
	for i := 0; i < 5000; i++ {
		v := Global.Float32Range(min, max)
		if v >= min && v < max {
			continue
		}
		t.Errorf("out of range number [%f,%f): %f",
			min, max, v)
	}
}

func TestFaker_Float64Range(t *testing.T) {
	Global.Seed(42)
	//we choose a range outside of float32
	//and small enough to check the out of bounds
	//for a series of 5000
	min := float64(math.SmallestNonzeroFloat64 + 100.2)
	max := float64(min + 15.3)
	for i := 0; i < 5000; i++ {
		v := Global.Float64Range(min, max)
		if v >= min && v < max {
			continue
		}
		t.Errorf("out of range number [%f,%f): %f",
			min, max, v)
	}
}

func TestFaker_RandIntRange(t *testing.T) {
	if Global.Number(5, 5) != 5 {
		t.Error("You should have gotten 5 back")
	}
}

func TestRandFloat32RangeSame(t *testing.T) {
	if Global.Float32Range(5.0, 5.0) != 5.0 {
		t.Error("You should have gotten 5.0 back")
	}
}

func TestRandFloat64RangeSame(t *testing.T) {
	if Global.Float64Range(5.0, 5.0) != 5.0 {
		t.Error("You should have gotten 5.0 back")
	}
}
