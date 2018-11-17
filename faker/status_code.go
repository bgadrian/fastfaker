package faker

// SimpleStatusCode will generate a random simple status code
func (f *Faker) SimpleStatusCode() int {
	return f.getRandIntValue([]string{"status_code", "simple"})
}

// StatusCode will generate a random status code
func (f *Faker) StatusCode() int {
	return f.getRandIntValue([]string{"status_code", "general"})
}
