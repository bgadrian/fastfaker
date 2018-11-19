package faker

import (
	"fmt"
	"testing"
)

func ExampleFaker_Address() {
	Global.Seed(11)
	address := Global.Address()
	fmt.Println(address.Address)
	fmt.Println(address.Street)
	fmt.Println(address.City)
	fmt.Println(address.State)
	fmt.Println(address.Zip)
	fmt.Println(address.Country)
	fmt.Println(address.Latitude)
	fmt.Println(address.Longitude)
	// 872 East Rapidsborough, Rutherfordstad, New Jersey 74853
	// 872 East Rapidsborough
	// Rutherfordstad
	// New Jersey
	// 74853
	// South Africa
	// 23.05875828427908
	// 89.02259415693374
}

func BenchmarkAddress(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.Address()
	}
}

func ExampleFaker_Street() {
	Global.Seed(11)
	fmt.Println(Global.Street())
	// Output: 364 East Rapidsborough
}

func BenchmarkStreet(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.Street()
	}
}

func ExampleFaker_StreetNumber() {
	Global.Seed(11)
	fmt.Println(Global.StreetNumber())
	// Output: 13645
}

func BenchmarkStreetNumber(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.StreetNumber()
	}
}

func ExampleFaker_StreetPrefix() {
	Global.Seed(11)
	fmt.Println(Global.StreetPrefix())
	// Output: Lake
}

func BenchmarkStreetPrefix(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.StreetPrefix()
	}
}

func ExampleFaker_StreetName() {
	Global.Seed(11)
	fmt.Println(Global.StreetName())
	// Output: View
}

func BenchmarkStreetName(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.StreetName()
	}
}

func ExampleFaker_StreetSuffix() {
	Global.Seed(11)
	fmt.Println(Global.StreetSuffix())
	// Output: land
}

func BenchmarkStreetSuffix(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.StreetSuffix()
	}
}

func ExampleFaker_City() {
	Global.Seed(11)
	fmt.Println(Global.City())
	// Output: Marcelside
}

func TestCity(t *testing.T) {
	for i := 0; i < 100; i++ {
		Global.City()
	}
}

func BenchmarkCity(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.City()
	}
}

func ExampleFaker_State() {
	Global.Seed(11)
	fmt.Println(Global.State())
	// Output: Hawaii
}

func BenchmarkState(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.State()
	}
}

func ExampleFaker_StateAbr() {
	Global.Seed(11)
	fmt.Println(Global.StateAbr())
	// Output: OR
}

func BenchmarkStateAbr(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.StateAbr()
	}
}

func ExampleFaker_Zip() {
	Global.Seed(11)
	fmt.Println(Global.Zip())
	// Output: 13645
}

func BenchmarkZip(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.Zip()
	}
}

func ExampleFaker_Country() {
	Global.Seed(11)
	fmt.Println(Global.Country())
	// Output: Tajikistan
}

func BenchmarkCountry(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.Country()
	}
}

func ExampleFaker_CountryAbr() {
	Global.Seed(11)
	fmt.Println(Global.CountryAbr())
	// Output: FI
}

func BenchmarkCountryAbr(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.CountryAbr()
	}
}

func ExampleFaker_Latitude() {
	Global.Seed(11)
	fmt.Println(Global.Latitude())
	// Output: -73.53405629980608
}

func BenchmarkLatitude(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.Latitude()
	}
}

func ExampleFaker_Longitude() {
	Global.Seed(11)
	fmt.Println(Global.Longitude())
	// Output: -147.06811259961216
}

func BenchmarkLongitude(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.Longitude()
	}
}

func TestLatitudeInRange(t *testing.T) {
	Global.Seed(11)
	lat, err := Global.LatitudeInRange(21, 42)
	if err != nil {
		t.Error("error should be nil")
	}

	if lat == 0 {
		t.Error("lat should be not be zero")
	}

	_, err = Global.LatitudeInRange(50, 42)
	if err == nil {
		t.Error("error should be not be nil")
	}

	_, err = Global.LatitudeInRange(-100, 42)
	if err == nil {
		t.Error("error should be not be nil")
	}
}

func ExampleFaker_LatitudeInRange() {
	Global.Seed(11)
	lat, _ := Global.LatitudeInRange(21, 42)
	fmt.Println(lat)
	// Output: 22.921026765022624
}

func BenchmarkLatitudeInRange(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.LatitudeInRange(-90, 90)
	}
}

func TestLongitudeInRange(t *testing.T) {
	Global.Seed(11)
	long, err := Global.LongitudeInRange(21, 42)
	if err != nil {
		t.Error("error should be nil")
	}

	if long == 0 {
		t.Error("long should be not be zero")
	}

	_, err = Global.LongitudeInRange(-32, -42)
	if err == nil {
		t.Error("error should be not be nil")
	}

	_, err = Global.LongitudeInRange(190, 192)
	if err == nil {
		t.Error("error should be not be nil")
	}
}

func ExampleFaker_LongitudeInRange() {
	Global.Seed(11)
	long, _ := Global.LongitudeInRange(-10, 10)
	fmt.Println(long)
	// Output: -8.170450699978453
}

func BenchmarkLongitudeInRange(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.LongitudeInRange(-180, 180)
	}
}
