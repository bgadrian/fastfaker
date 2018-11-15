package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleFaker_Company() {
	Global.Seed(11)
	fmt.Println(Global.Company())
	// Output: Moen, Pagac and Wuckert
}

func BenchmarkCompany(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.Company()
	}
}

func TestCompany(t *testing.T) {
	for i := 0; i < 100; i++ {
		Global.Company()
	}
}

func ExampleFaker_CompanySuffix() {
	Global.Seed(11)
	fmt.Println(Global.CompanySuffix())
	// Output: Inc
}

func BenchmarkCompanySuffix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.CompanySuffix()
	}
}

func ExampleFaker_BuzzWord() {
	Global.Seed(11)
	fmt.Println(Global.BuzzWord())
	// Output: disintermediate
}

func BenchmarkBuzzWord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.BuzzWord()
	}
}

func ExampleFaker_BS() {
	Global.Seed(11)
	fmt.Println(Global.BS())
	// Output: front-end
}

func BenchmarkBS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.BS()
	}
}
