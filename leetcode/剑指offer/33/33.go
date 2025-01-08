package _3

func verifyPostorder(postorder []int) bool {
	return recur(postorder, 0, (len(postorder) - 1))
}

func recur(postorder []int, left int, right int) bool {
	if left >= right { //如果结点少于2个就返回true，证明没有子树
		return true
	}
	p := left
	for postorder[p] < postorder[right] {
		p++
	} //z找到左子树的结束，右子树的开始
	m := p
	for postorder[p] > postorder[right] {
		p++
	} //一直判断右子树的所有结点是否都大于根结点
	//如果有右子树的结点小于根结点就证明该树不是二叉搜索树
	return p == right && recur(postorder, left, m-1) && recur(postorder, m, right-1)
}
