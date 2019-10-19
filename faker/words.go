package faker

import (
	"bytes"
	"strings"
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

// Word will generate a random word
func (f *Faker) Word() string {
	return f.getRandValue([]string{"lorem", "word"})
}

// Sentence will generate a random sentence, similar with Lorem Lipsum.
func (f *Faker) Sentence(wordCount int) string {
	wordSetCache, err := data.NewDataListCache(f, []string{"lorem", "word"})
	if err != nil {
		errorLogger.Printf("(Sentence) %s\n", err)
		return ""
	}
	return f.sentence(wordCount, wordSetCache, nil).String()
}

// SentenceAvg will generate a 18 word sentence, similar with Lorem Lipsum.
func (f *Faker) SentenceAvg() string {
	return f.Sentence(18)
}

// ParagraphAvg will generate a 4 paragraph separated by Unix new line \n text,
// each containing 32 words, similar with Lorem Lipsum.
func (f *Faker) ParagraphAvg() string {
	return f.Paragraph(4, 2, 18, "\n")
}

// Paragraph will generate a random paragraph, similar with Lorem Lipsum.
// Set Paragraph Count
// Set Sentence Count
// Set Word Count
// Set Paragraph Separator
func (f *Faker) Paragraph(paragraphCount int, sentenceCount int, wordCount int, separator string) string {
	wordSetCache, err := data.NewDataListCache(f, []string{"lorem", "word"})
	if err != nil {
		errorLogger.Printf("(Paragraph) %s\n", err)
		return ""
	}
	return f.paragraphGenerator(
		paragrapOptions{paragraphCount, sentenceCount, wordCount, separator},
		wordSetCache)
}

func (f *Faker) sentence(wordCount int, word data.SetCache, sentence *bytes.Buffer) *bytes.Buffer {
	if wordCount <= 0 {
		return &bytes.Buffer{}
	}

	if sentence == nil {
		sentence = &bytes.Buffer{}
		sentence.Grow(wordCount * bytesPerWordEstimation)
	}

	for i := 0; i < wordCount; i++ {
		word := word.GetRandValue()
		if i == 0 {
			runes := []rune(word)
			runes[0] = unicode.ToTitle(runes[0])
			word = string(runes)
		}
		//nolint buffer always return nil error
		sentence.WriteString(word)
		if i < wordCount-1 {
			//nolint buffer always return nil error
			sentence.WriteRune(' ')
		}
	}
	//nolint buffer always return nil error
	sentence.WriteRune('.')
	return sentence
}

func (f *Faker) paragraphGenerator(opts paragrapOptions, word data.SetCache) string {
	if opts.paragraphCount <= 0 || opts.sentenceCount <= 0 || opts.wordCount <= 0 {
		return ""
	}

	//to avoid making Go 1.10 dependency, we cannot use strings.Builder
	paragraphs := &bytes.Buffer{}
	//we presume the length
	paragraphs.Grow(opts.paragraphCount * opts.sentenceCount * opts.wordCount * bytesPerWordEstimation)
	wordSeparator := ' '

	for i := 0; i < opts.paragraphCount; i++ {
		for e := 0; e < opts.sentenceCount; e++ {
			f.sentence(opts.wordCount, word, paragraphs)
			if e < opts.sentenceCount-1 {
				//nolint buffer always return nil error
				paragraphs.WriteRune(wordSeparator)
			}
		}

		if i < opts.paragraphCount-1 {
			//nolint buffer always return nil error
			paragraphs.WriteString(opts.separator)
		}
	}

	return paragraphs.String()
}

// Question will return a random question
func (f *Faker) Question() string {
	return strings.Replace(f.HipsterSentence(f.Number(3, 10)), ".", "?", 1)
}

// Quote will return a random quote from a random person
func (f *Faker) Quote() string {
	return `"` + f.HipsterSentence(f.Number(3, 10)) + `" - ` + f.FirstName() + " " + f.LastName()
}
