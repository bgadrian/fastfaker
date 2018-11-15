package fastfaker

import (
	"reflect"
	"sort"
	"testing"

	"github.com/bgadrian/fastfaker/data"
)

func TestRandIntRange(t *testing.T) {
	if Global.randIntRange(5, 5) != 5 {
		t.Error("You should have gotten 5 back")
	}
}

func TestGetRandValueFail(t *testing.T) {
	for _, test := range [][]string{nil, {}, {"not", "found"}, {"person", "notfound"}} {
		if Global.getRandValue(test) != "" {
			t.Error("You should have gotten no value back")
		}
	}
}

func TestGetRandIntValueFail(t *testing.T) {
	for _, test := range [][]string{nil, {}, {"not", "found"}, {"status_code", "notfound"}} {
		if Global.getRandIntValue(test) != 0 {
			t.Error("You should have gotten no value back")
		}
	}
}

func TestRandFloat32RangeSame(t *testing.T) {
	if Global.Float32Range(5.0, 5.0) != 5.0 {
		t.Error("You should have gotten 5.0 back")
	}
}

func TestRandFloat64RangeSame(t *testing.T) {
	if Global.Float64Range(5.0, 5.0) != 5.0 {
		t.Error("You should have gotten 5.0 back")
	}
}

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

func TestCategories(t *testing.T) {
	var got, expected []string
	for k := range Categories() {
		got = append(got, k)
	}
	for k := range data.Data {
		expected = append(expected, k)
	}
	sort.Strings(got)
	sort.Strings(expected)
	if !reflect.DeepEqual(got, expected) {
		t.Error("Type arrays are not the same.\nExpected: ", expected, "\nGot: ", got)
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
