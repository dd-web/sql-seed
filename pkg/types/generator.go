package types

import (
	"strconv"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

/***************************/
/* LOREM / TEXT GENERATION */
/***************************/

var (
	lorem_word_min_length      int = 3
	lorem_word_max_length      int = 10
	lorem_sentence_min_length  int = 8
	lorem_sentence_max_length  int = 18
	lorem_paragraph_min_length int = 3
	lorem_paragraph_max_length int = 7
	lorem_min_paragraphs       int = 1
	lorem_max_paragraphs       int = 6

	uppercase_bitmask rune = 0x20
	lower_alpha_start rune = 0x61
	lower_alpha_end   rune = 0x7a

	special_chars = map[string]rune{
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

	safe_special_chars = []rune{
		special_chars["HYPHEN"],
		special_chars["UNDERSCORE"],
		special_chars["PERIOD"],
	}
)

type Lorem struct {
	Output string
	Cfg    *LoremConfig
}

type LoremConfigFunc func(*LoremConfig) *LoremConfig
type LoremConfig struct {
	minWordLength      int
	maxWordLength      int
	minSentenceLength  int
	maxSentenceLength  int
	minParagraphLength int
	maxParagraphLength int
	minParagraphs      int
	maxParagraphs      int

	capitalizeFirst bool
	punctuation     bool

	punctuationChars []string

	punctuationWeights map[string]int
}

func defaultLoremConfig() *LoremConfig {
	return &LoremConfig{
		minWordLength:      lorem_word_min_length,
		maxWordLength:      lorem_word_max_length,
		minSentenceLength:  lorem_sentence_min_length,
		maxSentenceLength:  lorem_sentence_max_length,
		minParagraphLength: lorem_paragraph_min_length,
		maxParagraphLength: lorem_paragraph_max_length,
		minParagraphs:      lorem_min_paragraphs,
		maxParagraphs:      lorem_max_paragraphs,
		capitalizeFirst:    true,
		punctuation:        true,
		punctuationChars:   []string{".", "!", "?"},
		punctuationWeights: map[string]int{".": 20, "!": 1, "?": 1},
	}
}

/***************************/
/* CONFIGURATION FUNCTIONS */
/***************************/

func LoremMinWordLength(i int) LoremConfigFunc {
	return func(c *LoremConfig) *LoremConfig {
		c.minWordLength = i
		return c
	}
}

func LoremMaxWordLength(i int) LoremConfigFunc {
	return func(c *LoremConfig) *LoremConfig {
		c.maxWordLength = i
		return c
	}
}

func LoremMinSentenceLength(i int) LoremConfigFunc {
	return func(c *LoremConfig) *LoremConfig {
		c.minSentenceLength = i
		return c
	}
}

func LoremMaxSentenceLength(i int) LoremConfigFunc {
	return func(c *LoremConfig) *LoremConfig {
		c.maxSentenceLength = i
		return c
	}
}

func LoremMinParagraphCount(i int) LoremConfigFunc {
	return func(c *LoremConfig) *LoremConfig {
		c.minParagraphs = i
		return c
	}
}

func LoremMaxParagraphCount(i int) LoremConfigFunc {
	return func(c *LoremConfig) *LoremConfig {
		c.maxParagraphs = i
		return c
	}
}

func LoremCapitalizeFirst(b bool) LoremConfigFunc {
	return func(c *LoremConfig) *LoremConfig {
		c.capitalizeFirst = b
		return c
	}
}

func LoremPunctuation(b bool) LoremConfigFunc {
	return func(c *LoremConfig) *LoremConfig {
		c.punctuation = b
		return c
	}
}

func LoremAddPunctuationChar(s string) LoremConfigFunc {
	return func(c *LoremConfig) *LoremConfig {
		c.punctuationChars = append(c.punctuationChars, s)
		return c
	}
}

func LoremRemovePunctuationChar(s string) LoremConfigFunc {
	return func(c *LoremConfig) *LoremConfig {
		for i, v := range c.punctuationChars {
			if v == s {
				c.punctuationChars = append(c.punctuationChars[:i], c.punctuationChars[i+1:]...)
				break
			}
		}
		return c
	}
}

func LoremSetPunctuationChars(s []string) LoremConfigFunc {
	return func(c *LoremConfig) *LoremConfig {
		c.punctuationChars = s
		return c
	}
}

func LoremAddPunctuationWeight(s string, i int) LoremConfigFunc {
	return func(c *LoremConfig) *LoremConfig {
		c.punctuationWeights[s] = i
		return c
	}
}

func NewLorem(cfg ...LoremConfigFunc) *Lorem {
	config := defaultLoremConfig()
	for _, fn := range cfg {
		config = fn(config)
	}
	return &Lorem{
		Cfg:    config,
		Output: "",
	}
}

// Generate begins the generation process and returns the generated string as well as setting
// the Output field of the Lorem struct so it can be accessed later if needed
func (l *Lorem) Generate() string {
	paragraphs := ""
	paragraphCount := RandomBetween[int](l.Cfg.minParagraphs, l.Cfg.maxParagraphs)
	for i := 0; i < paragraphCount; i++ {
		paragraphs += l.paragraph()
	}

	l.Output = paragraphs
	return l.Output
}

func (l *Lorem) GenerateParagraph() string {
	paragraph := l.paragraph()
	l.Output = paragraph
	return l.Output
}

func (l *Lorem) GenerateSentence() string {
	sentence := l.sentence()
	l.Output = sentence
	return l.Output
}

func (l *Lorem) GenerateWord() string {
	word := l.word()
	l.Output = word
	return l.Output
}

// generates a random word defined by the configuration
func (l *Lorem) word() string {
	word := ""
	wordLen := RandomBetween[int](l.Cfg.minWordLength, l.Cfg.maxWordLength)
	for i := 0; i < wordLen; i++ {
		word += string(RandomBetween[rune](lower_alpha_start, lower_alpha_end))
	}
	return word
}

// generates a random sentence defined by the configuration
func (l *Lorem) sentence() string {
	sentence := ""
	wordCount := RandomBetween[int](l.Cfg.minSentenceLength, l.Cfg.maxSentenceLength)
	for i := 0; i < wordCount; i++ {
		if i > 0 {
			sentence += string(special_chars["SPACE"])
		}
		sentence += l.word()
	}
	if l.Cfg.capitalizeFirst && len(sentence) > 0 {
		sentence = string(rune(sentence[0])-uppercase_bitmask) + sentence[1:]
	}
	if l.Cfg.punctuation && len(sentence) > 0 {
		sentence += RandomWeightedFromMap[string](l.Cfg.punctuationWeights)
	}
	return sentence
}

// generates a random paragraph defined by the configuration
func (l *Lorem) paragraph() string {
	paragraph := string(special_chars["SPACE"]) + string(special_chars["SPACE"])
	sentenceCount := RandomBetween[int](l.Cfg.minParagraphLength, l.Cfg.maxParagraphLength)
	for i := 0; i < sentenceCount; i++ {
		paragraph = paragraph + string(special_chars["SPACE"]) + l.sentence()
	}
	return paragraph + string(special_chars["CRLF"]) + string(special_chars["CRLF"])
}

/*******************/
/* SLUG GENERATION */
/*******************/

var (
	identity_slug_charset string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-"
	thread_slug_charset   string = "abcdefghijklmnopqrstuvwxyz0123456789-"
	article_slug_charset  string = "abcdefghijklmnopqrstuvwxyz-"

	identity_slug_min_length int = 8
	identity_slug_max_length int = 10
	thread_slug_min_length   int = 12
	thread_slug_max_length   int = 16
	article_slug_min_length  int = 10
	article_slug_max_length  int = 25
)

func slug(charset string, min, max int) string {
	sluglen := RandomBetween[int](min, max)
	slug, _ := gonanoid.Generate(charset, sluglen)
	return slug
}

func NewIdentitySlug() string {
	return slug(identity_slug_charset, identity_slug_min_length, identity_slug_max_length)
}

func NewThreadSlug() string {
	return slug(thread_slug_charset, thread_slug_min_length, thread_slug_max_length)
}

func NewArticleSlug() string {
	return slug(article_slug_charset, article_slug_min_length, article_slug_max_length)
}

/*************************/
/* NAME/EMAIL GENERATION */
/*************************/

var (
	username_min_length int = 6
	username_max_length int = 12

	email_domain_weights = map[string]int{
		"gmail.com":      100,
		"yahoo.com":      75,
		"hotmail.com":    60,
		"outlook.com":    20,
		"protonmail.com": 10,
		"live.com":       50,
		"icloud.com":     35,
		"fastmail.org":   5,
		"mit.edu":        3,
		"nyu.edu":        6,
		"brown.edu":      5,
		"fbi.gov":        4,
		"nsa.gov":        1,
	}

	uword_step_weights = map[int]int{
		1: 90, // add a lettr to the end
		2: 20, // add a letter to the end and capitalize it
		3: 4,  // add a number to the end
		4: 2,  // add a special character to the end
	}
)

func safeRandomSpecialChar() rune {
	return safe_special_chars[RandomBetween[int](0, len(safe_special_chars))]
}

// called for each letter of a username to generate the entire username
func uwordStep(current string) string {
	step := RandomWeightedFromMap[int](uword_step_weights)
	word := current

	switch step {
	case 1:
		word += string(RandomBetween[rune](lower_alpha_start, lower_alpha_end))
	case 2:
		word += string(RandomBetween[rune](lower_alpha_start-uppercase_bitmask, lower_alpha_end-uppercase_bitmask))
	case 3:
		word += strconv.Itoa(RandomBetween[int](0, 9))
	case 4:
		word += string(safeRandomSpecialChar())
	}

	// fmt.Println("step:", step, " og:", current, " new:", word)

	return word
}

func AddDomainSuffix(u string) string {
	return u + "@" + RandomWeightedFromMap[string](email_domain_weights)
}

func NewUsername() string {
	usernameLen := RandomBetween[int](username_min_length, username_max_length)
	username := ""
	for i := 0; i < usernameLen; i++ {
		username = uwordStep(username)
	}
	return username
}

/******************/
/* ENUM RESOLVERS */
/******************/

var (
	enum_account_role_weights = map[Enum]int{
		AccountRoleUser:      90,
		AccountRoleModerator: 5,
		AccountRoleAdmin:     2,
	}

	enum_account_status_weights = map[Enum]int{
		AccountStatusActive:    90,
		AccountStatusInactive:  10,
		AccountStatusSuspended: 10,
		AccountStatusBanned:    5,
	}

	enum_article_status_weights = map[Enum]int{
		ArticleStatusDraft:     20,
		ArticleStatusReview:    5,
		ArticleStatusPublished: 90,
		ArticleStatusArchived:  10,
		ArticleStatusRetracted: 2,
	}

	enum_thread_status_weights = map[Enum]int{
		ThreadStatusOpen:     90,
		ThreadStatusClosed:   5,
		ThreadStatusArchived: 10,
		ThreadStatusRemoved:  2,
	}

	enum_thread_role_weights = map[Enum]int{
		ThreadRoleUser:      90,
		ThreadRoleModerator: 5,
	}

	enum_identity_status_weights = map[Enum]int{
		IdentityStatusActive:    90,
		IdentityStatusInactive:  15,
		IdentityStatusSuspended: 6,
		IdentityStatusBanned:    2,
	}
)

func RandomEnumAccountRole() AccountRole {
	return RandomWeightedFromMap[Enum](enum_account_role_weights).(AccountRole)
}

func RandomEnumAccountStatus() AccountStatus {
	return RandomWeightedFromMap[Enum](enum_account_status_weights).(AccountStatus)
}

func RandomEnumArticleStatus() ArticleStatus {
	return RandomWeightedFromMap[Enum](enum_article_status_weights).(ArticleStatus)
}

func RandomEnumThreadStatus() ThreadStatus {
	return RandomWeightedFromMap[Enum](enum_thread_status_weights).(ThreadStatus)
}

func RandomEnumThreadRole() ThreadRole {
	return RandomWeightedFromMap[Enum](enum_thread_role_weights).(ThreadRole)
}

func RandomEnumIdentityStatus() IdentityStatus {
	return RandomWeightedFromMap[Enum](enum_identity_status_weights).(IdentityStatus)
}

func RandomEnumIdentityStyle() IdentityStyle {
	return IdentityStyleID[RandomBetween[int](1, len(IdentityStyleID))]
}

/*************************************************************************************************************/
/*************************************************************************************************************/

// wraps the content with the given tag as an html tag
func WrapHTMLTag(tag, content string) string {
	return string(special_chars["LESS-THAN"]) + tag + string(special_chars["GREATER-THAN"]) + content + string(special_chars["LESS-THAN"]) + string(special_chars["FORWARD-SLASH"]) + tag + string(special_chars["GREATER-THAN"])
}
