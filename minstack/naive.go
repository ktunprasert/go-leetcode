package minstack

type MinStack struct {
	hashMap map[int]int
	tail    int
	minMap  map[int]int
}

func Constructor() MinStack {
	stack := MinStack{
		make(map[int]int),
		0,
		make(map[int]int),
	}
	return stack
}

func (this *MinStack) Push(val int) {
	this.hashMap[this.tail] = val
	if len(this.minMap) == 0 {
		this.minMap[this.tail] = val
	} else {
		this.minMap[this.tail] = min(val, this.minMap[this.tail-1])
	}
	this.tail++
}

func (this *MinStack) Pop() {
	this.tail--
	delete(this.hashMap, this.tail)
	delete(this.minMap, this.tail)
}

func (this *MinStack) Top() int {
	return this.hashMap[this.tail-1]
}

func (this *MinStack) GetMin() int {
	return this.minMap[this.tail-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
