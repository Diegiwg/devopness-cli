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

func RenderArgs(arguments map[string]string, separator string) string {
	var result string

	for name, _type := range arguments {
		result += name + " " + _type + separator
	}

	return strings.TrimSuffix(result, separator)
}