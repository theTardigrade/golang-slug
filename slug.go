package slug

import (
	"regexp"
	"strings"

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
		replacement = format(replacement)

		if replacement == "" || !regexpReplace.MatchString(replacement) {
			replacement = DefaultOptions.Replacement
		}
	}

	slug = format(text)

	slug = regexpReplace.ReplaceAllString(slug, replacement)
	slug = strings.TrimSuffix(slug, replacement)

	if maxLen := options.MaxLen; len(slug) > maxLen {
		slug = slug[:maxLen]
	}

	return
}

func format(text string) string {
	return strings.ToLower(
		unidecode.Unidecode(
			strings.TrimSpace(text),
		),
	)
}
