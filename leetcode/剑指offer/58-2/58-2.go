package _8_2

func reverseLeftWords(s string, n int) string {
	left := s[0:n]
	right := s[n:]
	return right + left
}
