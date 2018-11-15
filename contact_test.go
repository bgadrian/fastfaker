package fastfaker

import (
	"fmt"
	"testing"
)

func ExampleFaker_Contact() {
	Global.Seed(11)
	contact := Global.Contact()
	fmt.Println(contact.Phone)
	fmt.Println(contact.Email)
	// Output: 3287271570
	// santinostanton@carroll.biz
}

func BenchmarkContact(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.Contact()
	}
}

func ExampleFaker_Phone() {
	Global.Seed(11)
	fmt.Println(Global.Phone())
	// Output: 3287271570
}

func BenchmarkPhone(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.Phone()
	}
}

func ExampleFaker_PhoneFormatted() {
	Global.Seed(11)
	fmt.Println(Global.PhoneFormatted())
	// Output: 287-271-5702
}

func BenchmarkPhoneFormatted(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.PhoneFormatted()
	}
}

func ExampleFaker_Email() {
	Global.Seed(11)
	fmt.Println(Global.Email())
	// Output: markusmoen@pagac.net
}

func BenchmarkEmail(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.Email()
	}
}
