package faker

import (
	"fmt"
	"testing"
)

func ExampleFaker_SimpleStatusCode() {
	Global.Seed(11)
	fmt.Println(Global.SimpleStatusCode())
	// Output: 200
}

func BenchmarkSimpleStatusCode(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.SimpleStatusCode()
	}
}

func ExampleFaker_StatusCode() {
	Global.Seed(11)
	fmt.Println(Global.StatusCode())
	// Output: 404
}

func BenchmarkStatusCode(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.StatusCode()
	}
}
