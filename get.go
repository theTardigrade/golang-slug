package slug

func Get(text string) (slug string, err error) {
	return GetWithOptions(text, NewOptionsWithDefaults())
}

func GetWithOptions(text string, options *Options) (slug string, err error) {
	options = initOptions(options)

	slug = formatText(text, options)

	if options.Unique {
		slug, err = uniqueFormatText(slug, options)
		if err != nil {
			return
		}
	}

	return
}
