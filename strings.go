package gotools

import "fmt"

func SmartPlural(number int, singularNoun, pluralNoun string) string {
	if pluralNoun == "" {
		pluralNoun = singularNoun + "s"
	}
	
	var result string
	result += fmt.Sprintf("%d ", number)
	if number == 1 {
		result += singularNoun
	} else {
		result += pluralNoun
	}

	return result
}