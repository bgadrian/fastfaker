package faker

import "strconv"

// Faker::Beer.blg #=> "18.5°Blg"

// BeerName will return a random beer name
func (f *Faker) BeerName() string {
	return f.getRandValue([]string{"beer", "name"})
}

// BeerStyle will return a random beer style
func (f *Faker) BeerStyle() string {
	return f.getRandValue([]string{"beer", "style"})
}

// BeerHop will return a random beer hop
func (f *Faker) BeerHop() string {
	return f.getRandValue([]string{"beer", "hop"})
}

// BeerYeast will return a random beer yeast
func (f *Faker) BeerYeast() string {
	return f.getRandValue([]string{"beer", "yeast"})
}

// BeerMalt will return a random beer malt
func (f *Faker) BeerMalt() string {
	return f.getRandValue([]string{"beer", "malt"})
}

// BeerIbu will return a random beer ibu value between 10 and 100
func (f *Faker) BeerIbu() string {
	return strconv.Itoa(f.Number(10, 100)) + " IBU"
}

// BeerAlcohol will return a random beer alcohol level between 2.0 and 10.0
func (f *Faker) BeerAlcohol() string {
	return strconv.FormatFloat(f.Float64Range(2.0, 10.0), 'f', 1, 64) + "%"
}

// BeerBlg will return a random beer blg between 5.0 and 20.0
func (f *Faker) BeerBlg() string {
	return strconv.FormatFloat(f.Float64Range(5.0, 20.0), 'f', 1, 64) + "°Blg"
}
