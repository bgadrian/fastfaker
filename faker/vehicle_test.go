package faker

import (
	"fmt"
	"testing"
)

func ExampleFaker_CarModel() {
	Global.Seed(11)
	fmt.Println(Global.CarModel())
	// Output: Aveo
}

func ExampleFaker_CarMaker() {
	Global.Seed(11)
	fmt.Println(Global.CarMaker())
	// Output: Nissan
}

func ExampleFaker_TransmissionGearType() {
	Global.Seed(11)
	fmt.Println(Global.TransmissionGearType())
	// Output: Manual
}

func ExampleFaker_FuelType() {
	Global.Seed(11)
	fmt.Println(Global.FuelType())
	// Output: CNG
}

func ExampleFaker_VehicleType() {
	Global.Seed(11)
	fmt.Println(Global.VehicleType())
	// Output: Passenger car mini
}

func ExampleFaker_Vehicle() {
	Global.Seed(11)
	vehicle := Global.Vehicle()
	fmt.Println(vehicle.Brand)
	fmt.Println(vehicle.Fuel)
	fmt.Println(vehicle.Model)
	fmt.Println(vehicle.TransmissionGear)
	fmt.Println(vehicle.VehicleType)
	fmt.Println(vehicle.Year)

	// Output: Fiat
	// Gasoline
	// Freestyle Fwd
	// Automatic
	// Passenger car mini
	// 1943
}

func BenchmarkFaker_CarMaker(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.CarMaker()
	}
}

func BenchmarkFaker_CarModel(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.CarModel()
	}
}

func BenchmarkFaker_Vehicle(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.Vehicle()
	}
}

func BenchmarkFaker_VehicleType(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.VehicleType()
	}
}
