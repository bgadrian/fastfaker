package faker

import (
	"math"

	"github.com/bgadrian/fastfaker/data"
)

// CurrencyInfo is a struct of currency information
type CurrencyInfo struct {
	Short string
	Long  string
}

// Currency will generate a struct with random currency information
func (f *Faker) Currency() *CurrencyInfo {
	index := f.Intn(len(data.Data["currency"]["short"]))
	return &CurrencyInfo{
		Short: data.Data["currency"]["short"][index],
		Long:  data.Data["currency"]["long"][index],
	}
}

// CurrencyShort will generate a random short currency value
func (f *Faker) CurrencyShort() string {
	return f.getRandValue([]string{"currency", "short"})
}

// CurrencyLong will generate a random long currency name
func (f *Faker) CurrencyLong() string {
	return f.getRandValue([]string{"currency", "long"})
}

// Price will take in a min and max value and return a formatted price
func (f *Faker) Price(min, max float64) float64 {
	return math.Floor(f.Float64Range(min, max)*100) / 100
}
