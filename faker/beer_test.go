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
	for i := 0; i < b.N; i++ {
		Global.BeerName()
	}
}

func ExampleFaker_BeerStyle() {
	Global.Seed(11)
	fmt.Println(Global.BeerStyle())
	// Output: European Amber Lager
}

func BenchmarkBeerStyle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.BeerStyle()
	}
}

func ExampleFaker_BeerHop() {
	Global.Seed(11)
	fmt.Println(Global.BeerHop())
	// Output: Glacier
}

func BenchmarkBeerHop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.BeerHop()
	}
}

func ExampleFaker_BeerYeast() {
	Global.Seed(11)
	fmt.Println(Global.BeerYeast())
	// Output: 1388 - Belgian Strong Ale
}

func BenchmarkBeerYeast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.BeerYeast()
	}
}

func ExampleFaker_BeerMalt() {
	Global.Seed(11)
	fmt.Println(Global.BeerMalt())
	// Output: Munich
}

func BenchmarkBeerMalt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.BeerMalt()
	}
}

func ExampleFaker_BeerIbu() {
	Global.Seed(11)
	fmt.Println(Global.BeerIbu())
	// Output: 29 IBU
}

func BenchmarkBeerIbu(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.BeerIbu()
	}
}

func ExampleFaker_BeerAlcohol() {
	Global.Seed(11)
	fmt.Println(Global.BeerAlcohol())
	// Output: 2.7%
}

func BenchmarkBeerAlcohol(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.BeerAlcohol()
	}
}

func ExampleFaker_BeerBlg() {
	Global.Seed(11)
	fmt.Println(Global.BeerBlg())
	// Output: 6.4°Blg
}

func BenchmarkBeerBlg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.BeerBlg()
	}
}
