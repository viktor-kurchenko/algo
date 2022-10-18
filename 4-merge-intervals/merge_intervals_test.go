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

func TestMinimumMeetingRooms(t *testing.T) {
	assert.Equal(t, 2, MinimumMeetingRooms([][]int{{1, 4}, {2, 5}, {7, 9}}))

	assert.Equal(t, 1, MinimumMeetingRooms([][]int{{6, 7}, {2, 4}, {8, 12}}))

	assert.Equal(t, 2, MinimumMeetingRooms([][]int{{1, 4}, {2, 3}, {3, 6}}))

	assert.Equal(t, 2, MinimumMeetingRooms([][]int{{4, 5}, {2, 3}, {2, 4}, {3, 5}}))
}

func TestMaxCPULoad(t *testing.T) {
	// Since [1,4,3] and [2,5,4] overlap, their maximum CPU load (3+4=7) will be when both the jobs are running at the same time i.e., during the time interval (2,4).
	assert.Equal(t, 7, MaxCPULoad([][]int{{1, 4, 3}, {2, 5, 4}, {7, 9, 6}}))

	// None of the jobs overlap, therefore we will take the maximum load of any job which is 15.
	assert.Equal(t, 15, MaxCPULoad([][]int{{6, 7, 10}, {2, 4, 11}, {8, 12, 15}}))

	// Maximum CPU load will be 8 as all jobs overlap during the time interval [3,4].
	assert.Equal(t, 8, MaxCPULoad([][]int{{1, 4, 2}, {2, 4, 1}, {3, 6, 5}}))
}

func TestEmployeeFreeTime(t *testing.T) {
	// Both the employees are free between [3,5].
	assert.Equal(t, [][]int{{3, 5}}, EmployeeFreeTime([][][]int{{{1, 3}, {5, 6}}, {{2, 3}, {6, 8}}}))

	// All employees are free between [4,6] and [8,9].
	assert.Equal(t, [][]int{{4, 6}, {8, 9}}, EmployeeFreeTime([][][]int{{{1, 3}, {9, 12}}, {{2, 4}}, {{6, 8}}}))

	// All employees are free between [5,7].
	assert.Equal(t, [][]int{{5, 7}}, EmployeeFreeTime([][][]int{{{1, 3}}, {{2, 4}}, {{3, 5}, {7, 9}}}))
}
