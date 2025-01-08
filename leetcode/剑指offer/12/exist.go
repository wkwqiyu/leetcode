package _2

type point struct {
	x int
	y int
}

func exist(board [][]byte, word string) bool {
	for index1, value1 := range board {
		for index2, value2 := range value1 {
			if value2 == word[0] {
				if dfs(board, word, index1, index2, 0) { //当找到真的路径时返回true
					return true
				}
			}
		}
	}
	return false
}

// i和j为board的当前矩阵元素的位置，k为word中当前矩阵的字母
func dfs(board [][]byte, word string, i, j, k int) bool {
	if i >= len(board) || j >= len(board[0]) || j < 0 || i < 0 || board[i][j] != word[k] {
		return false //当i和j超出检索范围或不等于目标字母时，返回false
	}
	if k == len(word)-1 {
		return true //字母表到了最后一个，也就是路径满足要求时，返回true
	}
	board[i][j] = ' '
	//该矩阵的四个方向全部递归查找，其中有一个返回true就证明有满足要求的路径
	res := dfs(board, word, i+1, j, k+1) || dfs(board, word, i-1, j, k+1) ||
		dfs(board, word, i, j+1, k+1) || dfs(board, word, i, j-1, k+1)
	board[i][j] = word[k] //修正数组到原值
	return res
}

/*class Solution {
public boolean exist(char[][] board, String word) {
char[] words = word.toCharArray();
for(int i = 0; i < board.length; i++) {
for(int j = 0; j < board[0].length; j++) {
if(dfs(board, words, i, j, 0)) return true;
}
}
return false;
}
boolean dfs(char[][] board, char[] word, int i, int j, int k) {
if(i >= board.length || i < 0 || j >= board[0].length || j < 0 || board[i][j] != word[k]) return false;
if(k == word.length - 1) return true;
board[i][j] = '\0';
boolean res = dfs(board, word, i + 1, j, k + 1) || dfs(board, word, i - 1, j, k + 1) ||
dfs(board, word, i, j + 1, k + 1) || dfs(board, word, i , j - 1, k + 1);
board[i][j] = word[k];
return res;
}
}

作者：jyd
链接：https://leetcode-cn.com/problems/ju-zhen-zhong-de-lu-jing-lcof/solution/mian-shi-ti-12-ju-zhen-zhong-de-lu-jing-shen-du-yo/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。*/
