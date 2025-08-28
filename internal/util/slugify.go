package util

import (
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func removeAccents(value string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)

	result, _, _ := transform.String(t, value)

	return result
}

func Slugify(input string) string {
	unaccented := removeAccents(input)

	return strings.ToLower(strings.Join(strings.Split(unaccented, " "), "-"))
}
