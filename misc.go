package fastfaker

import (
	"github.com/bgadrian/fastfaker/data"
)

const hashtag = '#'
const questionmark = '?'

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

// Replace # with numbers
func (f *Faker) replaceWithNumbers(str string) string {
	if str == "" {
		return str
	}
	bytestr := []byte(str)
	for i := 0; i < len(bytestr); i++ {
		if bytestr[i] == hashtag {
			bytestr[i] = byte(f.randDigit())
		}
	}
	if bytestr[0] == '0' {
		bytestr[0] = byte(f.Intn(8)+1) + '0'
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
func (f *Faker) randDigit() rune {
	return rune(byte(f.Intn(9)) + '0')
}

// Numerify will replace # with random numerical values (0-9 digits)
func (f *Faker) Numerify(str string) string {
	return f.replaceWithNumbers(str)
}
