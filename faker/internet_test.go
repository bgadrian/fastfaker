package faker

import (
	"fmt"
	"strings"
	"testing"
)

func ExampleFaker_DomainName() {
	Global.Seed(11)
	fmt.Println(Global.DomainName())
	// Output: centraltarget.biz
}

func BenchmarkDomainName(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.DomainName()
	}
}

func ExampleFaker_DomainSuffix() {
	Global.Seed(11)
	fmt.Println(Global.DomainSuffix())
	// Output: org
}

func BenchmarkDomainSuffix(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.DomainSuffix()
	}
}
func TestFaker_URL(t *testing.T) {
	Global.Seed(13)
	if !strings.Contains(Global.URL(), "https") {
		t.Errorf("some URLs must be https too")
	}

	Global.Seed(12)
	if !strings.Contains(Global.URL(), "http") {
		t.Errorf("some URLs must be http too")
	}
}

func ExampleFaker_URL() {
	Global.Seed(11)
	fmt.Println(Global.URL())
	// Output: https://www.nationalseamless.net/iterate/streamline/systems
}

func BenchmarkURL(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.URL()
	}
}

func ExampleFaker_HTTPMethod() {
	Global.Seed(11)
	fmt.Println(Global.HTTPMethod())
	// Output: HEAD
}

func BenchmarkHTTPMethod(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.HTTPMethod()
	}
}

func ExampleFaker_IPv4Address() {
	Global.Seed(11)
	fmt.Println(Global.IPv4Address())
	// Output: 222.83.191.222
}

func BenchmarkIPv4Address(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.IPv4Address()
	}
}

func ExampleFaker_IPv6Address() {
	Global.Seed(11)
	fmt.Println(Global.IPv6Address())
	// Output: 2001:cafe:8898:ee17:bc35:9064:5866:d019
}

func BenchmarkIPv6Address(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.IPv6Address()
	}
}

func ExampleMacAddress() {
	Global.Seed(11)
	fmt.Println(Global.MacAddress())
	// Output: e1:74:cb:01:77:91
}

func BenchmarkMacAddress(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.MacAddress()
	}
}
func ExampleFaker_Username() {
	Global.Seed(11)
	fmt.Println(Global.Username())
	// Output: Daniel1364
}

func BenchmarkUsername(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.Username()
	}
}
