package _8

/*func permutation(s string) []string {
	b := []byte(s)
	var sarray []string
	dfs(0, &b, &sarray)
	return sarray
}

func dfs(x int, b *[]byte, sarray *[]string) { //x为要固定的索引
	if x == len(*b)-1 {
		*sarray = append(*sarray, string(*b))
		return
	}
	m := make(map[byte]bool)
	for i := x; i < len(*b); i++ {
		if m[(*b)[i]] == true {
			continue
		}
		m[(*b)[i]] = true
		swap(i, x, b)
		dfs(x+1, b, sarray) //固定下一位
		swap(i, x, b)       //回溯
	}
	return
}

func swap(a, b int, by *[]byte) {
	tmp := (*by)[a]
	(*by)[a] = (*by)[b]
	(*by)[b] = tmp
}*/

// 另一个版本
func permutation(s string) []string {
	res := []string{}
	var cs string
	hmap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		hmap[s[i]]++
	}
	var dfs func(start int)

	dfs = func(start int) {
		if start == len(s) { //数组长度足够就添加进结果
			res = append(res, cs)
			return
		}
		for a, _ := range hmap {
			if hmap[a] != 0 { //判断该字母是否可用，可用添加
				cs += string(a)
				hmap[a]--
				dfs(start + 1)      //添加下一位
				cs = cs[:len(cs)-1] //回溯
				hmap[a]++           //表示再次可用
			}
		}
	}
	dfs(0)
	return res
}
