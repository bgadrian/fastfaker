package faker

import (
	"fmt"
	"testing"
)

func ExampleFaker_Bool() {
	Global.Seed(11)
	//it returns a Go boolean value
	fmt.Println(Global.Bool())
	// Output: false
}

func BenchmarkBool(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.Bool()
	}
}

func ExampleFaker_BoolText() {
	Global.Seed(11)
	//wil return a string value ("true" or "false")
	fmt.Println(Global.BoolText())
	fmt.Println(Global.BoolText())
	// Output: false
	// true
}

func BenchmarkBoolText(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.BoolText()
	}
}

func ExampleFaker_Binary() {
	Global.Seed(11)
	//will return a string value ("0" or "1")
	fmt.Println(Global.Binary())
	fmt.Println(Global.Binary())
	// Output: 0
	// 1
}

func BenchmarkBinary(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.Binary()
	}
}
