package fastfaker

// HipsterWord will return a single hipster word
func (f *Faker) HipsterWord() string {
	return f.getRandValue([]string{"hipster", "word"})
}

// HipsterSentence will generate a random sentence
func (f *Faker) HipsterSentence(wordCount int) string {
	return f.sentence(wordCount, f.HipsterWord)
}

// HipsterParagraph will generate a random paragraphGenerator
// Set Paragraph Count
// Set Sentence Count
// Set Word Count
// Set Paragraph Separator
func (f *Faker) HipsterParagraph(paragraphCount int, sentenceCount int, wordCount int, separator string) string {
	return f.paragraphGenerator(
		paragrapOptions{paragraphCount, sentenceCount, wordCount, separator},
		f.HipsterSentence)
}
