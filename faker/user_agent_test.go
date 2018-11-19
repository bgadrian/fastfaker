package faker

import (
	"fmt"
	"testing"
)

func ExampleFaker_UserAgent() {
	Global.Seed(11)
	fmt.Println(Global.UserAgent())
	// Output: Mozilla/5.0 (Windows NT 5.0) AppleWebKit/5362 (KHTML, like Gecko) Chrome/37.0.834.0 Mobile Safari/5362
}

func BenchmarkUserAgent(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.UserAgent()
	}
}

func TestUserAgent(t *testing.T) {
	for i := 0; i < 100; i++ {
		Global.UserAgent()
	}
}

func ExampleFaker_ChromeUserAgent() {
	Global.Seed(11)
	fmt.Println(Global.ChromeUserAgent())
	// Output: Mozilla/5.0 (X11; Linux i686) AppleWebKit/5312 (KHTML, like Gecko) Chrome/39.0.836.0 Mobile Safari/5312
}

func BenchmarkChromeUserAgent(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.ChromeUserAgent()
	}
}

func ExampleFaker_FirefoxUserAgent() {
	Global.Seed(11)
	fmt.Println(Global.FirefoxUserAgent())
	// Output: Mozilla/5.0 (Macintosh; U; PPC Mac OS X 10_8_3 rv:7.0) Gecko/1989-07-01 Firefox/37.0
}

func BenchmarkFirefoxUserAgent(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.FirefoxUserAgent()
	}
}

func ExampleFaker_SafariUserAgent() {
	Global.Seed(11)
	fmt.Println(Global.SafariUserAgent())
	// Output: Mozilla/5.0 (iPad; CPU OS 8_3_2 like Mac OS X; en-US) AppleWebKit/531.15.6 (KHTML, like Gecko) Version/4.0.5 Mobile/8B120 Safari/6531.15.6
}

func BenchmarkSafariUserAgent(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.SafariUserAgent()
	}
}

func ExampleFaker_OperaUserAgent() {
	Global.Seed(11)
	fmt.Println(Global.OperaUserAgent())
	// Output: Opera/8.39 (Macintosh; U; PPC Mac OS X 10_8_7; en-US) Presto/2.9.335 Version/10.00
}

func BenchmarkOperaUserAgent(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.OperaUserAgent()
	}
}
