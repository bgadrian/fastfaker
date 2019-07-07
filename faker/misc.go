package faker

import (
	"github.com/bgadrian/fastfaker/data"
)

const hashtag = '#'
const questionmark = '?'

// Get Random Value
func (f *Faker) getRandValue(dataVal []string) string {
	val, err := data.GetRandValue(f, dataVal)
	if err != nil {
		return ""
	}
	return val
}

// Get Random Integer Value
func (f *Faker) getRandIntValue(dataVal []string) int {
	val, err := data.GetRandIntValue(f, dataVal)
	if err != nil {
		return 0
	}
	return val
}

// Replace # with numbers
func (f *Faker) replaceWithNumbers(str string) string {
	if str == "" {
		return str
	}
	bytestr := []byte(str)
	for i := 0; i < len(bytestr); i++ {
		if bytestr[i] == hashtag {
			bytestr[i] = f.randDigit()
		}
	}

	return string(bytestr)
}

// Replace ? with ASCII lowercase letters
func (f *Faker) replaceWithLetters(str string) string {
	if str == "" {
		return str
	}
	bytestr := []byte(str)
	for i := 0; i < len(bytestr); i++ {
		if bytestr[i] == questionmark {
			bytestr[i] = byte(f.randLetter())
		}
	}

	return string(bytestr)
}

// Generate random lowercase ASCII letter
func (f *Faker) randLetter() rune {
	return rune(byte(f.Intn(26)) + 'a')
}

// Generate random ASCII digit
var numerifyDictonary = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

// randDigit will generate a random ASCII digit, returning its byte/rune value
func (f *Faker) randDigit() byte {
	return numerifyDictonary[f.Intn(10)]
}

// Numerify will replace # with random numerical values (0-9 digits)
func (f *Faker) Numerify(str string) string {
	return f.replaceWithNumbers(str)
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
