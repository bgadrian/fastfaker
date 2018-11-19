package faker

import (
	"fmt"
	"testing"
)

func ExampleFaker_Name() {
	Global.Seed(11)
	fmt.Println(Global.Name())
	// Output: Markus Moen
}

func BenchmarkName(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.Name()
	}
}

func ExampleFaker_FirstName() {
	Global.Seed(11)
	fmt.Println(Global.FirstName())
	// Output: Markus
}

func BenchmarkFirstName(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.FirstName()
	}
}

func ExampleFaker_LastName() {
	Global.Seed(11)
	fmt.Println(Global.LastName())
	// Output: Daniel
}

func BenchmarkLastName(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.LastName()
	}
}

func ExampleFaker_NamePrefix() {
	Global.Seed(11)
	fmt.Println(Global.NamePrefix())
	// Output: Mr.
}

func BenchmarkNamePrefix(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.NamePrefix()
	}
}

func ExampleFaker_NameSuffix() {
	Global.Seed(11)
	fmt.Println(Global.NameSuffix())
	// Output: Jr.
}

func BenchmarkNameSuffix(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.NameSuffix()
	}
}
