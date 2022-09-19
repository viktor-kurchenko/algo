package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubArrayAvg(t *testing.T) {
	assert.Equal(t, []float32{2.2, 2.8, 2.4, 3.6, 2.8}, SubArrayAvg([]float32{1, 3, 2, 6, -1, 4, 1, 8, 2}, 5))
}

func TestSubArrayMaxSum(t *testing.T) {
	// 9 -> [3, 4]
	assert.Equal(t, 9, SubArrayMaxSum([]int{2, 1, 5, 1, 3, 2}, 3))

	// 7 -> [5, 1, 3]
	assert.Equal(t, 7, SubArrayMaxSum([]int{2, 3, 4, 1, 5}, 2))
}

func TestSmallestSubArrayWithGreaterSum(t *testing.T) {
	// 7 -> [5, 2]
	assert.Equal(t, 2, SmallestSubArrayWithGreaterSum([]int{2, 1, 5, 2, 3, 2}, 7))

	// 7 -> [8]
	assert.Equal(t, 1, SmallestSubArrayWithGreaterSum([]int{2, 1, 5, 2, 8}, 7))

	// 8 -> [3, 4, 1] or [1, 1, 6]
	assert.Equal(t, 3, SmallestSubArrayWithGreaterSum([]int{3, 4, 1, 1, 6}, 8))
}
