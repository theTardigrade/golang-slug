# golang-slug

This Go package makes it easy to generate [slugs](https://en.wikipedia.org/wiki/Clean_URL#Slug) from human-readable text.

[![Go Reference](https://pkg.go.dev/badge/github.com/theTardigrade/golang-slug.svg)](https://pkg.go.dev/github.com/theTardigrade/golang-slug) [![Go Report Card](https://goreportcard.com/badge/github.com/theTardigrade/golang-slug)](https://goreportcard.com/report/github.com/theTardigrade/golang-slug)

## Example

```golang
package main

import (
	"fmt"

	slug "github.com/theTardigrade/golang-slug"
)

func main() {
	defaultOptions := slug.NewOptionsWithDefaults()

	result, err := slug.GetWithOptions("!=this is the text's slug=!", &slug.Options{
		WholeWords:    false,
		MaxLen:        20,
		Replacement:   "_",
		RunesToRemove: defaultOptions.RunesToRemove,
	})
	if err != nil {
		panic(err)
	}

	// prints "this_is_the_texts_sl"
	fmt.Println(result)

	result, err = slug.GetRandomWithOptions(&slug.Options{
		MaxLen: 20,
	})
	if err != nil {
		panic(err)
	}

	// prints slug made up of twenty randomly selected
	// uppercase letters, lowercase letters and digits
	fmt.Println(result)
}
```

## Support

If you use this package, or find any value in it, please consider donating:

[![ko-fi](https://ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/S6S2EIRL0)