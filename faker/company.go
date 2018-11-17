package faker

// Company will generate a random company name string
func (f *Faker) Company() (company string) {
	switch randInt := f.Number(1, 3); randInt {
	case 1:
		company = f.LastName() + ", " + f.LastName() + " and " + f.LastName()
	case 2:
		company = f.LastName() + "-" + f.LastName()
	case 3:
		company = f.LastName() + " " + f.CompanySuffix()
	}

	return
}

// CompanySuffix will generate a random company suffix string
func (f *Faker) CompanySuffix() string {
	return f.getRandValue([]string{"company", "suffix"})
}

// BuzzWord will generate a random company buzz word string
func (f *Faker) BuzzWord() string {
	return f.getRandValue([]string{"company", "buzzwords"})
}

// BS will generate a random company bs string
func (f *Faker) BS() string {
	return f.getRandValue([]string{"company", "bs"})
}
