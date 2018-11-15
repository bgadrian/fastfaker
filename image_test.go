package fastfaker

import (
	"fmt"
	"testing"
)

func ExampleFaker_ImageURL() {
	Global.Seed(11)
	fmt.Println(Global.ImageURL(640, 480))
	// Output: http://lorempixel.com/640/480
}

func BenchmarkImageURL(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.ImageURL(640, 480)
	}
}
