package slug

type Options struct {
	Replacement string
	MaxLen      int
}

const (
	defaultOptionsReplacement = "-"
	defaultOptionsMaxLen      = 128
)

var (
	DefaultOptions = Options{
		Replacement: defaultOptionsReplacement,
		MaxLen:      defaultOptionsMaxLen,
	}
)
