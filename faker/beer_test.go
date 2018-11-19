package faker

import (
	"fmt"
	"testing"
)

func ExampleFaker_BeerName() {
	Global.Seed(11)
	fmt.Println(Global.BeerName())
	// Output: Duvel
}

func BenchmarkBeerName(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.BeerName()
	}
}

func ExampleFaker_BeerStyle() {
	Global.Seed(11)
	fmt.Println(Global.BeerStyle())
	// Output: European Amber Lager
}

func BenchmarkBeerStyle(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.BeerStyle()
	}
}

func ExampleFaker_BeerHop() {
	Global.Seed(11)
	fmt.Println(Global.BeerHop())
	// Output: Glacier
}

func BenchmarkBeerHop(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.BeerHop()
	}
}

func ExampleFaker_BeerYeast() {
	Global.Seed(11)
	fmt.Println(Global.BeerYeast())
	// Output: 1388 - Belgian Strong Ale
}

func BenchmarkBeerYeast(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.BeerYeast()
	}
}

func ExampleFaker_BeerMalt() {
	Global.Seed(11)
	fmt.Println(Global.BeerMalt())
	// Output: Munich
}

func BenchmarkBeerMalt(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.BeerMalt()
	}
}

func ExampleFaker_BeerIbu() {
	Global.Seed(11)
	fmt.Println(Global.BeerIbu())
	// Output: 29 IBU
}

func BenchmarkBeerIbu(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.BeerIbu()
	}
}

func ExampleFaker_BeerAlcohol() {
	Global.Seed(11)
	fmt.Println(Global.BeerAlcohol())
	// Output: 2.7%
}

func BenchmarkBeerAlcohol(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.BeerAlcohol()
	}
}

func ExampleFaker_BeerBlg() {
	Global.Seed(11)
	fmt.Println(Global.BeerBlg())
	// Output: 6.4°Blg
}

func BenchmarkBeerBlg(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.BeerBlg()
	}
}
