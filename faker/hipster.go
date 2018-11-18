package faker

import "github.com/bgadrian/fastfaker/data"

// HipsterWord will return a single hipster word
func (f *Faker) HipsterWord() string {
	return f.getRandValue([]string{"hipster", "word"})
}

// HipsterSentenceAvg will generate a random sentence containing 18 Hipster words.
func (f *Faker) HipsterSentenceAvg() string {
	return f.HipsterSentence(18)
}

// HipsterSentence will generate a random sentence
func (f *Faker) HipsterSentence(wordCount int) string {
	wordSetCache, err := data.NewDataListCache(f, []string{"hipster", "word"})
	if err != nil {
		return ""
	}
	return f.sentence(wordCount, wordSetCache, nil).String()
}

// HipsterParagraphAvg will generate a 4 paragraph separated by Unix new line \n text,
// each containing 32 Hipster words
func (f *Faker) HipsterParagraphAvg() string {
	return f.HipsterParagraph(4, 2, 18, "\n")
}

// HipsterParagraph will generate a random paragraphGenerator
// Set Paragraph Count
// Set Sentence Count
// Set Word Count
// Set Paragraph Separator
func (f *Faker) HipsterParagraph(paragraphCount int, sentenceCount int, wordCount int, separator string) string {
	wordSetCache, err := data.NewDataListCache(f, []string{"hipster", "word"})
	if err != nil {
		return ""
	}
	return f.paragraphGenerator(
		paragrapOptions{paragraphCount, sentenceCount, wordCount, separator},
		wordSetCache)
}
