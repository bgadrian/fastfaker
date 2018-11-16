/* Package data contains all the raw data used by the Faker.
You can use it directly if you have your own selection criteria or a different random generator.
*/
package data

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

func TestCategories(t *testing.T) {
	var got, expected []string
	for k := range Categories() {
		got = append(got, k)
	}
	for k := range Data {
		expected = append(expected, k)
	}
	sort.Strings(got)
	sort.Strings(expected)
	if !reflect.DeepEqual(got, expected) {
		t.Error("Type arrays are not the same.\nExpected: ", expected, "\nGot: ", got)
	}
}

func TestGetRandValue(t *testing.T) {
	random := rand.New(rand.NewSource(42))
	for _, test := range [][]string{{"hacker", "abbreviation"}, {"person", "first"}} {
		if val, err := GetRandValue(random, test); val == "" || err != nil {
			t.Errorf("empty value for valid data {%s.%s}", test[0], test[1])
		}
	}
}

func TestGetRandValueFail(t *testing.T) {
	for _, test := range [][]string{nil, {}, {"not", "found"}, {"person", "notfound"}} {
		if val, err := GetRandValue(nil, test); val != "" || err == nil {
			t.Error("You should have gotten no value back")
		}
	}
}

func TestGetRandIntValue(t *testing.T) {
	random := rand.New(rand.NewSource(42))
	for _, test := range [][]string{{"status_code", "simple"}} {
		if _, err := GetRandIntValue(random, test); err != nil {
			t.Errorf("empty value for valid data {%s.%s}", test[0], test[1])
		}
	}
}

func TestGetRandIntValueFail(t *testing.T) {
	for _, test := range [][]string{nil, {}, {"not", "found"}, {"status_code", "notfound"}} {
		if _, err := GetRandIntValue(nil, test); err == nil {
			t.Error("You should have gotten no value back")
		}
	}
}

func Example() {
	//access the data directly
	fmt.Printf("I work at %s %s.\n", Company["buzzwords"][38], Company["suffix"][2])

	//OR provide a random generator
	random := rand.New(rand.NewSource(42))
	name, _ := GetRandValue(random, []string{"person", "first"})

	fmt.Printf("My name is %s.", name)
	// Output: I work at Innovative LLC.
	// My name is Jeromy.
}

func ExampleGetRandValue() {
	//you need to provide your own random generator
	random := rand.New(rand.NewSource(42))

	abbr, _ := GetRandValue(random, []string{"hacker", "abbreviation"})
	name, _ := GetRandValue(random, []string{"person", "first"})

	fmt.Printf("My name is %s, %s expert.", name, abbr)
	// Output: My name is Adriana, XSS expert.
}
