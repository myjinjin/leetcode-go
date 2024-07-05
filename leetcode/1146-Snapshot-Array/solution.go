package leetcode

type Snapshot struct {
	snapID int
	value  int
}

type SnapshotArray struct {
	snapshots [][]Snapshot
	snapID    int
}

func Constructor(length int) SnapshotArray {
	snapshots := make([][]Snapshot, length)
	for i := range snapshots {
		snapshots[i] = []Snapshot{{snapID: 0, value: 0}}
	}
	return SnapshotArray{
		snapshots: snapshots,
		snapID:    0,
	}
}

func (a *SnapshotArray) Set(index int, val int) {
	snapshots := a.snapshots[index]
	if snapshots[len(snapshots)-1].snapID == a.snapID {
		snapshots[len(snapshots)-1].value = val
	} else {
		a.snapshots[index] = append(snapshots, Snapshot{value: val, snapID: a.snapID})
	}
}

func (a *SnapshotArray) Snap() int {
	a.snapID++
	return a.snapID - 1
}

func (a *SnapshotArray) Get(index int, snapID int) int {
	snapshots := a.snapshots[index]
	left, right := 0, len(snapshots)-1
	for left <= right {
		mid := left + (right-left)/2
		if snapshots[mid].snapID <= snapID {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return snapshots[right].value
}
