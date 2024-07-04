package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	timeMap := Constructor()

	timeMap.Set("foo", "bar", 1)
	assert.Equal(t, "bar", timeMap.Get("foo", 1))
	assert.Equal(t, "bar", timeMap.Get("foo", 3))

	timeMap.Set("foo", "bar2", 4)
	assert.Equal(t, "bar2", timeMap.Get("foo", 4))
	assert.Equal(t, "bar2", timeMap.Get("foo", 5))
}
