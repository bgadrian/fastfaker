package faker

// Bool will generate a random boolean value
func (f *Faker) Bool() bool {
	return f.Intn(2) == 1
}

// BoolText will generate a random boolean value as text, "true" or "false".
// Can be used for JSON/YAML templates boolean values
func (f *Faker) BoolText() string {
	if f.Intn(2) == 1 {
		return "true"
	}

	return "false"
}

// Binary will generate a random binary value "0" or "1"
// Can be used for configs or generating binary values.
func (f *Faker) Binary() string {
	if f.Intn(2) == 1 {
		return "1"
	}

	return "0"
}
