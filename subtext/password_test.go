package subtext

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/bgadrian/fastfaker/randomizer"
)

func TestPassword(t *testing.T) {
	length := 10
	text := New(randomizer.NewRandWrapper(false, nil))

	pass := text.Password(true, true, true, true, true, length)

	if len(pass) != length {
		t.Error("Password length does not equal requested length")
	}

	// Test fully empty
	pass = text.Password(false, false, false, false, false, length)
	if pass == "" {
		t.Error("Password should not be empty")
	}
}

func ExampleFakerText_PasswordFull() {
	text := New(randomizer.NewRandWrapper(false, rand.NewSource(42)))
	fmt.Println(text.PasswordFull())
	// Output: +eHQa02X9n
}

func ExampleFakerText_Password() {
	text := New(randomizer.NewRandWrapper(false, rand.NewSource(11)))

	fmt.Println(text.Password(true, false, false, false, false, 32))
	fmt.Println(text.Password(false, true, false, false, false, 32))
	fmt.Println(text.Password(false, false, true, false, false, 32))
	fmt.Println(text.Password(false, false, false, true, false, 32))
	fmt.Println(text.Password(true, true, true, true, true, 32))
	fmt.Println(text.Password(true, true, true, true, true, 4))
	// Output: vodnqxzsuptgehrzylximvylxzoywexw
	// ZSRQWJFJWCSTVGXKYKWMLIAFGFELFJRG
	// 61718615932495608398906260648432
	// @=-%%#$=-%+++&-#?*&?%&=%?+#@@-&?
	// EEP+wwpk 4lU-eHNXlJZ4n K9%v&TZ9e
	// j ?9X
}

func BenchmarkPassword(b *testing.B) {
	text := New(randomizer.NewRandWrapper(false, nil))
	for i := 0; i < b.N; i++ {
		text.Password(true, true, true, true, true, 8)
	}
}
