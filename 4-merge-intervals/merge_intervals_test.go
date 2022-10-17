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

func TestInsertInterval(t *testing.T) {
	// After insertion, since [4,6] overlaps with [5,7], we merged them into one [4,7].
	assert.Equal(t, [][]int{{1, 3}, {4, 7}, {8, 12}}, InsertInterval([][]int{{1, 3}, {5, 7}, {8, 12}}, []int{4, 6}))

	// After insertion, since [4,10] overlaps with [5,7] & [8,12], we merged them into [4,12].
	assert.Equal(t, [][]int{{1, 3}, {4, 12}}, InsertInterval([][]int{{1, 3}, {5, 7}, {8, 12}}, []int{4, 10}))

	// After insertion, since [1,4] overlaps with [2,3], we merged them into one [1,4].
	assert.Equal(t, [][]int{{1, 4}, {5, 7}}, InsertInterval([][]int{{2, 3}, {5, 7}}, []int{1, 4}))
}

func TestIntervalsIntersection(t *testing.T) {
	assert.Equal(t, [][]int{{2, 3}, {5, 6}, {7, 7}}, IntervalsIntersection([][]int{{1, 3}, {5, 6}, {7, 9}}, [][]int{{2, 3}, {5, 7}}))

	assert.Equal(t, [][]int{{5, 7}, {9, 10}}, IntervalsIntersection([][]int{{1, 3}, {5, 7}, {9, 12}}, [][]int{{5, 10}}))
}

func TestConflictingAppointments(t *testing.T) {
	// Since [1,4] and [2,5] overlap, a person cannot attend both of these appointments.
	assert.False(t, ConflictingAppointments([][]int{{1, 4}, {2, 5}, {7, 9}}))

	// None of the appointments overlap, therefore a person can attend all of them.
	assert.True(t, ConflictingAppointments([][]int{{6, 7}, {2, 4}, {8, 12}}))

	// Since [4,5] and [3,6] overlap, a person cannot attend both of these appointments.
	assert.False(t, ConflictingAppointments([][]int{{4, 5}, {2, 3}, {3, 6}}))
}
