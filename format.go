package slug

import (
	"regexp"
	"strings"
	"unicode"

	"github.com/rainycape/unidecode"
)

var (
	formatWordAcceptedRuneRegexp = regexp.MustCompile(`[a-z0-9]+`)
)

func formatText(text string, options *Options) string {
	text = strings.TrimSpace(text)
	text = unidecode.Unidecode(text)

	words := formatWordsFromText(text, options)

	return formatBuildFromWords(words, options)
}

func formatWordsFromText(text string, options *Options) (words [][]rune) {
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
			if b := []byte(string(r)); !formatWordAcceptedRuneRegexp.Match(b) {
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

	return
}

func formatBuildFromWords(words [][]rune, options *Options) string {
	var builder strings.Builder

	if options.MaxLen > 0 {
		if options.WholeWords {
			formatBuildFromWordsHandlerWithMaxLenAndWithWholeWords(words, options, &builder)
		} else {
			formatBuildFromWordsHandlerWithMaxLenAndWithoutWholeWords(words, options, &builder)
		}
	} else {
		formatBuildFromWordsHandlerWithoutMaxLen(words, options, &builder)
	}

	return builder.String()
}

func formatBuildFromWordsHandlerWithMaxLenAndWithWholeWords(
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

func formatBuildFromWordsHandlerWithMaxLenAndWithoutWholeWords(
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

func formatBuildFromWordsHandlerWithoutMaxLen(
	words [][]rune, options *Options, builder *strings.Builder,
) {
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
