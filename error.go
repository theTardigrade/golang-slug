package slug

import "errors"

var (
	ErrUniqueAttempts = errors.New("too many unique slugs attempted with no success")
	ErrUniqueLength   = errors.New("not enough length to make all unique attempts")
)
