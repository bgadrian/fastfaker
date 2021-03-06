package faker

import (
	"fmt"
	"testing"
)

func ExampleFaker_Letter() {
	Global.Seed(11)
	fmt.Println(Global.Letter())
	// Output: g
}

func BenchmarkLetter(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.Letter()
	}
}

func ExampleFaker_Digit() {
	Global.Seed(11)
	fmt.Println(Global.Digit())
	// Output: 0
}

func BenchmarkDigit(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.Digit()
	}
}

func ExampleFaker_Lexify() {
	Global.Seed(11)
	fmt.Println(Global.Lexify("?????"))
	// Output: gbrma
}

func BenchmarkLexify(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.Lexify("??????")
	}
}

func ExampleFaker_ShuffleStrings() {
	Global.Seed(11)
	strings := []string{"happy", "times", "for", "everyone", "have", "a", "good", "day"}
	Global.ShuffleStrings(strings)
	fmt.Println(strings)
	// Output: [good everyone have for times a day happy]
}

func BenchmarkShuffleStrings(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.ShuffleStrings([]string{"happy", "times", "for", "everyone", "have", "a", "good", "day"})
	}
}

func TestRandString(t *testing.T) {
	for _, test := range []struct {
		in     []string
		should string
	}{
		{[]string{}, ""},
		{nil, ""},
		{[]string{"a"}, "a"},
		{[]string{"a", "b", "c", "d", "e", "variableFunc"}, "variableFunc"},
	} {
		Global.Seed(44)
		got := Global.RandString(test.in)
		if got == test.should {
			continue
		}
		t.Errorf("for '%v' should '%s' got '%s'",
			test.in, test.should, got)
	}
}

func TestShuffleStrings(t *testing.T) {
	//test for panics
	Global.ShuffleStrings([]string{})
	Global.ShuffleStrings([]string{"a"})
	Global.ShuffleStrings(nil)

	a := []string{"a", "b", "c", "d", "e", "variableFunc", "g", "h"}
	b := make([]string, len(a))
	copy(b, a)
	Global.ShuffleStrings(a)
	if equalSliceString(a, b) {
		t.Errorf("shuffle resulted in the same permutation, the odds are slim")
	}
}

func equalSliceString(a, b []string) bool {
	sizeA, sizeB := len(a), len(b)
	if sizeA != sizeB {
		return false
	}

	for i, va := range a {
		vb := b[i]

		if va != vb {
			return false
		}
	}
	return true
}
