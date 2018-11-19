package faker

import (
	"fmt"
	"testing"
)

func ExampleFaker_Currency() {
	Global.Seed(11)
	currency := Global.Currency()
	fmt.Printf("%s - %s", currency.Short, currency.Long)
	// Output: IQD - Iraq Dinar
}

func BenchmarkCurrency(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.Currency()
	}
}

func ExampleFaker_CurrencyShort() {
	Global.Seed(11)
	fmt.Println(Global.CurrencyShort())
	// Output: IQD
}

func BenchmarkCurrencyShort(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.CurrencyShort()
	}
}

func ExampleFaker_CurrencyLong() {
	Global.Seed(11)
	fmt.Println(Global.CurrencyLong())
	// Output: Iraq Dinar
}

func BenchmarkCurrencyLong(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.CurrencyLong()
	}
}

func ExampleFaker_Price() {
	Global.Seed(11)
	fmt.Printf("%.2f", Global.Price(0.8618, 1000))
	// Output: 92.26
}

func BenchmarkPrice(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.Price(0, 1000)
	}
}
