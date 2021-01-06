package gofakeit

import (
	"bytes"
	"errors"
	rand "math/rand"
	"strings"
	"unicode"
)

type paragrapOptions struct {
	paragraphCount int
	sentenceCount  int
	wordCount      int
	separator      string
}

const bytesPerWordEstimation = 6

type sentenceGenerator func(r *rand.Rand, wordCount int) string
type wordGenerator func(r *rand.Rand) string

// Noun will generate a random noun
func Noun() string { return noun(globalFaker.Rand) }

// Noun will generate a random noun
func (f *Faker) Noun() string { return noun(f.Rand) }

func noun(r *rand.Rand) string { return getRandValue(r, []string{"word", "noun"}) }

// Verb will generate a random verb
func Verb() string { return verb(globalFaker.Rand) }

// Verb will generate a random verb
func (f *Faker) Verb() string { return verb(f.Rand) }

func verb(r *rand.Rand) string { return getRandValue(r, []string{"word", "verb"}) }

// Adverb will generate a random adverb
func Adverb() string { return adverb(globalFaker.Rand) }

// Adverb will generate a random adverb
func (f *Faker) Adverb() string { return adverb(f.Rand) }

func adverb(r *rand.Rand) string { return getRandValue(r, []string{"word", "adverb"}) }

// Preposition will generate a random preposition
func Preposition() string { return preposition(globalFaker.Rand) }

// Preposition will generate a random preposition
func (f *Faker) Preposition() string { return preposition(f.Rand) }

func preposition(r *rand.Rand) string { return getRandValue(r, []string{"word", "preposition"}) }

// Adjective will generate a random adjective
func Adjective() string { return adjective(globalFaker.Rand) }

// Adjective will generate a random adjective
func (f *Faker) Adjective() string { return adjective(f.Rand) }

func adjective(r *rand.Rand) string { return getRandValue(r, []string{"word", "adjective"}) }

// Word will generate a random word
func Word() string { return word(globalFaker.Rand) }

// Word will generate a random word
func (f *Faker) Word() string { return word(f.Rand) }

func word(r *rand.Rand) string {
	if boolFunc(r) {
		return getRandValue(r, []string{"word", "noun"})
	}

	return getRandValue(r, []string{"word", "verb"})
}

// Sentence will generate a random sentence
func Sentence(wordCount int) string { return sentence(globalFaker.Rand, wordCount) }

// Sentence will generate a random sentence
func (f *Faker) Sentence(wordCount int) string { return sentence(f.Rand, wordCount) }

func sentence(r *rand.Rand, wordCount int) string {
	return sentenceGen(r, wordCount, word)
}

// Paragraph will generate a random paragraphGenerator
func Paragraph(paragraphCount int, sentenceCount int, wordCount int, separator string) string {
	return paragraph(globalFaker.Rand, paragraphCount, sentenceCount, wordCount, separator)
}

// Paragraph will generate a random paragraphGenerator
func (f *Faker) Paragraph(paragraphCount int, sentenceCount int, wordCount int, separator string) string {
	return paragraph(f.Rand, paragraphCount, sentenceCount, wordCount, separator)
}

func paragraph(r *rand.Rand, paragraphCount int, sentenceCount int, wordCount int, separator string) string {
	return paragraphGen(r, paragrapOptions{paragraphCount, sentenceCount, wordCount, separator}, sentence)
}

