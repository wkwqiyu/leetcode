package CQueue

type CQueue struct {
	array []int
}

func Constructor() CQueue {

	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.array = append(this.array, value)
}

func (this *CQueue) DeleteHead() int {
	var i int
	if len(this.array) == 0 {
		return -1
	} else {
		i = this.array[0]
		this.array = append(this.array[1:])
		return i
	}
}
