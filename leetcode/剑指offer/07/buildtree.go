package _7

/**
  Definition for a binary tree node.*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归，整个树的子树，并返回对应的结点
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 { //不存在的子树
		return nil
	}
	root := TreeNode{
		Val:   preorder[0],
		Left:  nil,
		Right: nil,
	}
	root.Val = preorder[0]
	i := 0
	//find root in inorder
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	//递归
	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return &root
}
