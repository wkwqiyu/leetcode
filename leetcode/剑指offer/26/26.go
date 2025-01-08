package _6

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	var a bool
	if A == nil || B == nil {
		return false
	}
	//如果遇到相同数值的结点，进入匹配模式，先序遍历每个结点并匹配，全部满足返回true
	if A.Val == B.Val {
		a = next(A.Left, B.Left) && next(A.Right, B.Right)

	}
	//不对就进行下一个A的结点进行匹配
	return a || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

// 先序遍历的单独步骤，如果左右子节点相同，就进入下一个左右子节点递归
// 如果不同，返回false
// 如果A遍历完成而不没有遍历完成，证明长度不匹配，返回false
// 如果B遍历完成，证明B的其中一个没有子节点的结点遍历完成，如果所有的没有子节点的结点都遍历完成，证明B是A的子树，返回true
func next(A *TreeNode, B *TreeNode) bool {
	if A == nil && B != nil {
		return false
	}
	if B == nil {
		return true
	}
	if A.Val == B.Val {
		return next(A.Left, B.Left) && next(A.Right, B.Right)
	} else {
		return false
	}
}
