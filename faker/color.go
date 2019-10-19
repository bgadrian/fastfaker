package faker

import "strings"

// Color will generate a random color string
func (f *Faker) Color() string {
	return f.getRandValue([]string{"color", "full"})
}

// SafeColor will generate a random safe color string
func (f *Faker) SafeColor() string {
	return f.getRandValue([]string{"color", "safe"})
}

// HexColor will generate a random hexadecimal color string
func (f *Faker) HexColor() string {
	b := strings.Builder{}
	b.Grow(7)
	b.WriteRune('#')
	for i := 0; i < 6; i++ {
		b.WriteByte(f.randHex())
	}
	return b.String()
}

// RGBColor will generate a random int slice color
func (f *Faker) RGBColor() []int {
	return []int{f.Number(0, 255), f.Number(0, 255), f.Number(0, 255)}
}
