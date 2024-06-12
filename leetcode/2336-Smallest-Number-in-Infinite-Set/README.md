# [2336. Smallest Number in Infinite Set](https://leetcode.com/problems/smallest-number-in-infinite-set/)

## Problem

You have a set which contains all positive integers `[1, 2, 3, 4, 5, ...]`.

Implement the `SmallestInfiniteSet` class:

- `SmallestInfiniteSet()` Initializes the SmallestInfiniteSet object to contain all positive integers.
- `int popSmallest()` Removes and returns the smallest integer contained in the infinite set.
- `void addBack(int num)` Adds a positive integer `num` back into the infinite set, if it is not already in the infinite set.
 

Example 1:

```
Input
["SmallestInfiniteSet", "addBack", "popSmallest", "popSmallest", "popSmallest", "addBack", "popSmallest", "popSmallest", "popSmallest"]
[[], [2], [], [], [], [1], [], [], []]
Output
[null, null, 1, 2, 3, null, 1, 4, 5]

Explanation
SmallestInfiniteSet smallestInfiniteSet = new SmallestInfiniteSet();
smallestInfiniteSet.addBack(2);    // 2 is already in the set, so no change is made.
smallestInfiniteSet.popSmallest(); // return 1, since 1 is the smallest number, and remove it from the set.
smallestInfiniteSet.popSmallest(); // return 2, and remove it from the set.
smallestInfiniteSet.popSmallest(); // return 3, and remove it from the set.
smallestInfiniteSet.addBack(1);    // 1 is added back to the set.
smallestInfiniteSet.popSmallest(); // return 1, since 1 was added back to the set and
                                   // is the smallest number, and remove it from the set.
smallestInfiniteSet.popSmallest(); // return 4, and remove it from the set.
smallestInfiniteSet.popSmallest(); // return 5, and remove it from the set.
```

Constraints:

- `1 <= num <= 1000`
- At most `1000` calls will be made in total to `popSmallest` and `addBack`.

## Solution

```go
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
```