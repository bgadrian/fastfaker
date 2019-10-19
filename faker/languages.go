package faker

// Language will return a random language
func (f *Faker) Language() string {
	return f.getRandValue([]string{"language", "long"})
}

// LanguageAbbreviation will return a random language abbreviation
func (f *Faker) LanguageAbbreviation() string {
	return f.getRandValue([]string{"language", "short"})
}

// ProgrammingLanguage will return a random programming language
func (f *Faker) ProgrammingLanguage() string {
	return f.getRandValue([]string{"language", "programming"})
}
