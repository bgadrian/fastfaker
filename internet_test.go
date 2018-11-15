package fastfaker

import (
	"fmt"
	"testing"
)

func ExampleFaker_DomainName() {
	Global.Seed(11)
	fmt.Println(Global.DomainName())
	// Output: centraltarget.biz
}

func BenchmarkDomainName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.DomainName()
	}
}

func ExampleFaker_DomainSuffix() {
	Global.Seed(11)
	fmt.Println(Global.DomainSuffix())
	// Output: org
}

func BenchmarkDomainSuffix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.DomainSuffix()
	}
}

func ExampleFaker_URL() {
	Global.Seed(11)
	fmt.Println(Global.URL())
	// Output: https://www.nationalseamless.net/iterate/streamline/systems
}

func BenchmarkURL(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.URL()
	}
}

func ExampleFaker_HTTPMethod() {
	Global.Seed(11)
	fmt.Println(Global.HTTPMethod())
	// Output: HEAD
}

func BenchmarkHTTPMethod(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.HTTPMethod()
	}
}

func ExampleFaker_IPv4Address() {
	Global.Seed(11)
	fmt.Println(Global.IPv4Address())
	// Output: 222.83.191.222
}

func BenchmarkIPv4Address(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.IPv4Address()
	}
}

func ExampleFaker_IPv6Address() {
	Global.Seed(11)
	fmt.Println(Global.IPv6Address())
	// Output: 2001:cafe:8898:ee17:bc35:9064:5866:d019
}

func BenchmarkIPv6Address(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.IPv6Address()
	}
}

func ExampleFaker_Username() {
	Global.Seed(11)
	fmt.Println(Global.Username())
	// Output: Daniel2872
}

func BenchmarkUsername(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.Username()
	}
}
