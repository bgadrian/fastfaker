package faker

var (
	mapRandStrings = []string{"{word}", "{bs}", "{jobtitle}", "{name}", "{street}, {city}, {state} {zip}"}
	mapRandTypes   = []string{"string", "int", "float", "slice"}
)

// Map will generate a random set of map data
func (f *Faker) Map() map[string]interface{} {
	size := int(f.Int32Range(3, 10))
	m := make(map[string]interface{}, size)

	for i := 0; i < size; i++ {
		switch f.RandString(mapRandTypes) {
		case "int":
			m[f.Word()] = f.Int64()
		case "float":
			m[f.Word()] = f.Float64()
		case "slice":
			m[f.Word()] = f.SliceStrings()
		default:
			m[f.Word()] = f.Template(f.RandString(mapRandStrings))
		}
	}

	return m
}

func (f *Faker) SliceStrings() []string {
	count := f.Int32Range(3, 10)
	result := make([]string, count)
	for i := range result {
		result[i] = f.Word()
	}
	return result
}
