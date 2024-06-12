package leetcode

type SmallestInfiniteSet struct {
	removed    map[int]bool
	currentMin int
}

func Constructor() SmallestInfiniteSet {
	return SmallestInfiniteSet{
		removed:    map[int]bool{},
		currentMin: 1,
	}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	var min = this.currentMin
	this.removed[this.currentMin] = true
	this.currentMin++
	for this.removed[this.currentMin] {
		this.currentMin++
	}
	return min
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if _, ok := this.removed[num]; ok {
		this.removed[num] = false
	}
	if num < this.currentMin {
		this.currentMin = num
	}
}

/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */
