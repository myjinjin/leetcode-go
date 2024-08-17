package leetcode

type LockingTree struct {
	parent   []int
	children [][]int
	locked   map[int]int
}

func Constructor(parent []int) LockingTree {
	n := len(parent)
	children := make([][]int, n)
	for i := 1; i < n; i++ {
		p := parent[i]
		children[p] = append(children[p], i)
	}
	return LockingTree{
		parent:   parent,
		children: children,
		locked:   make(map[int]int),
	}
}

func (l *LockingTree) Lock(num int, user int) bool {
	if _, exist := l.locked[num]; exist {
		return false
	}
	l.locked[num] = user
	return true
}

func (l *LockingTree) Unlock(num int, user int) bool {
	if u, exist := l.locked[num]; !exist || u != user {
		return false
	}
	delete(l.locked, num)
	return true
}

func (l *LockingTree) Upgrade(num int, user int) bool {
	if _, exist := l.locked[num]; exist {
		return false
	}

	lockedChildren := l.findLockedChildren(num)
	if len(lockedChildren) == 0 {
		return false
	}

	if l.hasLockedParents(num) {
		return false
	}

	l.locked[num] = user
	for _, c := range lockedChildren {
		delete(l.locked, c)
	}
	return true
}

func (l *LockingTree) findLockedChildren(num int) []int {
	lockedChildren := []int{}

	var dfs func(num int)
	dfs = func(num int) {
		if _, exist := l.locked[num]; exist {
			lockedChildren = append(lockedChildren, num)
		}
		for _, c := range l.children[num] {
			dfs(c)
		}
	}

	for _, c := range l.children[num] {
		dfs(c)
	}

	return lockedChildren
}

func (l *LockingTree) hasLockedParents(num int) bool {
	p := l.parent[num]
	for p != -1 {
		if _, exist := l.locked[p]; exist {
			return true
		}
		p = l.parent[p]
	}
	return false
}

/**
 * Your LockingTree object will be instantiated and called as such:
 * obj := Constructor(parent);
 * param_1 := obj.Lock(num,user);
 * param_2 := obj.Unlock(num,user);
 * param_3 := obj.Upgrade(num,user);
 */
