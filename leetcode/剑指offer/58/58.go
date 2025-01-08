package _8

import "strings"

func reverseWords(s string) string {
	words := strings.Split(s, " ")
	var res []string
	for i := len(words) - 1; i >= 0; i-- {
		words[i] = strings.Trim(words[i], " ")
		if len(words) > 0 {
			res = append(res, words[i])
		}
	}
	return strings.Join(res, " ")
}
