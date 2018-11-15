package fastfaker

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
	color := make([]byte, 6)
	hashQuestion := []byte("?#")
	for i := 0; i < 6; i++ {
		color[i] = hashQuestion[f.Intn(2)]
	}

	return "#" + f.replaceWithLetters(f.replaceWithNumbers(string(color)))

	// color := ""
	// for i := 1; i <= 6; i++ {
	// 	color += RandString([]string{"?", "#"})
	// }

	// // Replace # with number
	// color = replaceWithNumbers(color)

	// // Replace ? with letter
	// for strings.Count(color, "?") > 0 {
	// 	color = strings.Replace(color, "?", RandString(letters), 1)
	// }

	// return "#" + color
}

// RGBColor will generate a random int slice color
func (f *Faker) RGBColor() []int {
	return []int{f.randIntRange(0, 255), f.randIntRange(0, 255), f.randIntRange(0, 255)}
}
