package faker

import (
	"fmt"
	"testing"
)

func ExampleFaker_AvatarURL() {
	Global.Seed(42)
	fmt.Println(Global.AvatarURL())
	// Output: http://pipsum.com/80x80.jpg
}

func BenchmarkAvatarURL(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.AvatarURL()
	}
}

func ExampleFaker_ImageURL() {
	Global.Seed(11)
	fmt.Println(Global.ImageURL(640, 480))
	// Output: http://pipsum.com/640x480.jpg
}

func BenchmarkImageURL(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.ImageURL(640, 480)
	}
}
