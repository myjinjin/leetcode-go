# [1146. Snapshot Array](https://leetcode.com/problems/snapshot-array/)

## Problem

Implement a SnapshotArray that supports the following interface:

- `SnapshotArray(int length)` initializes an array-like data structure with the given length. Initially, each element equals 0.
- `void set(index, val)` sets the element at the given `index` to be equal to `val`.
- `int snap()` takes a snapshot of the array and returns the `snap_id`: the total number of times we called `snap()` minus `1`.
- `int get(index, snap_id)` returns the value at the given `index`, at the time we took the snapshot with the given `snap_id`
 

Example 1:

```
Input: ["SnapshotArray","set","snap","set","get"]
[[3],[0,5],[],[0,6],[0,0]]
Output: [null,null,0,null,5]
Explanation: 
SnapshotArray snapshotArr = new SnapshotArray(3); // set the length to be 3
snapshotArr.set(0,5);  // Set array[0] = 5
snapshotArr.snap();  // Take a snapshot, return snap_id = 0
snapshotArr.set(0,6);
snapshotArr.get(0,0);  // Get the value of array[0] with snap_id = 0, return 5
``` 

Constraints:

- `1 <= length <= 5 * 10^4`
- `0 <= index < length`
- `0 <= val <= 10^9`
- `0 <= snap_id <` (the total number of times we call `snap()`)
- At most `5 * 10^4` calls will be made to `set`, `snap`, and `get`.

## Solution

```go
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
```