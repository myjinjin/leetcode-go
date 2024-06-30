package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	lRUCache := Constructor(2)
	lRUCache.Put(1, 1)
	lRUCache.Put(2, 2)
	assert.Equal(t, 1, lRUCache.Get(1))

	lRUCache.Put(3, 3)
	assert.Equal(t, -1, lRUCache.Get(2))

	lRUCache.Put(4, 4)
	assert.Equal(t, -1, lRUCache.Get(1))
	assert.Equal(t, 3, lRUCache.Get(3))
	assert.Equal(t, 4, lRUCache.Get(4))
}
