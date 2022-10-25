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

func TestAllMissingNumbers(t *testing.T) {
	assert.Equal(t, []int{4, 6, 7}, AllMissingNumbers([]int{2, 3, 1, 8, 2, 3, 5, 1}))

	assert.Equal(t, []int{3}, AllMissingNumbers([]int{2, 4, 1, 2}))

	assert.Equal(t, []int{4}, AllMissingNumbers([]int{2, 3, 2, 1}))
}

func TestFindDuplicate(t *testing.T) {
	assert.Equal(t, 4, FindDuplicate([]int{1, 4, 4, 3, 2}))

	assert.Equal(t, 3, FindDuplicate([]int{2, 1, 3, 3, 5, 4}))

	assert.Equal(t, 4, FindDuplicate([]int{2, 4, 1, 4, 4}))
}

func TestFindAllDuplicates(t *testing.T) {
	assert.Equal(t, []int{4, 5}, FindAllDuplicates([]int{3, 4, 4, 5, 5}))

	assert.Equal(t, []int{3, 5}, FindAllDuplicates([]int{5, 4, 7, 2, 3, 5, 3}))
}

func TestFindCorruptPair(t *testing.T) {
	assert.Equal(t, []int{2, 4}, FindCorruptPair([]int{3, 1, 2, 5, 2}))

	assert.Equal(t, []int{3, 5}, FindCorruptPair([]int{3, 1, 2, 3, 6, 4}))
}

func TestFindSmallestMissingPositiveNumber(t *testing.T) {
	assert.Equal(t, 3, FindSmallestMissingPositiveNumber([]int{-3, 1, 5, 4, 2}))

	assert.Equal(t, 4, FindSmallestMissingPositiveNumber([]int{3, -2, 0, 1, 2}))

	assert.Equal(t, 4, FindSmallestMissingPositiveNumber([]int{3, 2, 5, 1}))
}
