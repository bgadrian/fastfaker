package fastfaker

import (
	"fmt"
	"testing"
)

func ExampleFaker_Bool() {
	Global.Seed(11)
	fmt.Println(Global.Bool())
	// Output: false
}

func BenchmarkBool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.Bool()
	}
}
