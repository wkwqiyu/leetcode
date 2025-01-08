package _5

import "fmt"

// 与下一个比大小
// 比他小，就把当前的指针指向目标，目标指向下一个
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	l := new(ListNode)
	head := l
	if l1.Val <= l2.Val {
		l.Val = l1.Val
		l1 = l1.Next
	} else {
		l.Val = l2.Val
	}
	for l2 != nil || l1 != nil {
		if l1.Val <= l2.Val {
			l.Next = l1
			fmt.Println(l.Val, "<-", l1.Val)
			l1 = l1.Next
			l = l.Next
		} else {
			l.Next = l2
			fmt.Println(l.Val, "<-", l2.Val)
			l2 = l2.Next
			l = l.Next
		}
	}
	if l1 != nil {
		l.Next = l1
	} else if l2 != nil {
		l.Next = l2
	}
	return head
}

func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else if l1.Val < l2.Val {
		l1.Next = mergeTwoLists1(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists1(l1, l2.Next)
		return l2
	}
}

/*class Solution {
public ListNode mergeTwoLists(ListNode l1, ListNode l2) {
if (l1 == null) {
return l2;
} else if (l2 == null) {
return l1;
} else if (l1.val < l2.val) {
l1.next = mergeTwoLists(l1.next, l2);
return l1;
} else {
l2.next = mergeTwoLists(l1, l2.next);
return l2;
}
}
}

作者：LeetCode-Solution
链接：https://leetcode-cn.com/problems/he-bing-liang-ge-pai-xu-de-lian-biao-lcof/solution/he-bing-liang-ge-pai-xu-de-lian-biao-by-g3z6g/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。*/
