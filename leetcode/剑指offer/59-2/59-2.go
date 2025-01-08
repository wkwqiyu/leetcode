package _9_2

type MaxQueue struct {
	max   []int
	queue []int
}

func Constructor() MaxQueue {
	return MaxQueue{}
}

func (this *MaxQueue) Max_value() int {
	if len(this.max) == 0 {
		return -1
	}
	return this.max[0]
}

func (this *MaxQueue) Push_back(value int) {
	for len(this.max) > 0 && value > this.max[len(this.max)-1] {
		//fmt.Println(this.max)
		this.max = this.max[:len(this.max)-1]
	}
	this.max = append(this.max, value)
	this.queue = append(this.queue, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.queue) == 0 {
		return -1
	}
	if this.max[0] == this.queue[0] {
		this.max = this.max[1:]
	}
	a := this.queue[0]
	this.queue = this.queue[1:]
	return a
}
