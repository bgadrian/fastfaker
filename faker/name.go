package faker

// Name will generate a random First and Last Name
func (f *Faker) Name() string {
	return f.getRandValue([]string{"person", "first"}) + " " + f.getRandValue([]string{"person", "last"})
}

// FirstName will generate a random first name
func (f *Faker) FirstName() string {
	return f.getRandValue([]string{"person", "first"})
}

// LastName will generate a random last name
func (f *Faker) LastName() string {
	return f.getRandValue([]string{"person", "last"})
}

// NamePrefix will generate a random name prefix
func (f *Faker) NamePrefix() string {
	return f.getRandValue([]string{"person", "prefix"})
}

// NameSuffix will generate a random name suffix
func (f *Faker) NameSuffix() string {
	return f.getRandValue([]string{"person", "suffix"})
}
