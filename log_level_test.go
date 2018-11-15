package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleFaker_LogLevel() {
	Global.Seed(11)
	fmt.Println(Global.LogLevel("")) // This will also use general
	fmt.Println(Global.LogLevel("syslog"))
	fmt.Println(Global.LogLevel("apache"))
	// Output: error
	// debug
	// trace1-8
}

func BenchmarkLogLevel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.LogLevel("general")
	}
}
