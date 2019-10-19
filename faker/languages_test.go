package faker

import (
	"fmt"
	"testing"
)

func ExampleLanguage() {
	Global.Seed(11)
	fmt.Println(Global.Language())
	// Output: Kazakh
}

func BenchmarkLanguage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.Language()
	}
}

func ExampleLanguageAbbreviation() {
	Global.Seed(11)
	fmt.Println(Global.LanguageAbbreviation())
	// Output: kk
}

func BenchmarkLanguageAbbreviation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.LanguageAbbreviation()
	}
}

func ExampleProgrammingLanguage() {
	Global.Seed(464)
	fmt.Println(Global.ProgrammingLanguage())
	// Output: Go
}

func BenchmarkProgrammingLanguage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.ProgrammingLanguage()
	}
}
