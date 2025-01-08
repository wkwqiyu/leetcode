package _0

func firstUniqChar(s string) byte {
	m := make(map[int32]int)
	for _, value := range s {
		m[value]++
	}
	for k := range m {
		if m[k] == 1 {
			return byte(k)
		}
	}
	return 0
}
