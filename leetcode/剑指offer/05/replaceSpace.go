package _5

import "strings"

func replaceSpace(s string) string {
	return strings.Replace(s, " ", "%20s", -1)
}
