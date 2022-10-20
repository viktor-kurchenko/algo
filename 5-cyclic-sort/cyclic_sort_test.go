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
