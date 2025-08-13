package util

import "strings"

func Slugify(input string) string {
	return strings.ToLower(strings.Join(strings.Split(input, " "), "-"))
}
