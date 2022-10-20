package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCyclicSort(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 4, 5}, CyclicSort([]int{3, 1, 5, 4, 2}))

	assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, CyclicSort([]int{2, 6, 4, 3, 1, 5}))

	assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, CyclicSort([]int{1, 5, 6, 4, 3, 2}))
}

func TestMissingNumber(t *testing.T) {
	assert.Equal(t, 2, MissingNumber([]int{4, 0, 3, 1}))

	assert.Equal(t, 7, MissingNumber([]int{8, 3, 5, 2, 4, 6, 0, 1}))
}
