package slug

type Options struct {
	MaxLen        int
	Replacement   string
	RunesToRemove []rune
	WholeWords    bool
}

const (
	defaultOptionsMaxLen      = 128
	defaultOptionsReplacement = "-"
	defaultOptionsWholeWords  = true
)

func defaultOptionsRunesToRemove() []rune {
	return []rune{
		'\'',
	}
}

var (
	DefaultOptions = Options{
		MaxLen:        defaultOptionsMaxLen,
		Replacement:   defaultOptionsReplacement,
		RunesToRemove: defaultOptionsRunesToRemove(),
		WholeWords:    defaultOptionsWholeWords,
	}
)
