package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeIntervals(t *testing.T) {
	// Since the first two intervals [1,4] and [2,5] overlap, we merged them into one [1,5].
	assert.Equal(t, [][]int{{1, 5}, {7, 9}}, MergeIntervals([][]int{{1, 4}, {2, 5}, {7, 9}}))

	// Since the intervals [6,7] and [5,9] overlap, we merged them into one [5,9].
	assert.Equal(t, [][]int{{2, 4}, {5, 9}}, MergeIntervals([][]int{{6, 7}, {2, 4}, {5, 9}}))

	// Since all the given intervals overlap, we merged them into one.
	assert.Equal(t, [][]int{{1, 6}}, MergeIntervals([][]int{{1, 4}, {2, 6}, {3, 5}}))
}
