package fastfaker

import (
	"github.com/bgadrian/fastfaker/data"
)

const hashtag = '#'
const questionmark = '?'

// Check if in lib
func dataCheck(dataVal []string) bool {
	var checkOk bool

	if len(dataVal) == 2 {
		_, checkOk = data.Data[dataVal[0]]
		if checkOk {
			_, checkOk = data.Data[dataVal[0]][dataVal[1]]
		}
	}

	return checkOk
}

// Check if in lib
func intDataCheck(dataVal []string) bool {
	if len(dataVal) != 2 {
		return false
	}

	_, checkOk := data.IntData[dataVal[0]]
	if checkOk {
		_, checkOk = data.IntData[dataVal[0]][dataVal[1]]
	}

	return checkOk
}

// Get Random Value
func (f *Faker) getRandValue(dataVal []string) string {
	if !dataCheck(dataVal) {
		return ""
	}
	return data.Data[dataVal[0]][dataVal[1]][f.Intn(len(data.Data[dataVal[0]][dataVal[1]]))]
}

// Get Random Integer Value
func (f *Faker) getRandIntValue(dataVal []string) int {
	if !intDataCheck(dataVal) {
		return 0
	}
	return data.IntData[dataVal[0]][dataVal[1]][f.Intn(len(data.IntData[dataVal[0]][dataVal[1]]))]
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

// Generate random integer between min and max
func (f *Faker) randIntRange(min, max int) int {
	if min == max {
		return min
	}
	return f.Intn((max+1)-min) + min
}

// Categories will return a map string array of available data categories and sub categories
func Categories() map[string][]string {
	types := make(map[string][]string)
	for category, subCategoriesMap := range data.Data {
		subCategories := make([]string, 0)
		for subType := range subCategoriesMap {
			subCategories = append(subCategories, subType)
		}
		types[category] = subCategories
	}
	return types
}
