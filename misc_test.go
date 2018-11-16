package fastfaker

import (
	"testing"
)

func TestReplaceWithNumbers(t *testing.T) {
	if "" != Global.replaceWithNumbers("") {
		t.Error("You should have gotten an empty string")
	}
}

func TestReplaceWithNumbersUnicode(t *testing.T) {
	for _, test := range []struct{ in, should string }{
		{"#界#世#", "8界8世5"},
		{"☺#☻☹#", "☺8☻☹8"},
		{"\x80#¼#語", "\x808¼8語"},
	} {
		Global.Seed(42)
		got := Global.replaceWithNumbers(test.in)
		if got == test.should {
			continue
		}
		t.Errorf("for '%s' got '%s' should '%s'",
			test.in, got, test.should)
	}
}

func TestReplaceWithLetters(t *testing.T) {
	if "" != Global.replaceWithLetters("") {
		t.Error("You should have gotten an empty string")
	}
}

func BenchmarkReplaceWithNumbers(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		Global.Seed(42)

		b.StartTimer()
		Global.replaceWithNumbers("###☺#☻##☹##")
		b.StopTimer()
	}
}
