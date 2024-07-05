package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	snapshots := Constructor(3)
	snapshots.Set(0, 5)

	assert.Equal(t, 0, snapshots.Snap())

	snapshots.Set(0, 6)
	assert.Equal(t, 5, snapshots.Get(0, 0))
}
