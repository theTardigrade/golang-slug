package slug

import (
	"regexp"
	"strings"
	"unicode"

	"github.com/rainycape/unidecode"
)

var (
	regexpReplace = regexp.MustCompile(`[^a-z0-9]+`)
)

func Get(text string) (slug string) {
	return GetWithOptions(text, &DefaultOptions)
}

func GetWithOptions(text string, options *Options) (slug string) {
	if options == nil {
		options = &DefaultOptions
	}

	replacement := options.Replacement

	if replacement != defaultOptionsReplacement {
		replacement = format(replacement, options)

		if replacement == "" || !regexpReplace.MatchString(replacement) {
			replacement = DefaultOptions.Replacement
		}
	}

	slug = format(text, options)
	slug = regexpReplace.ReplaceAllString(slug, replacement)

	if maxLen := options.MaxLen; maxLen > 0 && len(slug) > maxLen {
		slug = slug[:maxLen]
	}

	slug = strings.TrimSuffix(slug, replacement)

	return
}

func format(text string, options *Options) string {
	text = strings.TrimSpace(text)

	if len(options.RunesToRemove) > 0 {
		var builder strings.Builder

		for _, r := range text {
			shouldAdd := true

			for _, r2 := range options.RunesToRemove {
				if r2 == r {
					shouldAdd = false
					break
				}
			}

			if shouldAdd {
				builder.WriteRune(
					unicode.ToLower(r),
				)
			}
		}

		text = builder.String()
	} else {
		text = strings.ToLower(text)
	}

	text = strings.ReplaceAll(text, `'`, "")
	text = unidecode.Unidecode(text)

	return text
}
