package gofakeit

import "strings"

// HackerPhrase will return a random hacker sentence
func (f *Faker) HackerPhrase() string {
	words := strings.Split(f.Generate(f.getRandValue([]string{"hacker", "phrase"})), " ")
	words[0] = strings.Title(words[0])
	return strings.Join(words, " ")
}

// HackerAbbreviation will return a random hacker abbreviation
func (f *Faker) HackerAbbreviation() string {
	return f.getRandValue([]string{"hacker", "abbreviation"})
}

// HackerAdjective will return a random hacker adjective
func (f *Faker) HackerAdjective() string {
	return f.getRandValue([]string{"hacker", "adjective"})
}

// HackerNoun will return a random hacker noun
func (f *Faker) HackerNoun() string {
	return f.getRandValue([]string{"hacker", "noun"})
}

// HackerVerb will return a random hacker verb
func (f *Faker) HackerVerb() string {
	return f.getRandValue([]string{"hacker", "verb"})
}

// HackerIngverb will return a random hacker ingverb
func (f *Faker) HackerIngverb() string {
	return f.getRandValue([]string{"hacker", "ingverb"})
}
