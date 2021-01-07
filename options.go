package slug

type Options struct {
	Replacement   string
	MaxLen        int
	RunesToRemove []rune
}

const (
	defaultOptionsReplacement = "-"
	defaultOptionsMaxLen      = 128
)

func defaultOptionsRunesToRemove() []rune {
	return []rune{
		'\'',
	}
}

var (
	DefaultOptions = Options{
		Replacement:   defaultOptionsReplacement,
		MaxLen:        defaultOptionsMaxLen,
		RunesToRemove: defaultOptionsRunesToRemove(),
	}
)
