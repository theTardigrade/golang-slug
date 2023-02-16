package slug

type Options struct {
	MaxLen         int
	Replacement    string
	RunesToRemove  []rune
	WholeWords     bool
	Unique         bool
	UniqueAttempts int
}

const (
	optionsDefaultMaxLen         = 128
	optionsDefaultReplacement    = "-"
	optionsDefaultWholeWords     = true
	optionsDefaultUnique         = false
	optionsDefaultUniqueAttempts = 1 << 16
)

func optionsDefaultRunesToRemove() []rune {
	return []rune{
		'\'',
	}
}

func NewOptionsWithDefaults() *Options {
	return &Options{
		MaxLen:         optionsDefaultMaxLen,
		Replacement:    optionsDefaultReplacement,
		RunesToRemove:  optionsDefaultRunesToRemove(),
		WholeWords:     optionsDefaultWholeWords,
		Unique:         optionsDefaultUnique,
		UniqueAttempts: optionsDefaultUniqueAttempts,
	}
}

func initOptions(options *Options) *Options {
	if options == nil {
		options = NewOptionsWithDefaults()
	}

	if replacement := options.Replacement; replacement != optionsDefaultReplacement {
		if replacement == "" || formatWordAcceptedRuneRegexp.MatchString(replacement) {
			replacement = optionsDefaultReplacement
			options.Replacement = replacement
		}
	}

	return options
}
