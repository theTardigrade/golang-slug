package slug

import (
	"strconv"
	"sync"
)

var (
	uniqueCache      = make(map[string]struct{})
	uniqueCacheMutex sync.Mutex
)

func uniqueFormatText(slug string, options *Options) (uniqueSlug string, err error) {
	defer uniqueCacheMutex.Unlock()
	uniqueCacheMutex.Lock()

	potentialSlug := slug
	var suffixValue uint64 = 1

	for i := options.UniqueAttempts; i > 0; i-- {
		if _, found := uniqueCache[potentialSlug]; !found {
			uniqueCache[potentialSlug] = struct{}{}
			uniqueSlug = potentialSlug

			return
		}

		suffixValue++
		suffix := strconv.FormatUint(suffixValue, 16)

		slugLen := len(slug)
		replacementLen := len(options.Replacement)
		suffixLen := len(suffix)

		for {
			totalLen := slugLen + suffixLen

			if slug[slugLen-replacementLen:] != options.Replacement {
				totalLen += replacementLen
			}

			if totalLen <= options.MaxLen {
				break
			}

			if slugLen--; slugLen == 0 {
				err = ErrUniqueLength

				return
			}
		}

		potentialSlug = formatText(slug[:slugLen]+options.Replacement+suffix, options)
	}

	err = ErrUniqueAttempts

	return
}
