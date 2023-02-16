package slug

import (
	passwordGenerator "github.com/theTardigrade/golang-passwordGenerator"
)

func GetRandom() (slug string, err error) {
	return GetRandomWithOptions(NewOptionsWithDefaults())
}

func GetRandomWithOptions(options *Options) (slug string, err error) {
	options = initOptions(options)

	generator := passwordGenerator.New(
		passwordGenerator.Options{
			Len:                     options.MaxLen,
			IncludeUpperCaseLetters: true,
			IncludeLowerCaseLetters: true,
			IncludeDigits:           true,
			ExcludeAmbiguousRunes:   true,
		},
	)

	text, err := generator.Generate()
	if err != nil {
		return
	}

	slug = formatText(text, options)

	if options.Unique {
		slug, err = uniqueFormatText(slug, options)
		if err != nil {
			return
		}
	}

	return
}
