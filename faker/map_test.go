package faker

import (
	"fmt"
	"sort"
	"testing"
)

func ExampleMap() {
	Global.Seed(11)
	randomMap := Global.Map()

	//Caveat, older Go versions printed maps elements in random order, so we have to order them ourselves
	items := []string{}
	for key := range randomMap {
		items = append(items, key)
	}
	sort.Strings(items)
	for _, key := range items {
		fmt.Printf("%s: %#v\n", key, randomMap[key])
	}
	// Output: aut: -4163344001092914963
	//exercitationem: "consequatur"
	//harum: 5.235169171532564e+307
	//non: 1.785634529370745e+308
	//quas: -3673523645516654914
	//repellat: []string{"quidem", "nisi", "quo", "qui", "voluptatum", "accusantium", "quisquam"}
	//similique: "58302 Drivehaven, Alvenabury, Arizona 61321"
}

func TestMap(t *testing.T) {
	for i := 0; i < 100; i++ {
		Global.Map()
	}
}

func BenchmarkMap(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.Map()
	}
}
