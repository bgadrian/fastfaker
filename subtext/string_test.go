package subtext

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/bgadrian/fastfaker/randomizer"
)

func ExampleFakerText_Letter() {
	text := New(randomizer.NewRandWrapper(false, rand.NewSource(11)))
	fmt.Println(text.Letter())
	// Output: g
}

func BenchmarkLetter(b *testing.B) {
	text := New(randomizer.NewRandWrapper(false, rand.NewSource(42)))
	for i := 0; i < b.N; i++ {
		text.Letter()
	}
}

func ExampleFakerText_Digit() {
	text := New(randomizer.NewRandWrapper(false, rand.NewSource(11)))
	fmt.Println(text.Digit())
	// Output: 3
}

func BenchmarkDigit(b *testing.B) {
	text := New(randomizer.NewRandWrapper(false, rand.NewSource(42)))
	for i := 0; i < b.N; i++ {
		text.Digit()
	}
}

func ExampleFakerText_Lexify() {
	text := New(randomizer.NewRandWrapper(false, rand.NewSource(11)))
	fmt.Println(text.Lexify("?????"))
	// Output: gbrma
}

func BenchmarkLexify(b *testing.B) {
	text := New(randomizer.NewRandWrapper(false, rand.NewSource(42)))
	for i := 0; i < b.N; i++ {
		text.Lexify("??????")
	}
}

func ExampleFakerText_ShuffleStrings() {
	text := New(randomizer.NewRandWrapper(false, rand.NewSource(11)))
	strings := []string{"happy", "times", "for", "everyone", "have", "a", "good", "day"}
	text.ShuffleStrings(strings)
	fmt.Println(strings)
	// Output: [good everyone have for times a day happy]
}

func BenchmarkShuffleStrings(b *testing.B) {
	text := New(randomizer.NewRandWrapper(false, rand.NewSource(42)))
	for i := 0; i < b.N; i++ {
		text.ShuffleStrings([]string{"happy", "times", "for", "everyone", "have", "a", "good", "day"})
	}
}

func TestFakerText_RandString(t *testing.T) {
	text := New(randomizer.NewRandWrapper(false, rand.NewSource(42)))

	for _, test := range []struct {
		in     []string
		should string
	}{
		{[]string{}, ""},
		{nil, ""},
		{[]string{"a"}, "a"},
		{[]string{"a", "b", "c", "d", "e", "f"}, "f"},
	} {
		text.textRand.Seed(44)
		got := text.RandString(test.in)
		if got == test.should {
			continue
		}
		t.Errorf("for '%v' should '%s' got '%s'",
			test.in, test.should, got)
	}
}

func TestShuffleStrings(t *testing.T) {
	text := New(randomizer.NewRandWrapper(false, rand.NewSource(42)))
	//test for panics
	text.ShuffleStrings([]string{})
	text.ShuffleStrings([]string{"a"})
	text.ShuffleStrings(nil)

	a := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	b := make([]string, len(a))
	copy(b, a)
	text.ShuffleStrings(a)
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

func TestFakerText_Numerify(t *testing.T) {
	text := New(randomizer.NewRandWrapper(false, rand.NewSource(42)))
	if "" != text.Numerify("") {
		t.Error("You should have gotten an empty string")
	}
}

func TestReplaceWithNumbersUnicode(t *testing.T) {
	text := New(randomizer.NewRandWrapper(false, rand.NewSource(42)))
	for _, test := range []struct{ in, should string }{
		{"#界#世#", "8界8世5"},
		{"☺#☻☹#", "☺8☻☹8"},
		{"\x80#¼#語", "\x808¼8語"},
	} {
		text.textRand.Seed(42)
		got := text.Numerify(test.in)
		if got == test.should {
			continue
		}
		t.Errorf("for '%s' got '%s' should '%s'",
			test.in, got, test.should)
	}
}

func TestReplaceWithLetters(t *testing.T) {
	text := New(randomizer.NewRandWrapper(false, rand.NewSource(42)))
	if "" != text.Lexify("") {
		t.Error("You should have gotten an empty string")
	}
}

func BenchmarkReplaceWithNumbers(b *testing.B) {
	text := New(randomizer.NewRandWrapper(false, rand.NewSource(42)))
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		text.textRand.Seed(42)

		b.StartTimer()
		text.Numerify("###☺#☻##☹##")
		b.StopTimer()
	}
}

func ExampleFaker_Numerify() {
	text := New(randomizer.NewRandWrapper(false, rand.NewSource(11)))
	fmt.Println(text.Numerify("###-###-####"))
	// Output: 328-727-1570
}

func BenchmarkNumerify(b *testing.B) {
	text := New(randomizer.NewRandWrapper(false, rand.NewSource(42)))
	for i := 0; i < b.N; i++ {
		text.Numerify("###-###-####")
	}
}
