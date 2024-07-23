# gotools

This is my collection of useful Go functions.

## SmartPlural

`func SmartPlural(number int, singularNoun, pluralNoun string) string`

Enables you to easily deal with the grammatical exception that 1 presents, e.g. "0 cards", "2 cards", "999 cards" but "1 card".

```go
for i := 1; i <= 4; i++ {
	fmt.Printf("Status: %s\n", SmartPlural(i, "card", ""))
	fmt.Printf("Status: %s\n", SmartPlural(i, "bus", "buses"))
}
```