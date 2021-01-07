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
	s := slug.GetWithOptions("!=this is a test=!", &slug.Options{
		MaxLen:        12,
		Replacement:   "_",
		RunesToRemove: []rune{'=', '!'},
	})

	fmt.Println(s) // prints "this_is_a_te"
}
```