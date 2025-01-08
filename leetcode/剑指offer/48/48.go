package _8

func lengthOfLongestSubstring(s string) int {
	//双指针的滑动窗口加map
	if len(s) == 0 {
		return 0
	}
	res := 0
	m := make(map[uint8]int)
	for j, i := 0, 0; i < len(s); {
		if m[s[i]] == 0 { //无重复就往前移动一位
			m[s[i]]++
			i++
			if res < i-j {
				res = i - j
			}
		} else { //重复就判断是否是最大子字符串
			//哈希表j位的字母数量-1然后j++
			m[s[j]]--
			j++
		}
	}
	return res
}
