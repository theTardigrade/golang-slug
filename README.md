# slug

## Support

Please consider donating at PayPal: [https://www.paypal.me/jismithpp](https://www.paypal.me/jismithpp)

## Example

```golang
package main

import (
	"fmt"

	slug "github.com/theTardigrade/golang-slug"
)

func main() {
	s := slug.GetWithOptions("!=this is the text's slug=!", &slug.Options{
		WholeWords:    false,
		MaxLen:        20,
		Replacement:   "_",
		RunesToRemove: slug.DefaultOptions.RunesToRemove,
	})

	fmt.Println(s) // prints "this_is_the_texts_sl"
}
```