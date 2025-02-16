package generator

import "strings"

func Capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}

func CamelCase(s string) string {
	if len(s) == 0 {
		return s
	}

	s = strings.ReplaceAll(s, "-", " ")
	s = strings.ReplaceAll(s, "_", " ")

	words := strings.Split(s, " ")

	for i, word := range words {
		words[i] = Capitalize(word)
	}

	return strings.Join(words, "")
}
