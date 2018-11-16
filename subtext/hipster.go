package subtext

import "github.com/bgadrian/fastfaker/data"

// HipsterWord will return a single hipster word
func (f *FakerText) HipsterWord() string {
	val, _ := data.GetRandValue(f.textRand, []string{"hipster", "word"})
	return val
}

// HipsterSentenceAvg will generate a random sentence containing 18 Hipster words.
func (f *FakerText) HipsterSentenceAvg() string {
	return f.HipsterSentence(18)
}

// HipsterSentence will generate a random sentence
func (f *FakerText) HipsterSentence(wordCount int) string {
	return f.sentence(wordCount, f.HipsterWord)
}

// HipsterParagraphAvg will generate a 4 paragraph separated by Unix new line \n text,
// each containing 32 Hipster words
func (f *FakerText) HipsterParagraphAvg() string {
	return f.HipsterParagraph(4, 2, 18, "\n")
}

// HipsterParagraph will generate a random paragraphGenerator
// Set Paragraph Count
// Set Sentence Count
// Set Word Count
// Set Paragraph Separator
func (f *FakerText) HipsterParagraph(paragraphCount int, sentenceCount int, wordCount int, separator string) string {
	return f.paragraphGenerator(
		paragrapOptions{paragraphCount, sentenceCount, wordCount, separator},
		f.HipsterSentence)
}
