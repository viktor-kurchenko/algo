package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPairWithTargetSum(t *testing.T) {
	// The numbers at index 1 and 3 add up to 6: 2+4=6
	assert.Equal(t, []int{1, 3}, PairWithTargetSum([]int{1, 2, 3, 4, 6}, 6))

	// The numbers at index 0 and 2 add up to 11: 2+9=11
	assert.Equal(t, []int{0, 2}, PairWithTargetSum([]int{2, 5, 9, 11}, 11))
}

func TestRemoveDuplicates(t *testing.T) {
	// The first four elements after removing the duplicates will be [2, 3, 6, 9].
	assert.Equal(t, 4, RemoveDuplicates([]int{2, 3, 3, 3, 6, 9, 9}))

	// The first two elements after removing the duplicates will be [2, 11].
	assert.Equal(t, 2, RemoveDuplicates([]int{2, 2, 2, 11}))
}

func TestRemoveKey(t *testing.T) {
	// The first four elements after removing every 'Key' will be [2, 6, 10, 9].
	assert.Equal(t, 4, RemoveKey([]int{3, 2, 3, 6, 3, 10, 9, 3}, 3))

	// The first two elements after removing every 'Key' will be [11, 1].
	assert.Equal(t, 2, RemoveKey([]int{2, 11, 2, 2, 1}, 2))
}

func TestSquaringSortedArray(t *testing.T) {
	assert.Equal(t, []int{0, 1, 4, 4, 9}, SquaringSortedArray([]int{-2, -1, 0, 2, 3}))

	assert.Equal(t, []int{0, 1, 1, 4, 9}, SquaringSortedArray([]int{-3, -1, 0, 1, 2}))
}

func TestTripletSum2Zero(t *testing.T) {
	assert.Equal(t, [][]int{{-3, 1, 2}, {-2, 0, 2}, {-2, 1, 1}, {-1, 0, 1}}, TripletSum2Zero([]int{-3, 0, 1, 2, -1, 1, -2}))

	assert.Equal(t, [][]int{{-5, 2, 3}, {-2, -1, 3}}, TripletSum2Zero([]int{-5, 2, -1, -2, 3}))
}

func TestTripletSumCloseToTarget(t *testing.T) {
	// The triplet [-2, 1, 2] has the closest sum to the target.
	assert.Equal(t, 1, TripletSumCloseToTarget([]int{-2, 0, 1, 2}, 2))

	// The triplet [-3, 1, 2] has the closest sum to the target.
	assert.Equal(t, 0, TripletSumCloseToTarget([]int{-3, -1, 1, 2}, 1))

	// The triplet [1, 1, 1] has the closest sum to the target.
	assert.Equal(t, 3, TripletSumCloseToTarget([]int{1, 0, 1, 1}, 100))
}

func TestTripletWithSmallerSum(t *testing.T) {
	// There are two triplets whose sum is less than the target: [-1, 0, 3], [-1, 0, 2]
	assert.Equal(t, 2, TripletWithSmallerSum([]int{-1, 0, 2, 3}, 3))

	// There are four triplets whose sum is less than the target:
	// [-1, 1, 4], [-1, 1, 3], [-1, 1, 2], [-1, 2, 3]
	assert.Equal(t, 4, TripletWithSmallerSum([]int{-1, 4, 2, 1, 3}, 5))
}
