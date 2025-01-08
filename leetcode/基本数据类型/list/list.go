package list

type List struct {
	next   *List
	number interface{}
}

//用"container/list"包下自带的list
//queue :=list.New
