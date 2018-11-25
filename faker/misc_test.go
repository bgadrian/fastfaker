package faker

import (
	"fmt"
	"testing"
)

func TestReplaceWithNumbers(t *testing.T) {
	if "" != Global.replaceWithNumbers("") {
		t.Error("You should have gotten an empty string")
	}
}

func TestReplaceWithNumbersUnicode(t *testing.T) {
	for _, test := range []struct{ in, should string }{
		{"#界#世#", "5界7世8"},
		{"☺#☻☹#", "☺5☻☹7"},
		{"\x80#¼#語", "\x805¼7語"},
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
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.replaceWithNumbers("###☺#☻##☹##")
	}
}

func ExampleFaker_Numerify() {
	Global.Seed(11)
	fmt.Println(Global.Numerify("###-###-####"))
	// Output: 613-645-9948
}

func BenchmarkNumerify(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.Numerify("###-###-####")
	}
}

func TestClampInt(t *testing.T) {
	if clampInt(1, 10, 20) != 10 {
		t.Errorf("failed for 1, 10, 20")
	}

	if clampInt(100, 10, 20) != 20 {
		t.Errorf("failed for 1, 10, 20")
	}
}

func TestGeRandErrors(t *testing.T) {
	fastFaker := NewFastFaker()
	if fastFaker.getRandIntValue([]string{"XXX", "XXX"}) != 0 {
		t.Error("should return 0 when category not found")
	}
	if fastFaker.getRandValue([]string{"XXX", "XXX"}) != "" {
		t.Error("should return '' when category not found")
	}
}
