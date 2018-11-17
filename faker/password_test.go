package faker

import (
	"fmt"
	"testing"
)

func TestPassword(t *testing.T) {
	length := 10

	pass := Global.Password(true, true, true, true, true, length)

	if len(pass) != length {
		t.Error("Password length does not equal requested length")
	}

	// Test fully empty
	pass = Global.Password(false, false, false, false, false, length)
	if pass == "" {
		t.Error("Password should not be empty")
	}
}

func ExampleFaker_PasswordFull() {
	Global.Seed(42)
	fmt.Println(Global.PasswordFull())
	// Output: e+Hea2n9Xw0Q
}

func ExampleFaker_Password() {
	Global.Seed(11)
	fmt.Println(Global.Password(true, false, false, false, false, 32))
	fmt.Println(Global.Password(false, true, false, false, false, 32))
	fmt.Println(Global.Password(false, false, true, false, false, 32))
	fmt.Println(Global.Password(false, false, false, true, false, 32))
	fmt.Println(Global.Password(true, true, true, true, true, 32))
	fmt.Println(Global.Password(true, true, true, true, true, 4))
	// Output: vodnqxzsuptgehrzylximvylxzoywexw
	// ZSRQWJFJWCSTVGXKYKWMLIAFGFELFJRG
	// 61718615932495608398906260648432
	// @=-%%#$=-%+++&-#?*&?%&=%?+#@@-&?
	// EEP+wwpk 4lU-eHNXlJZ4n K9%v&TZ9e
	// j ?9X
}

func BenchmarkPassword(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.Password(true, true, true, true, true, 8)
	}
}
