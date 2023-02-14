# golang-slug

## Example

```golang
package main

import (
	"fmt"

	slug "github.com/theTardigrade/golang-slug"
)

func main() {
	result, err := slug.GetWithOptions("!=this is the text's slug=!", &slug.Options{
		WholeWords:    false,
		MaxLen:        20,
		Replacement:   "_",
		RunesToRemove: slug.DefaultOptions.RunesToRemove,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(result) // prints "this_is_the_texts_sl"
}
```

## Support

If you use this package, or find any value in it, please consider donating at [Ko-fi](https://ko-fi.com/thetardigrade).