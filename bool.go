package fastfaker

// Bool will generate a random boolean value
func (f *Faker) Bool() bool {
	if f.Number(0, 1) == 1 {
		return true
	}

	return false
}
