package faker

import (
	"fmt"
	"testing"
)

func ExampleFaker_Color() {
	Global.Seed(11)
	fmt.Println(Global.Color())
	// Output: MediumOrchid
}

func BenchmarkColor(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.Color()
	}
}

func ExampleFaker_SafeColor() {
	Global.Seed(11)
	fmt.Println(Global.SafeColor())
	// Output: black
}

func BenchmarkSafeColor(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.SafeColor()
	}
}

func ExampleFaker_HexColor() {
	Global.Seed(42)
	fmt.Println(Global.HexColor())
	// Output: #1b4ef1
}

func BenchmarkHexColor(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.HexColor()
	}
}

func ExampleFaker_RGBColor() {
	Global.Seed(11)
	fmt.Println(Global.RGBColor())
	// Output: [152 23 53]
}

func BenchmarkRGBColor(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.RGBColor()
	}
}
