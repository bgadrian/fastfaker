package faker

import (
	"fmt"
	"testing"
)

func TestUUID(t *testing.T) {
	id := Global.UUID()

	if len(id) != 36 {
		t.Error("unique length does not equal requested length")
	}
}

func ExampleFaker_UUID() {
	Global.Seed(11)
	fmt.Println(Global.UUID())
	// Output: 590c1440-9888-45b0-bd51-a817ee07c3f2
}

func BenchmarkUUID(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.UUID()
	}
}
