package subtext

import (
	"bytes"
	"unicode"

	"github.com/bgadrian/fastfaker/data"
)

type paragrapOptions struct {
	paragraphCount int
	sentenceCount  int
	wordCount      int
	separator      string
}

const bytesPerWordEstimation = 6

type sentenceGenerator func(wordCount int) string
type wordGenerator func() string

// Word will generate a random word
func (f *FakerText) Word() string {
	val, _ := data.GetRandValue(f.textRand, []string{"lorem", "word"})
	return val
}

// Sentence will generate a random sentence, similar with Lorem Lipsum.
func (f *FakerText) Sentence(wordCount int) string {
	return f.sentence(wordCount, f.Word)
}

// SentenceAvg will generate a 18 word sentence, similar with Lorem Lipsum.
func (f *FakerText) SentenceAvg() string {
	return f.Sentence(18)
}

// ParagraphAvg will generate a 4 paragraph separated by Unix new line \n text,
// each containing 32 words, similar with Lorem Lipsum.
func (f *FakerText) ParagraphAvg() string {
	return f.Paragraph(4, 2, 18, "\n")
}

// Paragraph will generate a random paragraph, similar with Lorem Lipsum.
// Set Paragraph Count
// Set Sentence Count
// Set Word Count
// Set Paragraph Separator
func (f *FakerText) Paragraph(paragraphCount int, sentenceCount int, wordCount int, separator string) string {
	return f.paragraphGenerator(
		paragrapOptions{paragraphCount, sentenceCount, wordCount, separator},
		f.Sentence)
}

func (f *FakerText) sentence(wordCount int, word wordGenerator) string {
	if wordCount <= 0 {
		return ""
	}

	wordSeparator := ' '
	sentence := bytes.Buffer{}
	sentence.Grow(wordCount * bytesPerWordEstimation)

	for i := 0; i < wordCount; i++ {
		//TODO optimize, collect more words in 1 pass, make GetRandValueN()
		word := word()
		if i == 0 {
			runes := []rune(word)
			runes[0] = unicode.ToTitle(runes[0])
			word = string(runes)
		}
		sentence.WriteString(word)
		if i < wordCount-1 {
			sentence.WriteRune(wordSeparator)
		}
	}
	sentence.WriteRune('.')
	return sentence.String()
}

func (f *FakerText) paragraphGenerator(opts paragrapOptions, sentecer sentenceGenerator) string {
	if opts.paragraphCount <= 0 || opts.sentenceCount <= 0 || opts.wordCount <= 0 {
		return ""
	}

	//to avoid making Go 1.10 dependency, we cannot use strings.Builder
	paragraphs := bytes.Buffer{}
	//we presume the length
	paragraphs.Grow(opts.paragraphCount * opts.sentenceCount * opts.wordCount * bytesPerWordEstimation)
	wordSeparator := ' '

	for i := 0; i < opts.paragraphCount; i++ {
		for e := 0; e < opts.sentenceCount; e++ {
			paragraphs.WriteString(sentecer(opts.wordCount))
			if e < opts.sentenceCount-1 {
				paragraphs.WriteRune(wordSeparator)
			}
		}

		if i < opts.paragraphCount-1 {
			paragraphs.WriteString(opts.separator)
		}
	}

	return paragraphs.String()
}
