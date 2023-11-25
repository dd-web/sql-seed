package types

import (
	"math/rand"
)

// word generation settings - words are counted by letters.
var lorem_word_min_length = 3
var lorem_word_max_length = 10

// sentence generation settings - sentences are counted by words.
var lorem_sentence_min_length = 8
var lorem_sentence_max_length = 18

// paragraph generation settings - paragraphs are counted by sentences.
var lorem_paragraph_min_length = 3
var lorem_paragraph_max_length = 7

// the minimum and maximum number of paragraphs to generate
var lorem_min_paragraphs = 1
var lorem_max_paragraphs = 6

// offset bitmask for converting upper case to lower case
var upper_to_lower_offset_mask rune = 0x20

// the lower case starting and end bound masks
var lower_alpha_start rune = 0x61
var lower_alpha_end rune = 0x7a

// special character map
var special_chars = map[string]rune{
	"SPACE":         0x20,
	"FULL-QUOTE":    0x22,
	"APOSTROPHE":    0x27,
	"HYPHEN":        0x2d,
	"PERIOD":        0x2e,
	"FORWARD-SLASH": 0x2f,

	"LESS-THAN":    0x3c,
	"EQUALS":       0x3d,
	"GREATER-THAN": 0x3e,
	"QUESTION":     0x3f,

	"LEFT-BRACKET":   0x5b,
	"BACKWARD-SLASH": 0x5c,
	"RIGHT-BRACKET":  0x5d,
	"CARET":          0x5e,
	"UNDERSCORE":     0x5f,
	"TILDA":          0x60,
	"CRLF":           0x0a,
}

/*************************************************************************************************************/
/*************************************************************************************************************/

// below defines lorem types and configuration types to make configuration options easier to use
// defaults will be used if no configuration is given (set via config fns)

// container struct for our generation
type Lorem struct {
	// output
	Output string

	// configuration
	Cfg *LoremConfig
}

// lorem config struct
type LoremConfig struct {
	/* word length settings */
	MinWordLength int
	MaxWordLength int

	/* sentence length settings */
	MinSentenceLength int
	MaxSentenceLength int

	/* paragraph length settings */
	MinParagraphLength int
	MaxParagraphLength int
	MinParagraphs      int
	MaxParagraphs      int

	/* sentence settings */
	CapitalizeFirst bool
	Punctuation     bool

	/* defines a set of characters available to be selected as a termination character for a sentence */
	PunctuationChars []string

	/* gives each punctuation character a weight, which is used to determine the likelihood of it being selected */
	/* weights don't need to add up to 100, it's just easier to set a default this way. */
	PunctuationWeights map[string]int
}

// default lorem configuration
func defaultLoremConfig() *LoremConfig {
	return &LoremConfig{
		MinWordLength:      lorem_word_min_length,
		MaxWordLength:      lorem_word_max_length,
		MinSentenceLength:  lorem_sentence_min_length,
		MaxSentenceLength:  lorem_sentence_max_length,
		MinParagraphLength: lorem_paragraph_min_length,
		MaxParagraphLength: lorem_paragraph_max_length,
		MinParagraphs:      lorem_min_paragraphs,
		MaxParagraphs:      lorem_max_paragraphs,
		CapitalizeFirst:    true,
		Punctuation:        true,
		PunctuationChars:   []string{".", "!", "?"},
		PunctuationWeights: map[string]int{".": 85, "!": 10, "?": 5},
	}
}

// defines a configuration function to modify the configuration - must be called on lorem instantiation
type LoremConfigFunc func(*LoremConfig) *LoremConfig

/*******************************/
/*** CONFIGURATION FUNCTIONS ***/
/*******************************/

// min word length
func LoremMinWordLength(i int) LoremConfigFunc {
	return func(c *LoremConfig) *LoremConfig {
		c.MinWordLength = i
		return c
	}
}

// max word length
func LoremMaxWordLength(i int) LoremConfigFunc {
	return func(c *LoremConfig) *LoremConfig {
		c.MaxWordLength = i
		return c
	}
}

// min sentence length
func LoremMinSentenceLength(i int) LoremConfigFunc {
	return func(c *LoremConfig) *LoremConfig {
		c.MinSentenceLength = i
		return c
	}
}

// max sentence length
func LoremMaxSentenceLength(i int) LoremConfigFunc {
	return func(c *LoremConfig) *LoremConfig {
		c.MaxSentenceLength = i
		return c
	}
}

// min paragraph count
func LoremMinParagraphCount(i int) LoremConfigFunc {
	return func(c *LoremConfig) *LoremConfig {
		c.MinParagraphs = i
		return c
	}
}

// max paragraph count
func LoremMaxParagraphCount(i int) LoremConfigFunc {
	return func(c *LoremConfig) *LoremConfig {
		c.MaxParagraphs = i
		return c
	}
}

// capitalize first letter of sentence
func LoremCapitalizeFirst(b bool) LoremConfigFunc {
	return func(c *LoremConfig) *LoremConfig {
		c.CapitalizeFirst = b
		return c
	}
}

