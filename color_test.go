package gofakeit

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
	for i := 0; i < b.N; i++ {
		Global.Color()
	}
}

func ExampleFaker_SafeColor() {
	Global.Seed(11)
	fmt.Println(Global.SafeColor())
	// Output: black
}

func BenchmarkSafeColor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.SafeColor()
	}
}

func ExampleFaker_HexColor() {
	Global.Seed(11)
	fmt.Println(Global.HexColor())
	// Output: #i15jb7
}

func BenchmarkHexColor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.HexColor()
	}
}

func ExampleFaker_RGBColor() {
	Global.Seed(11)
	fmt.Println(Global.RGBColor())
	// Output: [152 23 53]
}

func BenchmarkRGBColor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.RGBColor()
	}
}
