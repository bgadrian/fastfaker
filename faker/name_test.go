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
	for i := 0; i < b.N; i++ {
		Global.Name()
	}
}

func ExampleFaker_FirstName() {
	Global.Seed(11)
	fmt.Println(Global.FirstName())
	// Output: Markus
}

func BenchmarkFirstName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.FirstName()
	}
}

func ExampleFaker_LastName() {
	Global.Seed(11)
	fmt.Println(Global.LastName())
	// Output: Daniel
}

func BenchmarkLastName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.LastName()
	}
}

func ExampleFaker_NamePrefix() {
	Global.Seed(11)
	fmt.Println(Global.NamePrefix())
	// Output: Mr.
}

func BenchmarkNamePrefix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.NamePrefix()
	}
}

func ExampleFaker_NameSuffix() {
	Global.Seed(11)
	fmt.Println(Global.NameSuffix())
	// Output: Jr.
}

func BenchmarkNameSuffix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.NameSuffix()
	}
}