// add punctuation to the end of sentences
func LoremPunctuation(b bool) LoremConfigFunc {
	return func(c *LoremConfig) *LoremConfig {
		c.Punctuation = b
		return c
	}
}

// add a character to the list of punctuation characters
func LoremAddPunctuationChar(s string) LoremConfigFunc {
	return func(c *LoremConfig) *LoremConfig {
		c.PunctuationChars = append(c.PunctuationChars, s)
		return c
	}
}

// remove a character from the list of punctuation characters
func LoremRemovePunctuationChar(s string) LoremConfigFunc {
	return func(c *LoremConfig) *LoremConfig {
		for i, v := range c.PunctuationChars {
			if v == s {
				c.PunctuationChars = append(c.PunctuationChars[:i], c.PunctuationChars[i+1:]...)
				break
			}
		}
		return c
	}
}

// replace the list of punctuation characters with the given list
func LoremSetPunctuationChars(s []string) LoremConfigFunc {
	return func(c *LoremConfig) *LoremConfig {
		c.PunctuationChars = s
		return c
	}
}

// add a weight to the given punctuation character - this will replace any existing weight
func LoremAddPunctuationWeight(s string, i int) LoremConfigFunc {
	return func(c *LoremConfig) *LoremConfig {
		c.PunctuationWeights[s] = i
		return c
	}
}

/* asserts */
func (l *Lorem) check() {
	// no checks yet
	// @TODO runtime checks
}

/*********************************/
/** END CONFIGURATION FUNCTIONS **/
/*********************************/

// Creates a new lorem with the given configuration
// defaults will be used if no configuration is given (set via config fns)
func NewLorem(cfg ...LoremConfigFunc) *Lorem {
	config := defaultLoremConfig()
	for _, fn := range cfg {
		config = fn(config)
	}
	lorem := &Lorem{
		Cfg: config,
	}
	lorem.check()
	return lorem
}

/** GENERATION FUNCTIONS **/

// Main entry point for content generation. call after instantiation.
func (l *Lorem) Generate() string {
	paragraphs := ""
	paragraphCount := RandomBetween(l.Cfg.MinParagraphs, l.Cfg.MaxParagraphs)
	for i := 0; i < paragraphCount; i++ {
		paragraphs += l.paragraph()
	}

	l.Output = paragraphs
	return l.Output
}

// generates a random word defined by the configuration
func (l *Lorem) word() string {
	word := ""
	wordLen := RandomBetween(l.Cfg.MinWordLength, l.Cfg.MaxWordLength)
	for i := 0; i < wordLen; i++ {
		word += string(RandomBetween[rune](lower_alpha_start, lower_alpha_end))
	}
	return word
}

// generates a random sentence defined by the configuration
func (l *Lorem) sentence() string {
	sentence := ""
	wordCount := RandomBetween[int](l.Cfg.MinSentenceLength, l.Cfg.MaxSentenceLength)
	for i := 0; i < wordCount; i++ {
		if i > 0 {
			sentence += string(special_chars["SPACE"])
		}
		sentence += l.word()
	}
	if l.Cfg.CapitalizeFirst && len(sentence) > 0 {
		fl := rune(sentence[0])
		sentence = string(fl-upper_to_lower_offset_mask) + sentence[1:]
	}
	if l.Cfg.Punctuation && len(sentence) > 0 {
		sentence += l.punctuation()
	}
	return sentence
}

// generates a random paragraph defined by the configuration
func (l *Lorem) paragraph() string {
	paragraph := string(special_chars["SPACE"]) + string(special_chars["SPACE"])
	sentenceCount := RandomBetween(l.Cfg.MinParagraphLength, l.Cfg.MaxParagraphLength)
	for i := 0; i < sentenceCount; i++ {
		paragraph = paragraph + string(special_chars["SPACE"]) + l.sentence()
	}
	return paragraph + string(special_chars["CRLF"]) + string(special_chars["CRLF"])
}

// grabs a weighted punctuation character from the configuration
func (l *Lorem) punctuation() string {
	var cumulativeWeights []int
	var characters []string
	cumulative := 0

	for char, weight := range l.Cfg.PunctuationWeights {
		cumulative += weight
		cumulativeWeights = append(cumulativeWeights, cumulative)
		characters = append(characters, char)
	}

	r := rand.Intn(cumulativeWeights[len(cumulativeWeights)-1])

	for i, weight := range cumulativeWeights {
		if r < weight {
			return characters[i]
		}
	}

	return string(special_chars["PERIOD"])
}

// wraps the content with the given tag as an html tag
func WrapHTMLTag(tag, content string) string {
	return string(special_chars["LESS-THAN"]) + tag + string(special_chars["GREATER-THAN"]) + content + string(special_chars["LESS-THAN"]) + string(special_chars["FORWARD-SLASH"]) + tag + string(special_chars["GREATER-THAN"])
}
