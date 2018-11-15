package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleFaker_MimeType() {
	Global.Seed(11)
	fmt.Println(Global.MimeType())
	// Output: application/dsptype
}

func BenchmarkMimeType(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.MimeType()
	}
}

func ExampleFaker_Extension() {
	Global.Seed(11)
	fmt.Println(Global.Extension())
	// Output: nes
}

func BenchmarkExtension(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.Extension()
	}
}
