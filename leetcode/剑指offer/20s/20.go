package _0s

/*
本题的思想是把一串字符串分为起始，中间部分和结尾部分
思路为先判断是否为满足起始字符，满足就检查是否满足中间部分
如果满足其中一个中间部分，就检查是否满足该中间部分对应的下一个中间部分或者结尾部分，当字符串遍历完后，检查是否满足结尾部分。
若当中有一个不满足就return false
起始，中间部分和结尾部分均为某一些特殊字符
*/
//'d'为数字	's'为+-	'e'为幂符号	'.'为小数点
func isNumber(s string) bool {
	var states []map[interface{}]interface{}
	states = make([]map[interface{}]interface{}, 9)
	for i := 0; i < 9; i++ {
		states[i] = make(map[interface{}]interface{})
	}
	//起始部分
	states[0][' '] = 0
	states[0]['s'] = 1
	states[0]['d'] = 2
	states[0]['.'] = 4
	//第一个中间部分
	states[1]['d'] = 2
	states[1]['.'] = 4
	//中间部分为数字
	states[2]['d'] = 2
	states[2]['.'] = 3
	states[2]['e'] = 5
	states[2][' '] = 8
	//中间部分为中间小数点后的数字
	states[3]['d'] = 3
	states[3]['e'] = 5
	states[3][' '] = 8
	//中间部分为开头小数点后的数字，因为有开头小数点不能再跟小数点
	states[4]['d'] = 3
	//中间部分为e
	states[5]['s'] = 6
	states[5]['d'] = 7
	//+-号后的中间部分
	states[6]['d'] = 7
	//中间部分e后数字
	states[7]['d'] = 7
	states[7][' '] = 8
	//空格后的空格
	states[8][' '] = 8

	p := 0 //标识状态的变量
	t := ' '
	for _, value := range s {
		if value >= '0' && value <= '9' {
			t = 'd'
		} else if value == '+' || value == '-' {
			t = 's'
		} else if value == 'e' || value == 'E' {
			t = 'e'
		} else if value == '.' || value == ' ' {
			t = value
		} else {
			t = '?'
		}
		//fmt.Println(value,p, t, states[p][t])
		if states[p][t] == nil {
			return false
		}
		p = states[p][t].(int)
	}
	return p == 2 || p == 3 || p == 7 || p == 8
}
