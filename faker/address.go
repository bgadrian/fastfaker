package faker

import (
	"errors"
	"strings"
)

// ErrRangeInvalid signals a bad input
var ErrRangeInvalid = errors.New("input range is invalid")

// AddressInfo is a struct full of address information
type AddressInfo struct {
	Address   string
	Street    string
	City      string
	State     string
	Zip       string
	Country   string
	Latitude  float64
	Longitude float64
}

// Address will generate a struct of address information
func (f *Faker) Address() *AddressInfo {
	street := f.Street()
	city := f.City()
	state := f.State()
	zip := f.Zip()

	return &AddressInfo{
		Address:   street + ", " + city + ", " + state + " " + zip,
		Street:    street,
		City:      city,
		State:     state,
		Zip:       zip,
		Country:   f.Country(),
		Latitude:  f.Latitude(),
		Longitude: f.Longitude(),
	}
}

// Street will generate a random address street string
func (f *Faker) Street() (street string) {
	switch randInt := f.Number(1, 2); randInt {
	case 1:
		street = f.StreetNumber() + " " + f.StreetPrefix() + " " + f.StreetName() + f.StreetSuffix()
	case 2:
		street = f.StreetNumber() + " " + f.StreetName() + f.StreetSuffix()
	}

	return
}

// StreetNumber will generate a random address street number string
func (f *Faker) StreetNumber() string {
	return strings.TrimLeft(f.replaceWithNumbers(f.getRandValue([]string{"address", "number"})), "0")
}

// StreetPrefix will generate a random address street prefix string
func (f *Faker) StreetPrefix() string {
	return f.getRandValue([]string{"address", "street_prefix"})
}

// StreetName will generate a random address street name string
func (f *Faker) StreetName() string {
	return f.getRandValue([]string{"address", "street_name"})
}

// StreetSuffix will generate a random address street suffix string
func (f *Faker) StreetSuffix() string {
	return f.getRandValue([]string{"address", "street_suffix"})
}

// City will generate a random city string
func (f *Faker) City() (city string) {
	switch randInt := f.Number(1, 3); randInt {
	case 1:
		city = f.FirstName() + f.StreetSuffix()
	case 2:
		city = f.LastName() + f.StreetSuffix()
	case 3:
		city = f.StreetPrefix() + " " + f.LastName()
	}

	return
}

// State will generate a random state string
func (f *Faker) State() string {
	return f.getRandValue([]string{"address", "state"})
}

// StateAbr will generate a random abbreviated state string
func (f *Faker) StateAbr() string {
	return f.getRandValue([]string{"address", "state_abr"})
}

// Zip will generate a random Zip code string
func (f *Faker) Zip() string {
	return f.replaceWithNumbers(f.getRandValue([]string{"address", "zip"}))
}

// Country will generate a random country string
func (f *Faker) Country() string {
	return f.getRandValue([]string{"address", "country"})
}

// CountryAbr will generate a random abbreviated country string
func (f *Faker) CountryAbr() string {
	return f.getRandValue([]string{"address", "country_abr"})
}

// Latitude will generate a random latitude float64
func (f *Faker) Latitude() float64 {
	return (f.Float64Unary() * 180) - 90
}

// LatitudeInRange will generate a random latitude within the input range
func (f *Faker) LatitudeInRange(min, max float64) (float64, error) {
	if min > max || min < -90 || min > 90 || max < -90 || max > 90 {
		return 0, ErrRangeInvalid
	}
	return f.Float64Range(min, max), nil
}

// Longitude will generate a random longitude float64
func (f *Faker) Longitude() float64 {
	return (f.Float64Unary() * 360) - 180
}

// LongitudeInRange will generate a random longitude within the input range
func (f *Faker) LongitudeInRange(min, max float64) (float64, error) {
	if min > max || min < -180 || min > 180 || max < -180 || max > 180 {
		return 0, ErrRangeInvalid
	}
	return f.Float64Range(min, max), nil
}
