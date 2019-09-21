# slug

## Support

Please consider donating at PayPal: [https://www.paypal.me/jismithpp](https://www.paypal.me/jismithpp)

## Example

```golang
s := slug.GetWithOptions("this is a test", &slug.Options{
	Replacement: "_",
	MaxLen:       12,
})

fmt.Println(s) // prints "this_is_a_te"
```