func sentenceGen(r *rand.Rand, wordCount int, word wordGenerator) string {
	if wordCount <= 0 {
		return ""
	}

	wordSeparator := ' '
	sentence := bytes.Buffer{}
	sentence.Grow(wordCount * bytesPerWordEstimation)

	for i := 0; i < wordCount; i++ {
		word := word(r)
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

func paragraphGen(r *rand.Rand, opts paragrapOptions, sentecer sentenceGenerator) string {
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
			paragraphs.WriteString(sentecer(r, opts.wordCount))
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

// Question will return a random question
func Question() string {
	return question(globalFaker.Rand)
}

// Question will return a random question
func (f *Faker) Question() string {
	return question(f.Rand)
}

func question(r *rand.Rand) string {
	return strings.Replace(HipsterSentence(Number(3, 10)), ".", "?", 1)
}

// Quote will return a random quote from a random person
func Quote() string { return quote(globalFaker.Rand) }

// Quote will return a random quote from a random person
func (f *Faker) Quote() string { return quote(f.Rand) }

func quote(r *rand.Rand) string {
	return `"` + HipsterSentence(number(r, 3, 10)) + `" - ` + firstName(r) + " " + lastName(r)
}

// Phrase will return a random dictionary phrase
func Phrase() string { return phrase(globalFaker.Rand) }

// Phrase will return a random dictionary phrase
func (f *Faker) Phrase() string { return phrase(f.Rand) }

func phrase(r *rand.Rand) string { return getRandValue(r, []string{"word", "phrase"}) }

func addWordLookup() {
	AddFuncLookup("noun", Info{
		Display:     "Noun",
		Category:    "word",
		Description: "Random noun",
		Example:     "foot",
		Output:      "string",
		Call: func(r *rand.Rand, m *map[string][]string, info *Info) (interface{}, error) {
			return noun(r), nil
		},
	})

	AddFuncLookup("verb", Info{
		Display:     "Verb",
		Category:    "word",
		Description: "Random verb",
		Example:     "release",
		Output:      "string",
		Call: func(r *rand.Rand, m *map[string][]string, info *Info) (interface{}, error) {
			return verb(r), nil
		},
	})

	AddFuncLookup("adverb", Info{
		Display:     "Adverb",
		Category:    "word",
		Description: "Random adverb",
		Example:     "smoothly",
		Output:      "string",
		Call: func(r *rand.Rand, m *map[string][]string, info *Info) (interface{}, error) {
			return adverb(r), nil
		},
	})

	AddFuncLookup("preposition", Info{
		Display:     "Preposition",
		Category:    "word",
		Description: "Random preposition",
		Example:     "down",
		Output:      "string",
		Call: func(r *rand.Rand, m *map[string][]string, info *Info) (interface{}, error) {
			return preposition(r), nil
		},
	})

	AddFuncLookup("adjective", Info{
		Display:     "Adjective",
		Category:    "word",
		Description: "Random adjective",
		Example:     "genuine",
		Output:      "string",
		Call: func(r *rand.Rand, m *map[string][]string, info *Info) (interface{}, error) {
			return adjective(r), nil
		},
	})

	AddFuncLookup("word", Info{
		Display:     "Word",
		Category:    "word",
		Description: "Random word",
		Example:     "man",
		Output:      "string",
		Call: func(r *rand.Rand, m *map[string][]string, info *Info) (interface{}, error) {
			return word(r), nil
		},
	})

	AddFuncLookup("sentence", Info{
		Display:     "Sentence",
		Category:    "word",
		Description: "Random sentence",
		Example:     "Interpret context record river mind.",
		Output:      "string",
		Params: []Param{
			{Field: "wordcount", Display: "Word Count", Type: "int", Default: "5", Description: "Number of words in a sentence"},
		},
		Call: func(r *rand.Rand, m *map[string][]string, info *Info) (interface{}, error) {
			wordCount, err := info.GetInt(m, "wordcount")
			if err != nil {
				return nil, err
			}
			if wordCount <= 0 || wordCount > 50 {
				return nil, errors.New("Invalid word count, must be greater than 0, less than 50")
			}

			return sentence(r, wordCount), nil
		},
	})

	AddFuncLookup("paragraph", Info{
		Display:     "Paragraph",
		Category:    "word",
		Description: "Random paragraph",
		Example:     "Interpret context record river mind press self should compare property outcome divide. Combine approach sustain consult discover explanation direct address church husband seek army. Begin own act welfare replace press suspect stay link place manchester specialist. Arrive price satisfy sign force application hair train provide basis right pay. Close mark teacher strengthen information attempt head touch aim iron tv take. Handle wait begin look speech trust cancer visit capacity disease chancellor clean. Race aim function gain couple push faith enjoy admit ring attitude develop. Edge game prevent cast mill favour father star live search aim guess. West heart item adopt compete equipment miss output report communicate model cabinet. Seek worker variety step argue air improve give succeed relief artist suffer. Hide finish insist knowledge thatcher make research chance structure proportion husband implement. Town crown restaurant cost material compete lady climb football region discussion order. Place lee market ice like display mind stress compete weather station raise. Democracy college major recall struggle use cut intention accept period generation strike. Benefit defend recommend conclude justify result depend succeed address owner fill interpret.",
		Output:      "string",
		Params: []Param{
			{Field: "paragraphcount", Display: "Paragraph Count", Type: "int", Default: "2", Description: "Number of paragraphs"},
			{Field: "sentencecount", Display: "Sentence Count", Type: "int", Default: "2", Description: "Number of sentences in a paragraph"},
			{Field: "wordcount", Display: "Word Count", Type: "int", Default: "5", Description: "Number of words in a sentence"},
			{Field: "paragraphseparator", Display: "Paragraph Separator", Type: "string", Default: "<br />", Description: "String value to add between paragraphs"},
		},
		Call: func(r *rand.Rand, m *map[string][]string, info *Info) (interface{}, error) {
			paragraphCount, err := info.GetInt(m, "paragraphcount")
			if err != nil {
				return nil, err
			}
			if paragraphCount <= 0 || paragraphCount > 20 {
				return nil, errors.New("Invalid paragraph count, must be greater than 0, less than 20")
			}

			sentenceCount, err := info.GetInt(m, "sentencecount")
			if err != nil {
				return nil, err
			}
			if sentenceCount <= 0 || sentenceCount > 20 {
				return nil, errors.New("Invalid sentence count, must be greater than 0, less than 20")
			}

			wordCount, err := info.GetInt(m, "wordcount")
			if err != nil {
				return nil, err
			}
			if wordCount <= 0 || wordCount > 50 {
				return nil, errors.New("Invalid word count, must be greater than 0, less than 50")
			}

			paragraphSeparator, err := info.GetString(m, "paragraphseparator")
			if err != nil {
				return nil, err
			}

			return paragraph(r, paragraphCount, sentenceCount, wordCount, paragraphSeparator), nil
		},
	})

	AddFuncLookup("question", Info{
		Display:     "Question",
		Category:    "word",
		Description: "Random question",
		Example:     "Roof chia echo?",
		Output:      "string",
		Call: func(r *rand.Rand, m *map[string][]string, info *Info) (interface{}, error) {
			return question(r), nil
		},
	})

	AddFuncLookup("quote", Info{
		Display:     "Qoute",
		Category:    "word",
		Description: "Random quote",
		Example:     `"Roof chia echo." - Lura Lockman`,
		Output:      "string",
		Call: func(r *rand.Rand, m *map[string][]string, info *Info) (interface{}, error) {
			return quote(r), nil
		},
	})

	AddFuncLookup("phrase", Info{
		Display:     "Phrase",
		Category:    "word",
		Description: "Random phrase",
		Example:     "time will tell",
		Output:      "string",
		Call: func(r *rand.Rand, m *map[string][]string, info *Info) (interface{}, error) {
			return phrase(r), nil
		},
	})
}
