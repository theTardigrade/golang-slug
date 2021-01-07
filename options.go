package slug

type Options struct {
	MaxLen        int
	Replacement   string
	RunesToRemove []rune
}

const (
	defaultOptionsMaxLen      = 128
	defaultOptionsReplacement = "-"
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
	}
)
