package faker

import (
	"strings"

	"github.com/bgadrian/fastfaker/data"
)

// Generate fake information from given string. String should contain {category.subcategory}
//
// Ex: {person.first} - random firstname
//
// Ex: {person.first}###{person.last}@{person.last}.{internet.domain_suffix} - billy834smith@smith.com
//
// Ex: ### - 481 - random numbers
//
// Ex: ??? - fda - random letters
//
// For a complete list possible categories use the Categories() function.
func (f *Faker) Generate(dataVal string) string {
	// Identify items between brackets: {person.first}
	for strings.Count(dataVal, "{") > 0 && strings.Count(dataVal, "}") > 0 {
		catValue := ""
		startIndex := strings.Index(dataVal, "{")
		endIndex := strings.Index(dataVal, "}")
		replace := dataVal[(startIndex + 1):endIndex]
		categories := strings.Split(replace, ".")

		if len(categories) >= 2 && data.Check([]string{categories[0], categories[1]}) {
			catValue = f.getRandValue([]string{categories[0], categories[1]})
		}

		dataVal = strings.Replace(dataVal, "{"+replace+"}", catValue, 1)
	}

	// Replace # with numbers
	dataVal = f.replaceWithNumbers(dataVal)

	// Replace ? with letters
	dataVal = f.replaceWithLetters(dataVal)

	return dataVal
}
