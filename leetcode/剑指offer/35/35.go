package _5

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	var m map[*Node]*Node
	m = make(map[*Node]*Node)
	m[nil] = nil
	if head == nil {
		return nil
	}
	prehead := head.Next
	root := new(Node)
	root.Val = head.Val
	preroot := root.Next
	m[head] = root
	for prehead != nil {
		//通过map建立两表之间的连接
		m[prehead] = preroot
		node := new(Node)
		preroot.Val = prehead.Val
		preroot.Next = node
		preroot = node
		prehead = prehead.Next
	}
	//建立新链表的random:遍历新链表，赋值random
	preroot = root
	for preroot != nil {
		preroot.Random = m[head.Random]
		preroot = preroot.Next
		head = head.Next
	}
	return root
}

/*func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	// 建立一个哈希表，为原节点到新节点的映射关系
	m := map[*Node]*Node{}
	m[nil] = nil

	// 建立头部节点，用于返回，同时创建映射关系
	root := &Node{
		Val : head.Val,
	}
	m[head] = root

	// 建立临时节点，用于第一次迭代
	prev := root
	to := head.Next

	// 第一次迭代，创建层层相连的单链表，也就是创建节点并用Next穿起来
	for to != nil {
		node := &Node{}
		prev.Next = node
		m[to] = node

		node.Val = to.Val
		to = to.Next
		prev = node
	}

	// 第二次迭代，去哈希表中找原节点对于的新节点，用Random穿起来
	curr := root
	for head != nil {
		curr.Random = m[head.Random]
		curr = curr.Next
		head = head.Next
	}


	return root
}
*/
/*

func copyRandomList(head *Node) *Node {
	var copyhead *Node
	copyhead = new(Node)
	list := copyhead
	prehead := head
	var tmp *Node
	if head == nil {
		return nil
	}
	for prehead != nil {
		//通过原链表建立新链表
		list.Val = prehead.Val
		list.Next = new(Node)
		//把现链表的random指向原链表
		list.Random = prehead
		//把原链表的next指向现链表对应的结点
		tmp = prehead.Next
		prehead.Next = list
        //fmt.Println(list,prehead.Next)
		list = list.Next
		prehead = tmp
	}
    list = nil
    fmt.Println("-=-=-=-=-==-")
	//建立新链表的random:遍历新链表，把random指向的原结点的random结点的next赋给random
	list = copyhead
    for list !=nil{
        //fmt.Println(list,list.Random,list.Random.Next)
        list = list.Next
    }
    list = copyhead
	for list.Next != nil {
        //fmt.Print(list.Val)
        //fmt.Println(list.Random.Val,list.Random.Random.Val)
        //fmt.Println(list.Random.Random,list.Random.Random.Next,list)
		tmp = list.Random
		if tmp.Random != nil {
            fmt.Println(tmp.Random.Val,tmp.Random.Next.Val)
			list.Random = tmp.Random.Next
		} else {
			list.Random = nil
		}
		tmp.Next = list.Next.Random
        //fmt.Println(tmp.Val)
		list = list.Next
	}

	return copyhead
}*/
