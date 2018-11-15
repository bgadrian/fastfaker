package fastfaker

import (
	"fmt"
	"strconv"
	"testing"
)

func ExampleFaker_CreditCard() {
	Global.Seed(11)
	ccInfo := Global.CreditCard()
	fmt.Println(ccInfo.Type)
	fmt.Println(ccInfo.Number)
	fmt.Println(ccInfo.Exp)
	fmt.Println(ccInfo.Cvv)
	// Output: Visa
	// 6587271570245748
	// 03/25
	// 675
}

func BenchmarkCreditCard(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.CreditCard()
	}
}

func ExampleFaker_CreditCardType() {
	Global.Seed(11)
	fmt.Println(Global.CreditCardType())
	// Output: Visa
}

func BenchmarkCreditCardType(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.CreditCardType()
	}
}

func ExampleFaker_CreditCardNumber() {
	Global.Seed(11)
	fmt.Println(Global.CreditCardNumber())
	// Output: 4287271570245748
}

func BenchmarkCreditCardNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.CreditCardNumber()
	}
}

func ExampleFaker_CreditCardNumberLuhn() {
	Global.Seed(11)
	fmt.Println(Global.CreditCardNumberLuhn())
	// Output: 4007208855354357
}

func BenchmarkCreditCardNumberLuhn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.CreditCardNumberLuhn()
	}
}

func TestCreditCardNumberLuhn(t *testing.T) {
	Global.Seed(0)
	for i := 0; i < 100; i++ {
		cc := strconv.Itoa(Global.CreditCardNumberLuhn())
		if !Global.luhn(cc) {
			t.Errorf("not luhn valid: %s", cc)
		}
	}
}

func TestLuhn(t *testing.T) {
	// Lets make sure this card is invalid
	if Global.luhn("867gfsd5309") {
		t.Error("card should have failed")
	}
}

func ExampleFaker_CreditCardExp() {
	Global.Seed(11)
	fmt.Println(Global.CreditCardExp())
	// Output: 01/20
}

func BenchmarkCreditCardExp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.CreditCardExp()
	}
}

func ExampleFaker_CreditCardCvv() {
	Global.Seed(11)
	fmt.Println(Global.CreditCardCvv())
	// Output: 328
}

func BenchmarkCreditCardCvv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.CreditCardCvv()
	}
}
