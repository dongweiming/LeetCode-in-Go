package problem0155

type MinStack struct {
	items  []int
	helper []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.items = append(this.items, x)
	if len(this.helper) == 0 || x <= this.helper[len(this.helper)-1] {
		this.helper = append(this.helper, x)
	}
}

func (this *MinStack) Pop() {
	n := len(this.items) - 1
	if this.helper[len(this.helper)-1] == this.items[n] {
		this.helper = this.helper[:len(this.helper)-1]
	}
	this.items = this.items[:n]
}

func (this *MinStack) Top() int {
	return this.items[len(this.items)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.helper) > 0 {
		return this.helper[len(this.helper)-1]
	}
	return 0
}
