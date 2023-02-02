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
	defaultOptionsMaxLen         = 128
	defaultOptionsReplacement    = "-"
	defaultOptionsWholeWords     = true
	defaultOptionsUnique         = false
	defaultOptionsUniqueAttempts = 1 << 16
)

func defaultOptionsRunesToRemove() []rune {
	return []rune{
		'\'',
	}
}

var (
	DefaultOptions = Options{
		MaxLen:         defaultOptionsMaxLen,
		Replacement:    defaultOptionsReplacement,
		RunesToRemove:  defaultOptionsRunesToRemove(),
		WholeWords:     defaultOptionsWholeWords,
		Unique:         defaultOptionsUnique,
		UniqueAttempts: defaultOptionsUniqueAttempts,
	}
)
