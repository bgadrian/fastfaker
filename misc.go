package fastfaker

import (
	"github.com/bgadrian/fastfaker/data"
)

// Get Random Value
func (f *Faker) getRandValue(dataVal []string) string {
	//TODO treat the notfound case
	val, _ := data.GetRandValue(f, dataVal)
	return val
}

// Get Random Integer Value
func (f *Faker) getRandIntValue(dataVal []string) int {
	//TODO treat the notfound case
	val, _ := data.GetRandIntValue(f, dataVal)
	return val
}

func clampInt(val, min, max int) int {
	if val < min {
		return min
	}
	if val > max {
		return max
	}
	return val
}
