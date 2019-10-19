package faker

import (
	"fmt"
	"testing"
)

func ExampleMap() {
	Global.Seed(11)
	fmt.Println(Global.Map())
	// Output: map[accusantium: amet: aut:-3430133205295092491 et: nisi: non:-61868737515613618 qui: repellat:-3980178883679362219]
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
