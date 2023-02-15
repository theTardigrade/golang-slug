package slug

import (
	"regexp"
	"strings"
	"unicode"

	"github.com/rainycape/unidecode"
)

var (
	regexpWordAcceptedRune = regexp.MustCompile(`[a-z0-9]+`)
)

func Get(text string) (slug string, err error) {
	return GetWithOptions(text, NewOptionsWithDefaults())
}

func GetWithOptions(text string, options *Options) (slug string, err error) {
	if options == nil {
		options = NewOptionsWithDefaults()
	}

	if replacement := options.Replacement; replacement != optionsDefaultReplacement {
		if replacement == "" || regexpWordAcceptedRune.MatchString(replacement) {
			replacement = optionsDefaultReplacement
			options.Replacement = replacement
		}
	}

	slug = format(text, options)

	if options.Unique {
		slug, err = uniqueFormat(slug, options)
	}

	return
}

func format(text string, options *Options) string {
	text = strings.TrimSpace(text)
	text = unidecode.Unidecode(text)

	var words [][]rune
	wordBreak := true
	currentWord := []rune{}

	for _, r := range text {
		r = unicode.ToLower(r)
		shouldAdd := true

		for _, r2 := range options.RunesToRemove {
			if r2 == r {
				shouldAdd = false
				break
			}
		}

		if shouldAdd {
			if b := []byte(string(r)); !regexpWordAcceptedRune.Match(b) {
				shouldAdd = false
				wordBreak = true
			} else if wordBreak {
				if len(currentWord) > 0 {
					words = append(words, currentWord)
					currentWord = []rune{}
				}
				wordBreak = false
			}
		}

		if shouldAdd {
			currentWord = append(currentWord, r)
		}
	}

	if len(currentWord) > 0 {
		words = append(words, currentWord)
	}

	return buildFromWords(words, options)
}

func buildFromWords(words [][]rune, options *Options) string {
	var builder strings.Builder

	if options.MaxLen > 0 {
		if options.WholeWords {
			buildFromWordsHandlerWithMaxLenAndWithWholeWords(words, options, &builder)
		} else {
			buildFromWordsHandlerWithMaxLenAndWithoutWholeWords(words, options, &builder)
		}
	} else {
		buildFromWordsHandlerWithoutMaxLen(words, options, &builder)
	}

	return builder.String()
}

func buildFromWordsHandlerWithMaxLenAndWithWholeWords(
	words [][]rune, options *Options, builder *strings.Builder,
) {
	maxLen := options.MaxLen
	replacement := options.Replacement
	replacementLen := len(replacement)
	l := 0

	for i, w := range words {
		if i > 0 {
			if l+replacementLen+len(w) > maxLen {
				break
			}

			builder.WriteString(replacement)
			l += replacementLen
		} else {
			if l+len(w) > maxLen {
				break
			}
		}

		for _, r := range w {
			builder.WriteRune(r)
			l++
		}
	}
}

func buildFromWordsHandlerWithMaxLenAndWithoutWholeWords(
	words [][]rune, options *Options, builder *strings.Builder,
) {
	maxLen := options.MaxLen
	replacement := options.Replacement
	replacementLen := len(replacement)
	l := 0

	for i, w := range words {
		if i > 0 {
			if l+replacementLen+1 > maxLen {
				break
			}

			builder.WriteString(replacement)
			l += replacementLen
		} else {
			if l+1 > maxLen {
				break
			}
		}

		if l+len(w) <= maxLen {
			for _, r := range w {
				builder.WriteRune(r)
				l++
			}
		} else {
			deltaLen := maxLen - l

			for i, r := range w {
				if i == deltaLen {
					break
				}

				builder.WriteRune(r)
				l++
			}
		}
	}
}

func buildFromWordsHandlerWithoutMaxLen(words [][]rune, options *Options, builder *strings.Builder) {
	replacement := options.Replacement

	for i, w := range words {
		if i > 0 {
			builder.WriteString(replacement)
		}
		for _, r := range w {
			builder.WriteRune(r)
		}
	}
}
