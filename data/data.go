/*
Package data contains all the raw data used by the Faker, like Names, Streets, Countries and Companies.
You can use it directly if you have your own selection criteria or a different random generator.
*/
package data

import "github.com/pkg/errors"

// NotFound is triggered when a category of Data is not found
const NotFound = "not found"

// Intn is a part of the randomizer package. It is used to select an item from the database.
type Intn interface {
	Intn(n int) int
}

// Data consists of the main set of fake information
var Data = map[string]map[string][]string{
	"person":    Person,
	"contact":   Contact,
	"address":   Address,
	"company":   Company,
	"job":       Job,
	"lorem":     Lorem,
	"internet":  Internet,
	"file":      Files,
	"color":     Colors,
	"computer":  Computer,
	"payment":   Payment,
	"hipster":   Hipster,
	"beer":      Beer,
	"hacker":    Hacker,
	"currency":  Currency,
	"log_level": LogLevels,
	"timezone":  TimeZone,
	"vehicle":   Vehicle,
}

// IntData consists of the main set of fake information (integer only)
var IntData = map[string]map[string][]int{
	"status_code": StatusCodes,
}

// Check is used to see if a [category, subcategory] exists in Data.
func Check(dataVal []string) bool {
	var checkOk bool

	if len(dataVal) == 2 {
		_, checkOk = Data[dataVal[0]]
		if checkOk {
			_, checkOk = Data[dataVal[0]][dataVal[1]]
		}
	}

	return checkOk
}

// Check if in lib
func intDataCheck(dataVal []string) bool {
	if len(dataVal) != 2 {
		return false
	}

	_, checkOk := IntData[dataVal[0]]
	if checkOk {
		_, checkOk = IntData[dataVal[0]][dataVal[1]]
	}

	return checkOk
}

// GetRandValue from a [category, subcategory]. Returns error if not found.
func GetRandValue(f Intn, dataVal []string) (string, error) {
	if !Check(dataVal) {
		return "", errors.New(NotFound)
	}
	return Data[dataVal[0]][dataVal[1]][f.Intn(len(Data[dataVal[0]][dataVal[1]]))], nil
}

// GetRandIntValue Value from a [category, subcategory]. Returns error if not found.
func GetRandIntValue(f Intn, dataVal []string) (int, error) {
	if !intDataCheck(dataVal) {
		return 0, errors.New(NotFound)
	}
	return IntData[dataVal[0]][dataVal[1]][f.Intn(len(IntData[dataVal[0]][dataVal[1]]))], nil
}

// Categories will return a map string array of available data categories and sub categories
func Categories() map[string][]string {
	types := make(map[string][]string)
	for category, subCategoriesMap := range Data {
		subCategories := make([]string, 0)
		for subType := range subCategoriesMap {
			subCategories = append(subCategories, subType)
		}
		types[category] = subCategories
	}
	return types
}
