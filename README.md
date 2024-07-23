# gotools

This is useful collection of Go functions.

## SmartPlural

`func SmartPlural(number int, singularNoun, pluralNoun string) string`

Enables you to easily deal with the grammatical exception that 1 presents, e.g. "1 card" but "0 cards" and "2 cards".

```go
for i := 1; i <= 4; i++ {
	fmt.Printf("Status: %s\n", SmartPlural(i, "card", ""))
	fmt.Printf("Status: %s\n", SmartPlural(i, "bus", "buses"))
}
```